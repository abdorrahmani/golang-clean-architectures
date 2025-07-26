package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

func LoadConfig() (*Config, error) {
	var config Config
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.App.Name = os.Getenv("APP_NAME")
	config.App.Port = os.Getenv("APP_PORT")
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Name = os.Getenv("DB_NAME")

	return &config, nil
}
