package service

import (
	"Feature-Based/internal/auth"
	"Feature-Based/internal/auth/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(name, email, password string) error
	Login(email, password string) (uint, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &auth.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(user)
}

func (s *authService) Login(email, password string) (uint, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return 0, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return 0, errors.New("invalid credentials")
	}

	return user.ID, nil
}
