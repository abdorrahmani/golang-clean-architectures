package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vertical-slice/config"
	"vertical-slice/internal/modules/products"
	"vertical-slice/internal/pkg/database"
	"vertical-slice/internal/shared/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	database.ConnectDatabase(cfg)
	app.RunMigration()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	products.RegisterRoutes(r)
	err = r.Run(":" + cfg.App.Port)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
