package cib

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/containerd/containerd/platforms"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/builder/dockerignore"
	"github.com/moby/buildkit/client/llb"
	dockerfile "github.com/moby/buildkit/frontend/dockerfile/builder"
	"github.com/moby/buildkit/frontend/dockerfile/dockerfile2llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/moby/buildkit/util/apicaps"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
)

type service struct {
	// Context for actions.
	ctx context.Context
	// Client for dispatching low-level actions.
	client client.Client
	// Source code used for image build (build context).
	src client.Reference
	// Context of the metadata file.
	metadata []byte
	// Paths to exclude from the build context.
	excludes []string
}

// NewService creates an instance of container image building service.
func NewService(ctx context.Context, client client.Client) Service {
	return &service{ctx, client, nil, nil, nil}
}

func (s *service) GetClient() client.Client {
	return s.client
}

func (s *service) initSrc() error {
	// LLB
	state := llb.Local(dockerfile.DefaultLocalNameContext,
		llb.SessionID(s.client.BuildOpts().SessionID),
		dockerfile2llb.WithInternalName("load context"),
	)
	var err error
	s.src, err = s.Solve(s.ctx, state)
	return err
}

func (s *service) Src() (ref client.Reference, err error) {
	if s.src == nil {
		err = s.initSrc()
	}
	ref = s.src
	return
}

func (s *service) SrcState() (state llb.State, err error) {
	if s.src == nil {
		err = s.initSrc()
	}
	if err == nil {
		state, err = s.src.ToState()
	}
	return
}

func (s *service) GetOpts() map[string]string {
	return s.client.BuildOpts().Opts
}

func (s *service) GetCaps() apicaps.CapSet {
	return s.client.BuildOpts().LLBCaps
}

func (s *service) GetMarshalOpts() []llb.ConstraintsOpt {
	return []llb.ConstraintsOpt{llb.WithCaps(s.GetCaps())}
}

func (s *service) GetResolveMode() (llb.ResolveMode, error) {
	return parseResolveMode(s.GetOpts()[keyImageResolveMode])
}

func (s *service) GetMetadataFileName() string {
	if name, ok := s.GetOpts()[keyFilename]; ok {
		return name
	}
	return "Dockerfile"
}

func (s *service) initMetadata() error {
	filename := s.GetMetadataFileName()

	// LLB
	state := llb.Local(dockerfile.DefaultLocalNameDockerfile,
		llb.FollowPaths([]string{filename}),
		llb.SessionID(s.client.BuildOpts().SessionID),
		llb.SharedKeyHint(dockerfile.DefaultLocalNameDockerfile),
		dockerfile2llb.WithInternalName("load build definition from "+filename),
	)
	// Solve
	ref, err := s.Solve(s.ctx, state)
	if err != nil {
		return err
	}
	// Read
	s.metadata, err = ref.ReadFile(s.ctx, client.ReadRequest{
		Filename: filename,
	})
	if err != nil {
		return errors.Wrapf(err, "failed to read metadata")
	}
	return nil
}

func (s *service) GetMetadata() (data []byte, err error) {
	if s.metadata == nil {
		err = s.initMetadata()
	}

	data = s.metadata
	return
}

func (s *service) initExcludes() error {
	// LLB
	state := llb.Local(dockerfile.DefaultLocalNameContext,
		llb.SessionID(s.client.BuildOpts().SessionID),
		llb.FollowPaths([]string{dockerignoreFilename}),
		llb.SharedKeyHint(dockerfile.DefaultLocalNameContext+"-"+dockerignoreFilename),
		dockerfile2llb.WithInternalName("load "+dockerignoreFilename),
	)
	// Solve
	ref, err := s.Solve(s.ctx, state)
	if err != nil {
		return err
	}
	// Read
	data, _ := ref.ReadFile(s.ctx, client.ReadRequest{
		Filename: dockerignoreFilename,
	})
	if data == nil {
		s.excludes = []string{}
	} else {
		excludes, err := dockerignore.ReadAll(bytes.NewBuffer(data))
		if excludes == nil || err != nil {
			return errors.Wrap(err, "failed to parse dockerignore")
		}
		s.excludes = excludes
	}
	return nil
}

func (s *service) GetExcludes() (excludes []string, err error) {
	if s.excludes == nil {
		err = s.initExcludes()
	}

	excludes = s.excludes
	return
}

func (s *service) FetchImageConfig(name string, platform *specs.Platform) (img dockerfile2llb.Image, err error) {
	resolveMode, err := s.GetResolveMode()
	if err != nil {
		return
	}

	_, data, err := s.client.ResolveImageConfig(s.ctx, name, llb.ResolveImageConfigOpt{
		Platform:    platform,
		ResolveMode: resolveMode.String(),
		LogName:     "load metadata for " + name,
	})
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &img); err != nil {
		return
	}
	img.Created = nil
	return
}

func (s *service) GetBuildArgs() map[string]string {
	return filter(s.GetOpts(), buildArgPrefix)
}

func (s *service) GetBuildPlatform() *specs.Platform {
	platform := platforms.DefaultSpec()
	if workers := s.client.BuildOpts().Workers; len(workers) > 0 && len(workers[0].Platforms) > 0 {
		platform = workers[0].Platforms[0]
	}
	return &platform
}

func (s *service) GetTargetPlatforms() (platforms []*specs.Platform, err error) {
	if v := s.GetOpts()[keyTargetPlatform]; v != "" {
		platforms, err = parsePlatforms(v)
		if err != nil {
			return
		}
	} else {
		platforms = []*specs.Platform{s.GetBuildPlatform()}
	}
	return
}

func (s *service) GetIgnoreCache() bool {
	_, ignoreCache := s.GetOpts()[keyNoCache]
	return ignoreCache
}

func (s *service) From(base string, platform *specs.Platform, comment string) (llb.State, *dockerfile2llb.Image, error) {
	// Parse
	ref, err := reference.ParseNormalizedNamed(base)
	if err != nil {
		return llb.State{}, nil, err
	}
	base = reference.TagNameOnly(ref).String()
	// Resolve
	resolveMode, err := s.GetResolveMode()
	if err != nil {
		return llb.State{}, nil, err
	}
	// LLB
	state := llb.Image(base, resolveMode, llb.Platform(*platform), llb.WithCustomName(comment))
	img, err := s.FetchImageConfig(base, platform)
	// Initialise with environment
	for _, env := range img.Config.Env {
		k, v := parseKeyValue(env)
		state = state.AddEnv(k, v)
	}
	// Initialise with user
	if img.Config.User != "" {
		state = state.User(img.Config.User)
	}
	return state, &img, err
}

func (s *service) Solve(ctx context.Context, state llb.State) (client.Reference, error) {
	// Marshal
	def, err := state.Marshal(ctx, s.GetMarshalOpts()...)
	if err != nil {
		return nil, err
	}
	// Solve
	res, err := s.client.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, err
	}
	return res.SingleRef()
}
