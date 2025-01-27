package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vertical-slice/internal/modules/products/repository"
	"vertical-slice/internal/pkg/database"
)

func TestCreateProductHandler(t *testing.T) {
	db := database.InitTestDB()
	repo := repository.NewProductRepository(db)
	handler := NewCreateProductHandler(repo)

	t.Run("Successful Product Creation", func(t *testing.T) {
		cmd := NewCreateProductCommend("Test Product", "Test Description", 99.99)
		product, err := handler.Handle(cmd)

		assert.NoError(t, err)
		assert.NotZero(t, product.ID)

		p, _ := repo.GetProductByID(product.ID)
		assert.Equal(t, "Test Product", p.Name)
	})
}
