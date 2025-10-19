package handlers

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"oapi-codegen-layout/pkg/api/health"
	"oapi-codegen-layout/pkg/api/products"
	"oapi-codegen-layout/pkg/api/users"
)

// GetSwaggerJSON serves the combined OpenAPI specification as JSON
func GetSwaggerJSON(c *gin.Context) {
	// Get specs from all handlers
	usersSwagger, err := users.GetSwagger()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load users OpenAPI spec",
		})
		return
	}

	productsSwagger, err := products.GetSwagger()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load products OpenAPI spec",
		})
		return
	}

	healthSwagger, err := health.GetSwagger()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load health OpenAPI spec",
		})
		return
	}

	// Create combined spec
	combined := &openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:       "Example API",
			Description: "Example API using oapi-codegen with Gin - Split Handlers",
			Version:     "1.0.0",
		},
		Servers: openapi3.Servers{
			{
				URL:         "http://localhost:8080/api/v1",
				Description: "Development server",
			},
		},
		Paths:      openapi3.NewPaths(),
		Components: &openapi3.Components{
			Schemas: make(openapi3.Schemas),
		},
	}

	// Merge paths from all handlers
	if usersSwagger.Paths != nil {
		for path, pathItem := range usersSwagger.Paths.Map() {
			combined.Paths.Set(path, pathItem)
		}
	}

	if productsSwagger.Paths != nil {
		for path, pathItem := range productsSwagger.Paths.Map() {
			combined.Paths.Set(path, pathItem)
		}
	}

	if healthSwagger.Paths != nil {
		for path, pathItem := range healthSwagger.Paths.Map() {
			combined.Paths.Set(path, pathItem)
		}
	}

	// Merge schemas from all handlers
	if usersSwagger.Components != nil && usersSwagger.Components.Schemas != nil {
		for name, schema := range usersSwagger.Components.Schemas {
			combined.Components.Schemas[name] = schema
		}
	}

	if productsSwagger.Components != nil && productsSwagger.Components.Schemas != nil {
		for name, schema := range productsSwagger.Components.Schemas {
			combined.Components.Schemas[name] = schema
		}
	}

	if healthSwagger.Components != nil && healthSwagger.Components.Schemas != nil {
		for name, schema := range healthSwagger.Components.Schemas {
			combined.Components.Schemas[name] = schema
		}
	}

	c.JSON(http.StatusOK, combined)
}
