package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"oapi-codegen-layout/internal/handlers"
	"oapi-codegen-layout/pkg/api"
)

func main() {
	// Create Gin router
	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Create unified handler that implements all API endpoints
	handler := handlers.NewHandler()

	// Swagger endpoints - serve OpenAPI spec at a different path to avoid conflicts
	router.GET("/openapi.json", handlers.GetSwaggerJSON)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/openapi.json")))

	// Register routes with the API version prefix
	apiGroup := router.Group("/api/v1")
	api.RegisterHandlers(apiGroup, handler)

	// Start server
	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
