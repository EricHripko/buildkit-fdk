// Package cib contains the primitives for container image building.
package cib

import (
	"context"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/dockerfile/dockerfile2llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/moby/buildkit/util/apicaps"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

//go:generate mockgen -package cib_mock -destination mock/cib.go . Service
//go:generate mockgen -package cib_mock -destination mock/client.go github.com/moby/buildkit/frontend/gateway/client Client,Reference

// Service that wraps low-level BuildKit actions and provides a simple
// interface for container image build.
type Service interface {
	// GetOpts returns client options.
	GetOpts() map[string]string
	// GetCaps returns the capabilities of the service.
	GetCaps() apicaps.CapSet
	// GetMarshalOpts returns options for marshalling LLB.
	GetMarshalOpts() []llb.ConstraintsOpt
	// GetResolveMode returns the resolution mode for base images.
	GetResolveMode() (llb.ResolveMode, error)
	// GetMetadataFileName returns the name for the metadata file.
	GetMetadataFileName() string
	// GetMetadata returns the contents of the metadata file.
	GetMetadata() ([]byte, error)
	// GetExcludes returns the list of paths to exclude from the build context.
	GetExcludes() ([]string, error)
	// GetBuildArgs returns the list of build-time variables.
	GetBuildArgs() map[string]string
	// GetBuildPlatform returns the build platform for the client.
	GetBuildPlatform() *specs.Platform
	// GetTargetPlatforms returns the list of target platforms for the image(s) being built.
	GetTargetPlatforms() ([]*specs.Platform, error)
	// GetIgnoreCache indicates whether the cache should be ignored.
	GetIgnoreCache() bool

	// Src returns the reference to the source code (created lazily).
	Src() (ref client.Reference, err error)
	// SrcState returns the LLB state of the source code (created lazily).
	SrcState() (state llb.State, err error)

	// FetchImageConfig returns the configuration of the image for the given platform.
	FetchImageConfig(name string, platform *specs.Platform) (dockerfile2llb.Image, error)
	// From creates a new LLB state from the specified base image.
	From(base string, platform *specs.Platform, comment string) (llb.State, *dockerfile2llb.Image, error)
	// Solve the provided LLB state and return a single reference from it.
	Solve(ctx context.Context, state llb.State) (client.Reference, error)
}
