package main

import (
	"fmt"
	"log"
	"net/http"
	"oapi-codegen-layout/internal/config"
	"oapi-codegen-layout/internal/database"
	"oapi-codegen-layout/internal/router"
)

func main() {
	// Load configuration
	cfg, err := config.Load("")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Starting application in %s mode", cfg.Server.Mode)

	// Initialize database connection
	db, err := database.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup router with all routes and middleware
	r := router.Setup(&cfg.Server, db)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
