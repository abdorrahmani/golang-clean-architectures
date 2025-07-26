package http

import (
	"errors"
	"hexagonal/internal/app"
	"hexagonal/internal/config"
	"hexagonal/internal/domain"

	"gorm.io/gorm"
)

type GormPostRepository struct{}

func NewGormPostRepository() app.PostRepository {
	return &GormPostRepository{}
}

func (r *GormPostRepository) Create(post *domain.Post) error {
	return config.DB.Create(post).Error
}

func (r *GormPostRepository) GetByID(id int) (*domain.Post, error) {
	var post domain.Post
	result := config.DB.First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &post, result.Error
}

func (r *GormPostRepository) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	result := config.DB.Find(&posts)
	return posts, result.Error
}

func (r *GormPostRepository) Update(post *domain.Post) error {
	return config.DB.Save(post).Error
}

func (r *GormPostRepository) Delete(id int) error {
	return config.DB.Delete(&domain.Post{}, id).Error
}
