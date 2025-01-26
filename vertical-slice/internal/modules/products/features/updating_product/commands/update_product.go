package commands

type UpdateProductCommand struct {
	ID          uint
	Name        string
	Description string
	Price       float64
}

func NewUpdateProductCommand(id uint, name, description string, price float64) *UpdateProductCommand {
	return &UpdateProductCommand{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
