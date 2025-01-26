package app

import (
	"log"
	"vertical-slice/internal/modules/products/models"
	"vertical-slice/internal/pkg/database"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Failed to run migartions: %v", err)
	}
	log.Println("Database migration completed.")
}
