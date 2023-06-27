FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.16 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
# Build
ENV TREE 3
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o /demoapp
# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080
# Run

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
WORKDIR /
COPY --from=builder /demoapp /demoapp
ENTRYPOINT ["/demoapp"]
