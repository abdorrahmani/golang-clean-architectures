package commands

type DeleteProductCommand struct {
	ID uint
}

func NewDeleteProductCommand(id uint) *DeleteProductCommand {
	return &DeleteProductCommand{ID: id}
}
