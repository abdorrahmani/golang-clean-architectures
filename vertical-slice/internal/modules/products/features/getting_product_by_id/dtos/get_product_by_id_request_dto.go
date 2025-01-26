package dtos

type GetProductByIdRequestDto struct {
	ID uint `uri:"id" binding:"required"`
}

func NewGetRequestByIDRequestDTO(id uint) *GetProductByIdRequestDto {
	return &GetProductByIdRequestDto{
		ID: id,
	}
}
