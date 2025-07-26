package app

import (
	"hexagonal/internal/domain"
)

type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *domain.Post) error {
	return s.repo.Create(post)
}

func (s *PostService) GetPostByID(id int) (*domain.Post, error) {
	return s.repo.GetByID(id)
}

func (s *PostService) GetAllPosts() ([]domain.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) UpdatePost(post *domain.Post) error {
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.Delete(id)
}
