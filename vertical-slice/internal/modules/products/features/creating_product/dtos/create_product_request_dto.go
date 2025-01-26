package dtos

type CreateProductRequestDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewCreateProductRequestDTO(name, description string, price float64) *CreateProductRequestDTO {
	return &CreateProductRequestDTO{
		Name:        name,
		Description: description,
		Price:       price,
	}
}
