# oapi-codegen-layout

A Go project following the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) for API code generation using Gin framework.

## Project Structure

```
.
├── api/                    # OpenAPI/Swagger specs and code generation config
│   ├── openapi.yaml       # OpenAPI specification
│   └── oapi-codegen.yaml  # Code generation configuration
├── build/                  # Build output directory
├── cmd/                    # Main applications
│   └── server/            # API server application
│       └── main.go
├── configs/               # Configuration files
├── docs/                  # Design and user documents
├── internal/              # Private application code
│   ├── handlers/          # HTTP handlers
│   └── models/            # Data models
├── pkg/                   # Public libraries
│   └── api/              # Generated API code (*.gen.go)
├── scripts/               # Build, install, and analysis scripts
├── Makefile              # Build automation
├── go.mod                # Go module definition
└── README.md             # This file
```

## Prerequisites

- Go 1.25 or higher
- Make (optional, for using Makefile commands)
- Docker and Docker Compose (for containerized deployment)

## Getting Started

### 1. Install Dependencies

```bash
# Install oapi-codegen tool
make install-tools

# Download Go dependencies
make deps
```

Or manually:

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
go mod download
```

### 2. Generate API Code

Generate Go code from the OpenAPI specification:

```bash
make generate
```

This will create `pkg/api/api.gen.go` with:
- Type definitions for all models
- Gin server interface
- Request/response bindings
- Embedded OpenAPI spec

### 3. Build the Application

```bash
make build
```

The binary will be created at `build/server`.

### 4. Run the Application

```bash
make run
```

Or run the built binary:

```bash
./build/server
```

The API server will start on `http://localhost:8080`.

## Swagger UI

Once the server is running, you can access the interactive Swagger UI at:

**http://localhost:8080/swagger/index.html**

The Swagger UI provides:
- Interactive API documentation
- Try out endpoints directly from the browser
- View request/response schemas
- Test API with different parameters

You can also access the raw OpenAPI spec JSON at:
- **http://localhost:8080/openapi.json**

## Available Endpoints

- `GET /api/v1/health` - Health check endpoint
- `GET /api/v1/users` - List all users
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users/{userId}` - Get user by ID
- `PUT /api/v1/users/{userId}` - Update user
- `DELETE /api/v1/users/{userId}` - Delete user

## Testing the API

### Health Check

```bash
curl http://localhost:8080/api/v1/health
```

### List Users

```bash
curl http://localhost:8080/api/v1/users
```

### Create User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","name":"Test User"}'
```

### Get User by ID

```bash
curl http://localhost:8080/api/v1/users/{userId}
```

## Docker Deployment

### Quick Start with Docker Compose

Start the entire application stack (API + MySQL) with a single command:

```bash
make docker-compose-up
```

This will:
- Start a MySQL 8.0 container
- Start the API server container
- Automatically create and migrate the database schema
- Expose the API on http://localhost:8080
- Expose MySQL on localhost:3306

Access the application:
- **API**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **OpenAPI Spec**: http://localhost:8080/openapi.json

### Docker Commands

```bash
# Build and start all services
make docker-compose-up

# View logs from all services
make docker-compose-logs

# Stop all services
make docker-compose-down

# Rebuild services after code changes
make docker-compose-build

# Restart services
make docker-compose-restart

# Clean up all Docker resources (containers, images, volumes)
make docker-clean
```

### Environment Configuration

The application uses environment variables for database configuration. Create a `.env` file in the project root:

```env
# Database Configuration
DB_USER=appuser
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=example_db
```

**Note**: The `.env` file is already included in `.gitignore`. Never commit sensitive credentials to version control.

### Database

The Docker setup includes:
- **MySQL 8.0** container with persistent volume
- **Automatic schema migration** on startup using GORM
- **Health checks** to ensure MySQL is ready before starting the API

Database models are defined in `internal/models/`:
- `User` - User entity with UUID, email, name, and timestamps
- `Product` - Product entity with UUID, name, description, price, category, stock, and timestamps

Both models support soft deletes (records are marked as deleted but not actually removed).

### Testing the Dockerized API

Once the services are running, test the API:

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Create a user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","name":"Test User"}'

