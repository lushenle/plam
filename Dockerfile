# syntax=docker/dockerfile:1
# Build the binary
FROM --platform=$BUILDPLATFORM golang:bullseye as builder

WORKDIR /workspace

# Install upx for compress binary file
RUN apt update && apt install -y upx

# Copy the go source
COPY . .

#ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Build and compression
ARG TARGETARCH

RUN GOARCH=$TARGETARCH go build -a -installsuffix cgo -ldflags="-s -w" -o bin/server main.go
RUN upx bin/server

# build server
FROM alpine:3.17.2
WORKDIR /

COPY --from=builder /workspace/bin .

COPY start.sh /
COPY db/migration ./db/migration

ENV GIN_MODE=release

EXPOSE 8080/tcp

CMD ["/server"]
ENTRYPOINT ["/start.sh"]
