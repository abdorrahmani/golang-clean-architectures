package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vertical-slice/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err = r.Run(":" + cfg.App.Port)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
