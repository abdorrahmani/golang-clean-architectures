package commands

import "vertical-slice/internal/modules/products/repository"

type DeleteProductHandler struct {
	repo repository.ProductRepository
}

func NewDeleteProductHandler(repo repository.ProductRepository) *DeleteProductHandler {
	return &DeleteProductHandler{repo: repo}
}

func (h *DeleteProductHandler) Handle(cmd *DeleteProductCommand) error {
	return h.repo.DeleteProduct(cmd.ID)
}
