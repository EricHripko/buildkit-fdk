package cib

import (
	"strings"

	"github.com/containerd/containerd/platforms"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/solver/pb"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
)

// Taken from https://github.com/moby/buildkit/blob/master/frontend/dockerfile/builder/build.go#L33
// for compatibility

const (
	buildArgPrefix       = "build-arg:"
	dockerignoreFilename = ".dockerignore"
	keyImageResolveMode  = "image-resolve-mode"
	keyFilename          = "filename"
	keyTargetPlatform    = "platform"
	keyNoCache           = "no-cache"
)

// Taken from https://github.com/moby/buildkit/blob/master/frontend/dockerfile/dockerfile2llb/convert.go#L1160
// for compatibility

func parseKeyValue(env string) (string, string) {
	parts := strings.SplitN(env, "=", 2)
	v := ""
	if len(parts) > 1 {
		v = parts[1]
	}

	return parts[0], v
}

// Taken from https://github.com/moby/buildkit/blob/master/frontend/dockerfile/builder/build.go
// for compatibility

func parsePlatforms(v string) ([]*specs.Platform, error) {
	var pp []*specs.Platform
	for _, v := range strings.Split(v, ",") {
		p, err := platforms.Parse(v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse target platform %s", v)
		}
		p = platforms.Normalize(p)
		pp = append(pp, &p)
	}
	return pp, nil
}

func parseResolveMode(v string) (llb.ResolveMode, error) {
	switch v {
	case pb.AttrImageResolveModeDefault, "":
		return llb.ResolveModeDefault, nil
	case pb.AttrImageResolveModeForcePull:
		return llb.ResolveModeForcePull, nil
	case pb.AttrImageResolveModePreferLocal:
		return llb.ResolveModePreferLocal, nil
	default:
		return 0, errors.Errorf("invalid image-resolve-mode: %s", v)
	}
}

func filter(opt map[string]string, key string) map[string]string {
	m := map[string]string{}
	for k, v := range opt {
		if strings.HasPrefix(k, key) {
			m[strings.TrimPrefix(k, key)] = v
		}
	}
	return m
}
