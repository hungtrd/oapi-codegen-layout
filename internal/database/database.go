package database

import (
	"fmt"
	"log"
	"oapi-codegen-layout/internal/config"
	"oapi-codegen-layout/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB initializes the database connection and runs migrations
func InitDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	// Create DSN (Data Source Name) from config
	dsn := cfg.GetDSN()

	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate database schemas
	if err := db.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}
