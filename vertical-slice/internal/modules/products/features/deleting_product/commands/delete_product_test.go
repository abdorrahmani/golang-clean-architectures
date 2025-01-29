package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vertical-slice/internal/modules/products/models"
	"vertical-slice/internal/modules/products/repository"
	"vertical-slice/internal/pkg/database"
)

func TestDeleteProductHandler(t *testing.T) {
	db := database.InitTestDB()
	repo := repository.NewProductRepository(db)
	handler := NewDeleteProductHandler(repo)

	t.Run("Delete Existing Product", func(t *testing.T) {
		product := &models.Product{Name: "Test Product", Price: 50.0, Description: "Test Description"}
		repo.CreateProduct(product)

		cmd := NewDeleteProductCommand(product.ID)
		err := handler.Handle(cmd)

		assert.NoError(t, err)

		_, err = repo.GetProductByID(product.ID)
		assert.Error(t, err)
	})
}
