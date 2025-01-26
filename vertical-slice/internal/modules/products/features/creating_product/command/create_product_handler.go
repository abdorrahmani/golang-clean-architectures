package command

import (
	"log"
	"vertical-slice/internal/modules/products/models"
	"vertical-slice/internal/modules/products/repository"
)

type CreateProductHandler struct {
	ProductRepo repository.ProductRepository
}

func NewCreateProductHandler(repo repository.ProductRepository) *CreateProductHandler {
	return &CreateProductHandler{
		ProductRepo: repo,
	}
}

func (h *CreateProductHandler) Handle(commend *CreateProductCommend) (*models.Product, error) {
	product := &models.Product{
		Name:        commend.Name,
		Description: commend.Description,
		Price:       commend.Price,
	}

	createProduct, err := h.ProductRepo.CreateProduct(product)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return nil, err
	}

	return createProduct, nil
}
