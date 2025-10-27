# Simple Decomposed Key-Value Store

This repository contains a simple decomposed Key-Value store implemented as two services:
- **gRPC Service**: kv-service - a backend Key-Value store service (using Go and gRPC) that handles storage operations in-memory.
- **REST Service**: rest-service - a JSON REST API that serves as the public interface (using Go and gorilla/mux).

The services communicate over gRPC. The REST service forwards requests to the gRPC service.

## Features
- **REST API Endpoints**:
   - `POST /kv`: Store a key-value pair (e.g., `{"key": "foo", "value": "bar"}`).
   - `GET /kv/:key`: Retrieve a value by key.
   - `DELETE /kv/:key`: Delete a key.
- **gRPC Service**: Handles Put, Get, and Delete operations with typed values.
- **Modular Design**: Clean separation of concerns for scalability and maintenance.
- **Dockerized**: Services run in separate containers with Docker Compose orchestration.
- **Extensible**: Easy to add new value types or swap storage backends.

## Project Structure
The project follows a clean, modular architecture that separates concerns between service entry points, business logic, communication layers, and deployment configuration. It is designed for clarity, scalability, and easy maintenance, allowing independent development and testing of each service.

### Layer Overview
- **cmd/**: Contains entry points for each service (`kv-service` and `rest-service`).
- **internal/app/**: Houses application logic. Each subdirectory (`kv-service`, `rest-service`) represents an independent service with its own server, handler, and logic layers.
- **proto/**: Defines and generates gRPC communication code between services.
- **docker/**: Provides Dockerfiles for isolated builds of each service.
- **docker-compose.yml**: Simplifies local orchestration of both services.
- **pkg/**: Reserved for utility or shared packages (currently empty).

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

4. **Error cases**:
- Invalid JSON in POST: 400 Bad Request.
- Get/Delete non-existent key: 404 Not Found.

### Automated Testing
No full test suite provided due to time constraints, but basic unit tests are in each service:
- Run `go test ./...` in `grpc-service` and `rest-service` directories.
- Tests cover basic Set/Get/Delete and gRPC stubs.

## Extensibility Notes
- **Add Functionality**: Extend `proto/kv.proto` with new RPCs (e.g., ListKeys), implement in gRPC server, and add endpoints in REST.
- **Change Transport**: The gRPC client in REST is designed to be modular, allowing the transport layer to be replaced with another protocol (such as HTTP) in the future if needed by swapping the client logic in `rest-service/kv_client.go`.
- **Storage**: Currently in-memory map; replace `store` in `grpc-service/server.go` with Redis, DB, etc.
- **Error Handling**: Centralized in both services; easy to add logging/monitoring.
- **Data Types**: The current KVStore implementation supports only `string` keys and values for simplicity. In future versions, this can be extended to handle arbitrary data types (e.g., JSON objects, integers, or binary data) by updating the protobuf definitions and storage logic.

