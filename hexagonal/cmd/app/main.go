package main

import (
	adaptershttp "hexagonal/internal/adapters/http"
	"hexagonal/internal/app"
	"hexagonal/internal/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	config.ConnectToDB(cfg)
	config.RunMigration()

	repo := adaptershttp.NewGormPostRepository()
	service := app.NewPostService(repo)

	r := gin.Default()
	adaptershttp.RegisterPostRoutes(r, service)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(":" + cfg.App.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
