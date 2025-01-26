package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vertical-slice/internal/modules/products/features/getting_all_products/queries"
	"vertical-slice/internal/modules/products/repository"
	"vertical-slice/internal/pkg/database"
)

func GetAllProductsEndpoint(c *gin.Context) {
	handler := queries.NewGetAllProductsHandler(repository.NewProductRepository(database.DB))
	query := queries.NewGetAllProductsQuery()

	products, err := handler.Handle(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
