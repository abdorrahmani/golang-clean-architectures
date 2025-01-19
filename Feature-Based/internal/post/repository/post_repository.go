package repository

import (
	"Feature-Based/internal/post"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *post.Post) error
	FindAll() ([]post.Post, error)
	FindByID(id uint) (*post.Post, error)
	Update(post *post.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) FindAll() ([]post.Post, error) {
	var posts []post.Post

	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) Create(post *post.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindByID(id uint) (*post.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) Update(post *post.Post) error {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}
