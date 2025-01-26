package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vertical-slice/internal/modules/products/features/updating_product/commands"
	"vertical-slice/internal/modules/products/repository"
	"vertical-slice/internal/pkg/database"
)

func UpdateProductEndpoint(c *gin.Context) {
	var request struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	handler := commands.NewUpdateProductHandler(repository.NewProductRepository(database.DB))
	cmd := commands.NewUpdateProductCommand(
		uint(id),
		request.Name,
		request.Description,
		request.Price,
	)

	if _, err := handler.Handle(cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
