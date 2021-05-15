// Simple frontend that concatenates combines pre-configures Dockerfiles with
// the one provided by the build context.
package main

import (
	"context"
	"io/ioutil"

	"github.com/EricHripko/buildkit-fdk/pkg/dtp"

	dockerfile "github.com/moby/buildkit/frontend/dockerfile/builder"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	"github.com/moby/buildkit/util/appcontext"
	"github.com/sirupsen/logrus"
)

func build(ctx context.Context, c client.Client) (*client.Result, error) {
	// Read vendored Dockerfiles
	header, err := ioutil.ReadFile("header.Dockerfile")
	if err != nil {
		return nil, err
	}
	footer, err := ioutil.ReadFile("footer.Dockerfile")
	if err != nil {
		return nil, err
	}

	// Concatenate user-supplied Dockerfile with vendored ones
	transform := func(dockerfile []byte) ([]byte, error) {
		dockerfile = append(header, dockerfile...)
		dockerfile = append(dockerfile, footer...)
		return dockerfile, nil
	}
	if err := dtp.InjectDockerfileTransform(transform, c); err != nil {
		return nil, err
	}

	// Pass control to the upstream Dockerfile frontend
	return dockerfile.Build(ctx, c)
}

func main() {
	if err := grpcclient.RunFromEnvironment(appcontext.Context(), build); err != nil {
		logrus.Errorf("fatal error: %+v", err)
		panic(err)
	}
}
