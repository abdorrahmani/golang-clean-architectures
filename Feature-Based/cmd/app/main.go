package main

import (
	"Feature-Based/configs"
	"Feature-Based/internal/post"
	"Feature-Based/internal/post/handler"
	"Feature-Based/internal/post/repository"
	"Feature-Based/internal/post/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := configs.LoadConfig()

	db := configs.ConnectDB(config.DatabaseDSN)

	db.AutoMigrate(&post.Post{})

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := gin.Default()

	postRoutes := r.Group("/api/v1/posts")
	{
		postRoutes.GET("/", postHandler.GetAllPosts)
		postRoutes.POST("/", postHandler.CreatePost)
		postRoutes.GET("/:id", postHandler.GetPostByID)
		postRoutes.PUT("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
	}
	// Start the server
	log.Printf("Server running at http://localhost:%s", config.ServerPort)
	err := r.Run(":" + config.ServerPort)
	if err != nil {
		log.Fatalf("Faild to start server : %v", err)
	}
}
