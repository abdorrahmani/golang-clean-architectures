package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectToDB establishes a connection to the MySQL database using GORM.
// It constructs the Data Source Name (DSN) from the configuration and opens a connection.
// If the connection fails, it logs a fatal error and exits the application.
func ConnectToDB(cfg *Config) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Could not connect to the database: %v", err)
	}
	DB = db
	log.Println("✅ Database connected successfully")
}

// RunMigration runs the database migrations using GORM's AutoMigrate feature.
// It will automatically create or update the database schema based on the models defined in the application.
func RunMigration() {
	err := DB.AutoMigrate()
	if err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	log.Println("✅ Database migration completed successfully")
}
