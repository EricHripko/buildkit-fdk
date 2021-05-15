# Frontend Development kit for BuildKit

This repository houses the frontend Development Kit for [BuildKit](https://github.com/moby/buildkit).
It aims to make the development of custom frontends easier by removing
the unnecessary boilerplate. Generally speaking, there's _3_ ways to make use
of this (listed in the order of increasing complexity):

- Using [simple-frontend base image](https://hub.docker.com/r/erichripko/simple-frontend)
- Custom frontend that transforms the `Dockerfile`
- Custom frontend from scratch

## Using `simple-frontend` base image

[simple-frontend base image](https://hub.docker.com/r/erichripko/simple-frontend)
is a thin layer on top of a [docker/dockerfile frontend](https://hub.docker.com/r/docker/dockerfile).
It works by concatenating user-supplied `Dockerfile` with a pre-configured
header and footer. This enables you to create simple frontends with nearly
zero effort (and only requires knowledge of `Dockerfile` syntax).

Follow the steps below to create a frontend:

- Create a `Dockerfile` for your frontend with the following contents. Note
  that header and/or footer _can be omitted_ if not needed.

```dockerfile
FROM erichripko/simple-frontend
COPY header.Dockerfile footer.Dockerfile /
```

- As the name implies, `header.Dockerfile` is prepended to the file in the
  user's build context. See below for an example:

```dockerfile
FROM ubuntu
```

- Similarly, `footer.Dockerfile` will be appended to the file in the user's
  build context. See below for an example:

```dockerfile
LABEL org.opencontainers.image.licenses="Apache-2.0"
```

- Build the frontend: `docker build -t ubuntu-dockerfile .`
- Our frontend is now ready to be used. To illustrate this, let's create a
  sample partial `Dockerfile`. Note that `# syntax=` is a special stanza that
  instructs BuildKit to use a custom frontend. This should match the tag
  specified by `-t` earlier.

```dockerfile
# syntax=ubuntu-dockerfile
RUN apt-get update && \
    apt-get install -y \
        git \
    && rm -rf /var/lib/apt/lists/*
```

- When this `Dockerfile` is built, our frontend will concatenate it with the
  header and footer we specified. This will result in the following
  `Dockerfile` being used:

```dockerfile
FROM ubuntu
RUN apt-get update && \
    apt-get install -y \
        git \
    && rm -rf /var/lib/apt/lists/*
LABEL org.opencontainers.image.licenses="Apache-2.0"
```

That's it! If you'd like to see a complete ready example, check out
[examples](examples) directory in this repository.

## Custom frontend that transforms the `Dockerfile`

[Previous method](#using-simple-frontend-base-image) is quite limited in terms
of what can be achieved. The approach discussed in this section enables you to
create custom frontends that execute arbitrary transformations (written in Go)
on user-supplied `Dockerfile`. It works by intercepting the _read_ operations
from BuildKit and returning the transformed `Dockerfile` when it's read by the
upstream [docker/dockerfile frontend](https://hub.docker.com/r/docker/dockerfile).
Unfortunately, the transformation _does not_ have access any files in the build
context beyond the `Dockerfile`. This constraint is here because of how [docker/dockerfile frontend](https://hub.docker.com/r/docker/dockerfile)
works.

However, this is still quite a powerful pattern since you can use any existing
metadata file as long as it supports comments that start with `#` (so that you
can specify `# syntax=` stanza).

Follow the steps below to create a frontend:

- Create a new Go project with a main file. This will be the entry point for
  your frontend. See below for an example:

```go
package main

func main() {
	if err := grpcclient.RunFromEnvironment(appcontext.Context(), build); err != nil {
		logrus.Errorf("fatal error: %+v", err)
		panic(err)
	}
}
```

- Next, we need to define the function to build the container image. This
  would look roughly similar to below.

```go
func build(ctx context.Context, c client.Client) (*client.Result, error) {
	transform := func(dockerfile []byte) ([]byte, error) {
        // Do something to dockerfile here
		return dockerfile, nil
	}
	if err := dtp.InjectDockerfileTransform(transform, c); err != nil {
		return nil, err
	}

	// Pass control to the upstream Dockerfile frontend
	return dockerfile.Build(ctx, c)
}
```

- Once your transformation is defined, your frontend needs to be packaged.
  Frontends are distributed as container images, so feel free to refer to
  online resources to learn how to package up a Go binary. Key thing to
  keep in mind is that the binary should be specified as the entry point of
  the frontend container image:

```dockerfile
ENTRYPOINT ["/my-frontend"]
```

- As previously, we need to build our frontend: `docker build -t my-dockerfile .`.
  Our frontend is now ready to be used.

And we're done! If you'd like to see a complete ready example, check out
[Dockerfile](Dockerfile) and [cmd/simple-frontend/main.go](cmd/simple-frontend/main.go)
in this repository. These together define the frontend discussed in the first
section. Feel free to also refer to [docs of the dtp package](https://pkg.go.dev/github.com/EricHripko/buildkit-fdk@v0.1.0/pkg/dtp)
for additional guidance (dtp stands for Dockerfile Transforming Proxy).

## Custom frontend from scratch

If your use case cannot be satisfied with any of the approaches discussed
above, don't worry - you can still create a completely custom BuildKit
frontend. Similarly to before, you'll need to create a new Go project in order
to implement your frontend. The key differences will be in your _build_
function. Instead of passing control to the a different frontend, we'll be
implementing ours from scratch. The starter template for this can be seen
below:

```go
package main

const (
	keyMultiPlatform = "multi-platform"
)

func Build(ctx context.Context, c client.Client) (*client.Result, error) {
	return BuildWithService(ctx, c, cib.NewService(ctx, c))
}

func BuildWithService(ctx context.Context, c client.Client, svc cib.Service) (*client.Result, error) {
	opts := svc.GetOpts()

	// Identify target platforms
	targetPlatforms, err := svc.GetTargetPlatforms()
	if err != nil {
		return nil, err
	}
	exportMap := len(targetPlatforms) > 1
	if v := opts[keyMultiPlatform]; v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return nil, errors.Errorf("invalid boolean value %s", v)
		}
		if !b && exportMap {
			return nil, errors.Errorf("returning multiple target plaforms is not allowed")
		}
		exportMap = b
	}
	expPlatforms := &exptypes.Platforms{
		Platforms: make([]exptypes.Platform, len(targetPlatforms)),
	}

	// Build an image for each platform
	res := client.NewResult()
	eg, ctx := errgroup.WithContext(ctx)
	for i, tp := range targetPlatforms {
		func(i int, tp *specs.Platform) {
			eg.Go(func() error {
				// Declare the build process for your frontend
				// You'll need to create an LLB state and solve it to get
				// to get a ref(erence). This is then passed to BuildKit
				// to export as a container image.

				if !exportMap {
					res.AddMeta(exptypes.ExporterImageConfigKey, config)
					res.SetRef(ref)
				} else {
					p := platforms.DefaultSpec()
					if tp != nil {
						p = *tp
					}

					k := platforms.Format(p)
					res.AddMeta(fmt.Sprintf("%s/%s", exptypes.ExporterImageConfigKey, k), config)
					res.AddRef(k, ref)
					expPlatforms.Platforms[i] = exptypes.Platform{
						ID:       k,
						Platform: p,
					}
				}
				return nil
			})
		}(i, tp)
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// Export image(s)
	if exportMap {
		dt, err := json.Marshal(expPlatforms)
		if err != nil {
			return nil, err
		}
		res.AddMeta(exptypes.ExporterPlatformsKey, dt)
	}
	return res, nil
}
```

More flexibility comes with a great jump in complexity. To avoid boilerplate,
we can rely on [cib package](https://pkg.go.dev/github.com/EricHripko/buildkit-fdk@v0.1.0/pkg/cib).
It provides various primites for getting the build configuration, build
inputs, constructing LLB graphs, solving them and iterating the produced
artefacts. cib stands for Container Image Build.

Similarly, you'll need to package and build your frontend before it can be
used. If you'd like to see a complete example, check out the following custom
frontends using this approach:

- [pack.yaml](https://github.com/EricHripko/pack.yaml) - an opinionated
  frontend that builds the source code based on the YAML configuration
  provided.
- [cnbp](https://github.com/EricHripko/cnbp) - a proof-of-concept frontend
  that builds the source code using [Cloud Native Buildpacks](https://buildpacks.io/).
