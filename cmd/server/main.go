package main

import (
	"log"
	"net/http"
	"oapi-codegen-layout/internal/database"
	"oapi-codegen-layout/internal/router"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup router with all routes and middleware
	r := router.Setup(db)

	// Start server
	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
