package service

import (
	"golang-chapter-37/implem-gin-gonic/model"
	"golang-chapter-37/implem-gin-gonic/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type customerService struct {
	repo     repository.CustomerRepository
	validate *validator.Validate
}

type CustomerService interface {
	GetAllCustomers() ([]model.Customer, error)
	CreateCustomer(customer *model.Customer) error
}

func NewCustomerService(db *gorm.DB) CustomerService {
	return &customerService{
		repo:     repository.NewCustomerRepository(db),
		validate: validator.New(),
	}
}

func (s *customerService) GetAllCustomers() ([]model.Customer, error) {
	return s.repo.FindAll()
}

func (s *customerService) CreateCustomer(customer *model.Customer) error {
	if err := s.validate.Struct(customer); err != nil {
		return err
	}
	return s.repo.Save(customer)
}
