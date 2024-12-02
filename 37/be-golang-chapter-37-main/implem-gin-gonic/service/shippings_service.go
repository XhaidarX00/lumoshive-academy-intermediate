package service

import (
	"golang-chapter-37/implem-gin-gonic/model"
	"golang-chapter-37/implem-gin-gonic/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type shippingService struct {
	repo     repository.ShippingRepository
	validate *validator.Validate
}

type ShippingService interface {
	GetShippingServices(data *[]model.ShippingService) error
}

func NewShippingService(db *gorm.DB, cs repository.CustomerRepository, slr repository.SellerRepository) ShippingService {
	return &shippingService{
		repo:     repository.NewShippingRepository(db, cs, slr),
		validate: validator.New(),
	}
}

func (s *shippingService) GetShippingServices(data *[]model.ShippingService) error {
	err := s.repo.GetShippingRepo(data)
	if err != nil {
		return err
	}

	return nil
}
