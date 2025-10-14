# Quick Start Guide

## Installation Complete!

All dependencies have been installed and the project has been successfully built.

## Installed Dependencies

- **oapi-codegen** v2.5.0 - OpenAPI code generator
- **Gin** v1.11.0 - HTTP web framework
- **oapi-codegen/runtime** v1.1.2 - Runtime types and utilities
- **google/uuid** v1.6.0 - UUID generation

## Project Status

- Generated API code: `pkg/api/api.gen.go`
- Server binary: `build/server`
- Module name: `oapi-codegen-layout`

## Running the Server

### Option 1: Using Make
```bash
make run
```

### Option 2: Using the Binary
```bash
./build/server
```

### Option 3: Using Go Run
```bash
go run ./cmd/server/main.go
```

The server will start on `http://localhost:8080`

## Testing the API

### Health Check
```bash
curl http://localhost:8080/api/v1/health
```

Expected response:
```json
{
  "status": "ok",
  "timestamp": "2025-10-14T15:07:00Z"
}
```

### List Users
```bash
curl http://localhost:8080/api/v1/users
```

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "name": "John Doe"
  }'
```

### Get User by ID
```bash
# Replace {userId} with a valid UUID
curl http://localhost:8080/api/v1/users/{userId}
```

### Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/{userId} \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "name": "New Name"
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/{userId}
```

## Development Workflow

1. **Modify the API**: Edit `api/openapi.yaml`
2. **Regenerate Code**: Run `make generate`
3. **Update Handlers**: Implement new endpoints in `internal/handlers/`
4. **Build**: Run `make build`
5. **Test**: Run `make test`

## Key Files

- `api/openapi.yaml` - OpenAPI 3.0 specification (source of truth)
- `api/oapi-codegen.yaml` - Code generation configuration
- `pkg/api/api.gen.go` - Generated API code (DO NOT EDIT)
- `internal/handlers/users.go` - Handler implementation
- `cmd/server/main.go` - Server entry point
- `Makefile` - Build automation

## Important Notes

1. **Never edit generated files** (`*.gen.go`) - they will be overwritten
2. **UUID handling**: The generated code uses `openapi_types.UUID` from `github.com/oapi-codegen/runtime/types`
3. **Validation**: Request validation is handled automatically by the generated middleware
4. **Type safety**: All request/response types are strongly typed based on the OpenAPI spec

## Next Steps

1. Customize the handlers in `internal/handlers/users.go`
2. Add database integration
3. Add authentication/authorization
4. Add logging and monitoring
5. Add unit and integration tests
6. Configure environment-specific settings in `configs/`

## Common Make Commands

```bash
make help          # Show all available commands
make generate      # Generate API code from OpenAPI spec
make build         # Build the application
make run           # Run the application
make test          # Run tests with coverage
make clean         # Clean build artifacts
make deps          # Download dependencies
make fmt           # Format code
```

## Troubleshooting

### Port Already in Use
If port 8080 is already in use, you can change it in `cmd/server/main.go`:
```go
port := ":8080"  // Change this to your desired port
```

### Regenerate Code After OpenAPI Changes
Always run `make generate` after modifying the OpenAPI spec to update the generated code.

### Type Mismatches
Make sure handler method signatures match the generated `ServerInterface` in `pkg/api/api.gen.go`.
