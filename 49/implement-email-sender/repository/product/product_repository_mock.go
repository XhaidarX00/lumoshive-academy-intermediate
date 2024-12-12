package productrepository

import (
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (p *ProductRepositoryMock) Create(productInput models.Product) error {
	panic("unimplemented")
}

// Delete implements ProductRepositoryMock.
func (p *ProductRepositoryMock) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements ProductRepositoryMock.
func (p *ProductRepositoryMock) GetAll() ([]models.Product, error) {
	panic("unimplemented")
}

// GetByID implements ProductRepositoryMock.
func (p *ProductRepositoryMock) GetByID(id int) (*models.Product, error) {
	panic("unimplemented")
}

// Update implements ProductRepository.
func (p *ProductRepositoryMock) Update(productInput models.Product) error {
	panic("unimplemented")
}
