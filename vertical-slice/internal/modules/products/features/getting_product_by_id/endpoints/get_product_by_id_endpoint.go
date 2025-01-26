package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vertical-slice/internal/modules/products/features/getting_product_by_id/dtos"
	"vertical-slice/internal/modules/products/features/getting_product_by_id/queries"
	"vertical-slice/internal/modules/products/repository"
	"vertical-slice/internal/pkg/database"
)

func GetProductByIDEndpoint(c *gin.Context) {
	var request dtos.GetProductByIdRequestDto
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	query := queries.NewGetProductByIDQuery(request.ID)

	handler := queries.NewGetProductByIDHandler(repository.NewProductRepository(database.DB))
	product, err := handler.Handle(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	response := dtos.NewGetProductByIDResponseDTO(product.ID, product.Name, product.Description, product.Price, product.CreatedAt, product.UpdatedAt)
	c.JSON(http.StatusOK, response)
}
