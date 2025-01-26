package commands

import (
	"vertical-slice/internal/modules/products/models"
	"vertical-slice/internal/modules/products/repository"
)

type UpdateProductHandler struct {
	repo repository.ProductRepository
}

func NewUpdateProductHandler(repo repository.ProductRepository) *UpdateProductHandler {
	return &UpdateProductHandler{repo: repo}
}

func (h *UpdateProductHandler) Handle(cmd *UpdateProductCommand) (*models.Product, error) {
	product, err := h.repo.GetProductByID(cmd.ID)
	if err != nil {
		return nil, err
	}

	product.Name = cmd.Name
	product.Description = cmd.Description
	product.Price = cmd.Price
	productUpdated, err := h.repo.UpdateProduct(product)
	return productUpdated, nil
}
