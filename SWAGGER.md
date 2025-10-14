# Swagger UI Setup

## Access Swagger UI

After starting the server with `make run` or `./build/server`, access the Swagger UI at:

**http://localhost:8080/swagger/index.html**

## What You Can Do with Swagger UI

### 1. Interactive API Documentation
- Browse all available endpoints
- View request/response schemas
- See required and optional parameters
- Check response status codes and error messages

### 2. Test Endpoints
- Click "Try it out" on any endpoint
- Fill in parameters and request body
- Click "Execute" to send real requests
- View the response directly in the browser

### 3. View OpenAPI Specification
- Access the raw OpenAPI spec at: http://localhost:8080/openapi.json
- Use this for API client generation
- Import into other tools like Postman

## Example Usage

### Testing the Health Endpoint

1. Open http://localhost:8080/swagger/index.html
2. Find the `GET /api/v1/health` endpoint
3. Click "Try it out"
4. Click "Execute"
5. See the response:
   ```json
   {
     "status": "ok",
     "timestamp": "2025-10-14T15:07:00Z"
   }
   ```

### Testing User Creation

1. Find the `POST /api/v1/users` endpoint
2. Click "Try it out"
3. Edit the request body:
   ```json
   {
     "email": "test@example.com",
     "name": "Test User"
   }
   ```
4. Click "Execute"
5. View the created user with generated UUID

### Testing User Retrieval

1. Copy a user UUID from the create response
2. Find the `GET /api/v1/users/{userId}` endpoint
3. Click "Try it out"
4. Paste the UUID in the userId field
5. Click "Execute"
6. View the user details

## Available Routes

### Swagger Routes
- `GET /swagger/index.html` - Swagger UI interface
- `GET /openapi.json` - OpenAPI specification JSON

### API Routes (via Swagger)
- `GET /api/v1/health` - Health check
- `GET /api/v1/users` - List users (with optional limit query param)
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/{userId}` - Get user by ID
- `PUT /api/v1/users/{userId}` - Update user
- `DELETE /api/v1/users/{userId}` - Delete user

## Implementation Details

### Files Added
- `internal/handlers/swagger.go` - Handler to serve OpenAPI JSON
- Updated `cmd/server/main.go` - Registered Swagger routes

### Dependencies Added
- `github.com/swaggo/gin-swagger` - Swagger middleware for Gin
- `github.com/swaggo/files` - Embedded Swagger UI files

### How It Works

1. **OpenAPI Spec**: The spec is embedded in `pkg/api/api.gen.go` by oapi-codegen
2. **JSON Endpoint**: `handlers.GetSwaggerJSON()` serves the spec at `/openapi.json`
3. **Swagger UI**: `gin-swagger` wraps the Swagger UI and points it to our JSON endpoint
4. **Interactive**: Swagger UI reads the JSON and creates the interactive interface

## Customization

### Change Swagger UI URL
Edit `cmd/server/main.go`:
```go
router.GET("/docs/*any", ginSwagger.WrapHandler(
    swaggerFiles.Handler,
    ginSwagger.URL("/openapi.json"),
))
```

### Add Authentication to Swagger
```go
router.GET("/swagger/*any",
    gin.BasicAuth(gin.Accounts{"admin": "password"}),
    ginSwagger.WrapHandler(swaggerFiles.Handler))
```

### Disable in Production
```go
if os.Getenv("ENV") != "production" {
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

## Troubleshooting

### Swagger UI Not Loading
- Ensure server is running: `make run`
- Check URL: http://localhost:8080/swagger/index.html
- Check server logs for errors

### OpenAPI Spec Not Found
- Verify `/openapi.json` returns JSON
- Regenerate API code: `make generate`
- Rebuild: `make build`

### UI Shows Empty or Error
- Clear browser cache
- Check browser console for JavaScript errors
- Verify OpenAPI spec is valid: http://localhost:8080/openapi.json

## Benefits

1. **No Manual Documentation**: Documentation auto-generated from OpenAPI spec
2. **Always Up-to-Date**: Regenerate code to update docs automatically
3. **Interactive Testing**: Test endpoints without curl or Postman
4. **Client Generation**: Use the spec to generate API clients in any language
5. **Contract First**: OpenAPI spec serves as the contract between frontend and backend

## Next Steps

1. Start the server: `make run`
2. Open Swagger UI: http://localhost:8080/swagger/index.html
3. Try out the endpoints interactively
4. Modify `api/openapi.yaml` to add new endpoints
5. Run `make generate` to update
6. Restart server to see changes in Swagger UI
