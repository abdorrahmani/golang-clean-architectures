package service

import (
	"Feature-Based/internal/post"
	"Feature-Based/internal/post/repository"
)

type PostService interface {
	CreatePost(post *post.Post) error
	GetAllPosts() ([]post.Post, error)
	GetPostByID(id uint) (*post.Post, error)
	UpdatePost(post *post.Post) error
	DeletePost(id uint) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(post *post.Post) error {
	return s.repo.Create(post)
}

func (s *postService) GetAllPosts() ([]post.Post, error) {
	return s.repo.FindAll()
}

func (s *postService) GetPostByID(id uint) (*post.Post, error) {
	return s.repo.FindByID(id)
}

func (s *postService) UpdatePost(post *post.Post) error {
	return s.repo.Update(post)
}

func (s *postService) DeletePost(id uint) error {
	return s.repo.Delete(id)
}
