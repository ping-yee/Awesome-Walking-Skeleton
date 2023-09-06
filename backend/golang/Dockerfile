FROM golang:1.15-alpine as build-env
# All these steps will be cached
RUN mkdir -p /github.com/taimoor99/three-tier-golang
WORKDIR /github.com/taimoor99/three-tier-golang
COPY go.mod .
# <- COPY go.mod and go.sum files to the workspace
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN go build -o three-tier-golang ./cmd/three-tier-golang

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./three-tier-golang"]

