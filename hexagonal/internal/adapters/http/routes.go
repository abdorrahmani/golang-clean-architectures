package http

import (
	"github.com/gin-gonic/gin"
	"hexagonal/internal/app"
)

func RegisterPostRoutes(r *gin.Engine, service *app.PostService) {
	handler := NewPostHandler(service)
	posts := r.Group("/posts")
	{
		posts.POST("", handler.CreatePost)
		posts.GET("", handler.GetAllPosts)
		posts.GET(":id", handler.GetPostByID)
		posts.PUT(":id", handler.UpdatePost)
		posts.DELETE(":id", handler.DeletePost)
	}
}
