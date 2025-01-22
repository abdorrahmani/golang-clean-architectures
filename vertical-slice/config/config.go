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
	var c Config
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	c.App.Name = os.Getenv("APP_NAME")
	c.App.Port = os.Getenv("APP_PORT")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.Name = os.Getenv("DB_NAME")

	return &c, nil
}
