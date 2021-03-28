FROM --platform=${BUILDPLATFORM} golang:1.16-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /out/go-git-staged .

FROM scratch AS bin-unix
COPY --from=build /out/go-git-staged /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/go-git-staged /go-git-staged.exe

FROM bin-${TARGETOS} AS bin
