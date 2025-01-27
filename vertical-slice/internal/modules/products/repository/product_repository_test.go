package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vertical-slice/internal/modules/products/models"
	"vertical-slice/internal/pkg/database"
)

func TestProductRepository(t *testing.T) {
	db := database.InitTestDB()
	repo := NewProductRepository(db)

	t.Run("Create Product", func(t *testing.T) {
		product := &models.Product{Name: "Test Product", Price: 100.0, Description: "test"}
		_, err := repo.CreateProduct(product)
		assert.NoError(t, err)
		assert.NotZero(t, product.ID)
	})

	t.Run("Get Product By ID", func(t *testing.T) {
		product := &models.Product{Name: "Another Product", Price: 200.0, Description: "test"}
		repo.CreateProduct(product)

		foundProduct, err := repo.GetProductByID(product.ID)
		assert.NoError(t, err)
		assert.Equal(t, product.Name, foundProduct.Name)
	})

	t.Run("Update Product", func(t *testing.T) {
		product := &models.Product{Name: "Old Product", Price: 300.0, Description: "test"}
		repo.CreateProduct(product)

		product.Name = "Updated Product"
		_, err := repo.UpdateProduct(product)
		assert.NoError(t, err)

		updatedProduct, _ := repo.GetProductByID(product.ID)
		assert.Equal(t, "Updated Product", updatedProduct.Name)
	})

	t.Run("Delete Product", func(t *testing.T) {
		product := &models.Product{Name: "To Be Deleted", Price: 50.0, Description: "test"}
		repo.CreateProduct(product)

		err := repo.DeleteProduct(product.ID)
		assert.NoError(t, err)

		deletedProduct, err := repo.GetProductByID(product.ID)
		assert.Error(t, err)
		assert.Nil(t, deletedProduct)
	})
}
