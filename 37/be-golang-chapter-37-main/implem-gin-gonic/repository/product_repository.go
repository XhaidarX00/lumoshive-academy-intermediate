package repository

import (
	"golang-chapter-37/implem-gin-gonic/model"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	Save(product *model.Product) error
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) Save(product *model.Product) error {
	return r.db.Create(product).Error
}
