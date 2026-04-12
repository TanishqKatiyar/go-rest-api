# Go Modular REST API

A highly performant, lightweight RESTful API built natively in Go using standard library `net/http`.

## Architecture
- Standalone Go binary (no heavy framework overhead)
- Native JSON serialization/deserialization
- Mock in-memory database configuration

## Running the API
```bash
go run main.go
```
API runs exclusively on `http://localhost:8080`.

### Routes
* `GET /api/v1/health` - Healthcheck node status
* `GET /api/v1/users` - Retrieve all registered administrative users.
