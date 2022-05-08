# Build the manager binary
FROM golang:1.18 as builder

# Make sure we use go modules
WORKDIR /vcluster

# Copy the Go Modules manifests
COPY . .

# Install dependencies
RUN go mod vendor

# Build cmd
RUN CGO_ENABLED=0 GO111MODULE=on go build -mod vendor -o /plugin ./cmd/main.go

# we use alpine for easier debugging
FROM alpine

# Set root path as working directory
WORKDIR /

COPY /tls /tls
COPY --from=builder /plugin .

COPY manifests/ /manifests/

ENTRYPOINT ["/plugin"]