# List users
curl http://localhost:8080/api/v1/users

# Create a product
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","price":1299.99,"category":"Electronics","stock":10}'

# List products
curl http://localhost:8080/api/v1/products
```

### Docker Architecture

The docker-compose setup includes:
- **API Service**: Go application built with multi-stage Dockerfile
- **MySQL Service**: MySQL 8.0 with persistent volume storage
- **Network**: Bridge network for service communication
- **Health Checks**: Ensures MySQL is ready before starting API

## Development

### Code Generation

The project uses oapi-codegen to generate server code from OpenAPI specs. The configuration is in `api/oapi-codegen.yaml`:

```yaml
package: api
generate:
  gin-server: true
  models: true
  embedded-spec: true
output: pkg/api/api.gen.go
```

After modifying `api/openapi.yaml`, regenerate the code:

```bash
make generate
```

### Implementing Handlers

Handlers are located in `internal/handlers/`. To implement a new endpoint:

1. Add the endpoint to `api/openapi.yaml`
2. Run `make generate` to update the generated code
3. Implement the handler method in `internal/handlers/`

The handler must implement the `ServerInterface` from the generated code.

### Adding New Features

1. Update the OpenAPI spec in `api/openapi.yaml`
2. Regenerate code: `make generate`
3. Implement handlers in `internal/handlers/`
4. Add business logic in `internal/` packages
5. Update tests

## Makefile Commands

### Development Commands
- `make help` - Display all available commands
- `make install-tools` - Install required development tools
- `make generate` - Generate API code from OpenAPI spec
- `make build` - Build the application
- `make run` - Run the application
- `make clean` - Clean build artifacts
- `make test` - Run tests with coverage
- `make deps` - Download and tidy dependencies
- `make fmt` - Format code
- `make lint` - Run linter

### Docker Commands
- `make docker-compose-up` - Start all services with docker-compose
- `make docker-compose-down` - Stop all services
- `make docker-compose-logs` - Show logs from all services
- `make docker-compose-build` - Build docker-compose services
- `make docker-compose-restart` - Restart all services
- `make docker-clean` - Remove all containers, images, and volumes
- `make docker-build` - Build Docker image
- `make docker-run` - Run Docker container

## Project Layout

This project follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

- `/cmd` - Main applications for this project
- `/internal` - Private application and library code
- `/pkg` - Library code that's ok to use by external applications
- `/api` - OpenAPI/Swagger specs, JSON schema files, protocol definition files
- `/configs` - Configuration file templates or default configs
- `/docs` - Design and user documents
- `/scripts` - Scripts to perform various build, install, analysis, etc operations
- `/build` - Packaging and Continuous Integration

## Code Generation Details

The `oapi-codegen` tool generates:

1. **Models** - Go structs for all OpenAPI schemas
2. **Server Interface** - Interface that handlers must implement
3. **Request/Response Bindings** - Parameter parsing and validation
4. **Route Registration** - Helper to register all routes

Example generated interface:

```go
type ServerInterface interface {
    GetHealth(c *gin.Context)
    ListUsers(c *gin.Context, params ListUsersParams)
    CreateUser(c *gin.Context)
    GetUserById(c *gin.Context, userId string)
    UpdateUser(c *gin.Context, userId string)
    DeleteUser(c *gin.Context, userId string)
}
```

## Configuration

Configuration can be managed through:
- Environment variables
- Configuration files in `/configs`
- Command-line flags

## Testing

Run tests with coverage:

```bash
make test
```

This generates:
- `coverage.out` - Coverage data
- `coverage.html` - HTML coverage report

## Best Practices

1. **Never edit generated files** (`*.gen.go`) - they will be overwritten
2. **Update OpenAPI spec first** - let code generation drive your API
3. **Keep handlers thin** - business logic goes in `internal/` packages
4. **Use dependency injection** - pass dependencies to handler constructors
5. **Validate at API boundary** - OpenAPI validation catches issues early

## Contributing

1. Update the OpenAPI spec
2. Run `make generate`
3. Implement the handlers
4. Add tests
5. Run `make fmt` and `make lint`
6. Submit a pull request

## License

MIT
