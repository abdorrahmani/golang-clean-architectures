package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	ServerPort  string
	DatabaseDSN string
}

func LoadConfig() Config {
	dsn := "user:password@tcp(db:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	return Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseDSN: getEnv("DATABASE_DSN", dsn),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully!")
	return db
}
