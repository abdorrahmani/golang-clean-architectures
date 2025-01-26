package dtos

import "time"

type GetProductByIdResponseDto struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewGetProductByIDResponseDTO(id uint, name, description string, price float64, createdAt, updatedAt time.Time) *GetProductByIdResponseDto {
	return &GetProductByIdResponseDto{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
