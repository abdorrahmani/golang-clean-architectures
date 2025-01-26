package queries

import (
	"errors"
	"vertical-slice/internal/modules/products/models"
	"vertical-slice/internal/modules/products/repository"
)

type GetProductByIDHandler struct {
	ProductRepo repository.ProductRepository
}

func NewGetProductByIDHandler(repo repository.ProductRepository) *GetProductByIDHandler {
	return &GetProductByIDHandler{
		ProductRepo: repo,
	}
}

func (h *GetProductByIDHandler) Handle(query *GetProductByIdQuery) (*models.Product, error) {
	product, err := h.ProductRepo.GetProductByID(query.ID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
