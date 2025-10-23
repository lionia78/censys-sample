FROM golang:1.24-alpine AS builder
WORKDIR /root/app

# Install dependencies
RUN apk add --no-cache git protoc protobuf

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o kv-service ./cmd/kv-service

FROM alpine:3.18
WORKDIR /root/app

# Copy binary from builder stage
COPY --from=builder /root/app/kv-service /usr/local/bin/kv-service

EXPOSE 50051
ENTRYPOINT ["/usr/local/bin/kv-service"]
