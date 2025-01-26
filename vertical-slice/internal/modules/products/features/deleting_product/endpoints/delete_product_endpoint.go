package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vertical-slice/internal/modules/products/features/deleting_product/commands"
	"vertical-slice/internal/modules/products/repository"
	"vertical-slice/internal/pkg/database"
)

func DeleteProductEndpoint(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product ID",
		})
		return
	}

	handler := commands.NewDeleteProductHandler(repository.NewProductRepository(database.DB))
	cmd := commands.NewDeleteProductCommand(uint(id))
	if err := handler.Handle(cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
