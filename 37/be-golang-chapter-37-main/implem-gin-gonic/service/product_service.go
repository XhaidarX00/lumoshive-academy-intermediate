package service

import (
	"golang-chapter-37/implem-gin-gonic/model"
	"golang-chapter-37/implem-gin-gonic/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type productService struct {
	repo     repository.ProductRepository
	validate *validator.Validate
}

type ProductService interface {
	GetAllProducts() ([]model.Product, error)
	CreateProduct(product *model.Product) error
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{
		repo:     repository.NewProductRepository(db),
		validate: validator.New(),
	}
}

func (s *productService) GetAllProducts() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) CreateProduct(product *model.Product) error {
	if err := s.validate.Struct(product); err != nil {
		return err
	}
	return s.repo.Save(product)
}
