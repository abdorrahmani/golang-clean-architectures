package dtos

type CreateProductResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewCreateProductResponseDTO(id uint, name, description string, price float64) *CreateProductResponseDTO {
	return &CreateProductResponseDTO{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
