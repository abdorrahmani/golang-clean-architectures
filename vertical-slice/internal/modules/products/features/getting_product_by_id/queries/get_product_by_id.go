package queries

type GetProductByIdQuery struct {
	ID uint `json:"id"`
}

func NewGetProductByIDQuery(id uint) *GetProductByIdQuery {
	return &GetProductByIdQuery{
		ID: id,
	}
}
