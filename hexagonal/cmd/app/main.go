package main

import (
	"github.com/gin-gonic/gin"
	"hexagonal/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(":", cfg.App.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
