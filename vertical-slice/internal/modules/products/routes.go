package products

import (
	"github.com/gin-gonic/gin"
	createEndpoint "vertical-slice/internal/modules/products/features/creating_product/endpoints"
	deleteEndpoint "vertical-slice/internal/modules/products/features/deleting_product/endpoints"
	getAllEndpoint "vertical-slice/internal/modules/products/features/getting_all_products/endpoints"
	getByIDEndpoint "vertical-slice/internal/modules/products/features/getting_product_by_id/endpoints"
	updateEndpoint "vertical-slice/internal/modules/products/features/updating_product/endpoints"
)

func RegisterRoutes(router *gin.Engine) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("", getAllEndpoint.GetAllProductsEndpoint)
		productRoutes.POST("", createEndpoint.CreateProductEndpoint)
		productRoutes.GET("/:id", getByIDEndpoint.GetProductByIDEndpoint)
		productRoutes.PUT("/:id", updateEndpoint.UpdateProductEndpoint)
		productRoutes.DELETE("/:id", deleteEndpoint.DeleteProductEndpoint)
	}
}
