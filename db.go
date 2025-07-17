package main

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"

	"github.com/lukegrady1/to-do-list-api/models"
)

var DB *gorm.DB

func initDB() {
	driver := os.Getenv("DB_DRIVER")
	source := os.Getenv("DB_SOURCE")

	var err error
	switch driver {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(source), &gorm.Config{})
	default:
		log.Fatalf("unsupported DB_DRIVER: %s", driver)
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate our models
	if err := DB.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("auto-migration failed: %v", err)
	}
}
