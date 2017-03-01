# See https://blog.golang.org/docker for more on deploying Go servers with Docker

# Start from a Debian image with the specified version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.7.5

# Working directory - mapped to local source through Docker Compose
RUN mkdir -p /go/src/treadmill
WORKDIR /go/src/treadmill

# Install gvt for lightweight dependency management
RUN go get github.com/FiloSottile/gvt

# Forward logs to docker log collector
RUN mkdir -p /var/log/treadmill \
	&& ln -sf /dev/stdout /var/log/treadmill/treadmill.log
