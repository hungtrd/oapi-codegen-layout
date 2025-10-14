package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"oapi-codegen-layout/pkg/api"
)

// GetSwaggerJSON serves the OpenAPI specification as JSON
func GetSwaggerJSON(c *gin.Context) {
	swagger, err := api.GetSwagger()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to load OpenAPI spec",
		})
		return
	}

	c.JSON(http.StatusOK, swagger)
}
