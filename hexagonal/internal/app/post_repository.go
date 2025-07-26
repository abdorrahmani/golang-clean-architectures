package app

import "hexagonal/internal/domain"

type PostRepository interface {
	Create(post *domain.Post) error
	GetByID(id int) (*domain.Post, error)
	GetAll() ([]domain.Post, error)
	Update(post *domain.Post) error
	Delete(id int) error
}
