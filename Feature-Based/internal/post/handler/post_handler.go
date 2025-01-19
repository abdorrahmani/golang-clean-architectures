package handler

import (
	"Feature-Based/internal/post"
	"Feature-Based/internal/post/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{service: postService}
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var newPost post.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreatePost(&newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		return
	}

	c.JSON(http.StatusCreated, newPost)
}
