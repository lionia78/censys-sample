# Simple Decomposed Key-Value Store

This repository contains a simple decomposed Key-Value store implemented as two services:
- **REST Service**: A JSON REST API that serves as the public interface (using Go and gorilla/mux).
- **gRPC Service**: A backend Key-Value store service (using Go and gRPC) that handles storage operations in-memory.

The services communicate over gRPC. The REST service forwards requests to the gRPC service.

## Features
- Store a value at a given key (POST /kv).
- Retrieve the value for a given key (GET /kv/:key).
- Delete a given key (DELETE /kv/:key).
- In-memory storage (using a Go map for simplicity; supports concurrent access and extensible to other backends).
- Built as two separate Docker containers.
- gRPC for inter-service communication.

## Prerequisites
- Docker
- Docker Compose (for easy running)
- Go 1.24 (if building locally)

## Building and Running

### Using Docker Compose
1. Clone the repository.
2. Run `docker-compose up --build`.
   - This builds and starts both services.
   - gRPC service listens on port 50051 (internal).
   - REST service listens on port 8080 (exposed).

## Testing Instructions

### Manual Testing
Use curl or Postman to test the REST API (assuming running on localhost:8080).

1. **Set a key-value**:

```sh
  curl -X POST http://localhost:8080/kv/foo -H "Content-Type: application/json" -d '{"value": "bar"}'
```
Expected: HTTP Code 204 (No Content)

2. **Get a value**:

```sh
curl -X GET http://localhost:8080/kv/foo
```
Expected: `{"value": "bar"}`

3. **Delete a key**:

```sh
curl -X DELETE http://localhost:8080/kv/foo
```

Expected: HTTP Code 204 (No Content)

4. Error cases:
- Invalid JSON in POST: 400 Bad Request.
- Get/Delete non-existent key: 404 Not Found.

### Automated Testing
No full test suite provided due to time constraints, but basic unit tests are in each service:
- Run `go test ./...` in `grpc-service` and `rest-service` directories.
- Tests cover basic Set/Get/Delete and gRPC stubs.

## Extensibility Notes
- **Add Functionality**: Extend `proto/kv.proto` with new RPCs (e.g., ListKeys), implement in gRPC server, and add endpoints in REST.
- **Change Transport**: The gRPC client in REST is modular; replace with HTTP or other protocols by swapping the client logic in `rest-service/kv_client.go`.
- **Storage**: Currently in-memory map; replace `store` in `grpc-service/server.go` with Redis, DB, etc.
- **Error Handling**: Centralized in both services; easy to add logging/monitoring.
- **Data Types**: The current KVStore implementation supports only `string` keys and values for simplicity. In future versions, this can be extended to handle arbitrary data types (e.g., JSON objects, integers, or binary data) by updating the protobuf definitions and storage logic.

