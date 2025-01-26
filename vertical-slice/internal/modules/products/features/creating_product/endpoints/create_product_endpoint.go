package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vertical-slice/internal/modules/products/features/creating_product/command"
	"vertical-slice/internal/modules/products/features/creating_product/dtos"
	"vertical-slice/internal/modules/products/repository"
	"vertical-slice/internal/pkg/database"
)

func CreateProductEndpoint(c *gin.Context) {
	var request dtos.CreateProductRequestDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	commend := command.NewCreateProductCommend(request.Name, request.Description, request.Price)

	handler := command.NewCreateProductHandler(repository.NewProductRepository(database.DB))
	createProduct, err := handler.Handle(commend)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}
	response := dtos.NewCreateProductResponseDTO(createProduct.ID, createProduct.Name,
		createProduct.Description, createProduct.Price)

	c.JSON(http.StatusOK, response)
}
