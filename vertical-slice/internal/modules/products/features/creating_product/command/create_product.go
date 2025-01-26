package command

type CreateProductCommend struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewCreateProductCommend(name, description string, price float64) *CreateProductCommend {
	return &CreateProductCommend{
		Name:        name,
		Description: description,
		Price:       price,
	}
}
