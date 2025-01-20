package repository

import (
	"Feature-Based/internal/auth"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *auth.User) error
	FindByEmail(email string) (*auth.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *auth.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*auth.User, error) {
	var user auth.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
