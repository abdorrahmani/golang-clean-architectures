package main

import (
	"Feature-Based/configs"
	"Feature-Based/internal/auth"
	authHandler "Feature-Based/internal/auth/handler"
	authMiddleware "Feature-Based/internal/auth/middleware"
	authRepository "Feature-Based/internal/auth/repository"
	authService "Feature-Based/internal/auth/service"
	"Feature-Based/internal/post"
	postHandler "Feature-Based/internal/post/handler"
	postRepository "Feature-Based/internal/post/repository"
	postService "Feature-Based/internal/post/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := configs.LoadConfig()

	db := configs.ConnectDB(config.DatabaseDSN)

	db.AutoMigrate(&auth.User{}, &post.Post{})

	jwtService := auth.NewJWTService(config.JWTSecret)
	userRepo := authRepository.NewUserRepository(db)
	authService := authService.NewAuthService(userRepo)
	authHandler := authHandler.NewAuthHandler(authService, jwtService)
	_ = authMiddleware.NewAuthMiddleware(jwtService)

	postRepo := postRepository.NewPostRepository(db)
	postService := postService.NewPostService(postRepo)
	postHandler := postHandler.NewPostHandler(postService)

	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

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
