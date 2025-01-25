package repository

import (
	"errors"
	"gorm.io/gorm"
	"vertical-slice/internal/modules/products/models"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(id uint) error
}

type GormProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &GormProductRepository{
		db: db,
	}
}

func (r *GormProductRepository) GetAllProducts() ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r *GormProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *GormProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}

func (r *GormProductRepository) UpdateProduct(product *models.Product) (*models.Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *GormProductRepository) DeleteProduct(id uint) error {
	if err := r.db.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
