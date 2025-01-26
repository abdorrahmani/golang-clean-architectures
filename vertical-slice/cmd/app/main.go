package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vertical-slice/config"
	createEndpoint "vertical-slice/internal/modules/products/features/creating_product/endpoints"
	"vertical-slice/internal/modules/products/features/getting_all_products/endpoints"
	getEndpoint "vertical-slice/internal/modules/products/features/getting_product_by_id/endpoints"
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

	r.GET("/products", endpoints.GetAllProductsEndpoint)
	r.POST("/products", createEndpoint.CreateProductEndpoint)
	r.GET("/products/:id", getEndpoint.GetProductByIDEndpoint)

	err = r.Run(":" + cfg.App.Port)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
