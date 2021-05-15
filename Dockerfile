FROM golang:1.15 as build
WORKDIR /go/src/app
ENV CGO_ENABLED=0
COPY . .
RUN go install -v ./...
RUN touch /go/bin/header.Dockerfile /go/bin/footer.Dockerfile

FROM scratch
LABEL moby.buildkit.frontend.network.none="true"
COPY --from=build \
    /go/bin/header.Dockerfile \
    /go/bin/footer.Dockerfile \
    /go/bin/simple-frontend \
    /
ENTRYPOINT ["/simple-frontend"]
LABEL \
    org.opencontainers.image.authors="Eric Hripko" \
    org.opencontainers.image.documentation="https://github.com/EricHripko/buildkit-fdk" \
    org.opencontainers.image.source="https://github.com/EricHripko/buildkit-fdk.git"\
    org.opencontainers.image.licenses="Apache-2.0" \
    org.opencontainers.image.title="Simple BuildKit frontend" \
    org.opencontainers.image.description="Frontend that concatenates user-specified Dockerfile with the pre-configured header and footer"
