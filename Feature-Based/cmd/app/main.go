package main

import (
	"Feature-Based/configs"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := configs.LoadConfig()

	db := configs.ConnectDB(config.DatabaseDSN)

	log.Println(db)

	r := gin.Default()

	// Start the server
	log.Printf("Server running at http://localhost:%s", config.ServerPort)
	err := r.Run(":" + config.ServerPort)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
