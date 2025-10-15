package main

import (
	"log"
	"net/http"
	"oapi-codegen-layout/internal/database"
	"oapi-codegen-layout/internal/handlers"
	"oapi-codegen-layout/pkg/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Create unified handler that implements all API endpoints
	handler := handlers.NewHandler(db)

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
