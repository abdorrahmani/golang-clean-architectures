package queries

import (
	"vertical-slice/internal/modules/products/models"
	"vertical-slice/internal/modules/products/repository"
)

type GetAllProductsHandler struct {
	repo repository.ProductRepository
}

func NewGetAllProductsHandler(repo repository.ProductRepository) *GetAllProductsHandler {
	return &GetAllProductsHandler{repo: repo}
}

func (h *GetAllProductsHandler) Handle(query *GetAllProductsQuery) ([]models.Product, error) {
	return h.repo.GetAllProducts()
}
