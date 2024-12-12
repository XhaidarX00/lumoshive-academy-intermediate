package userrepository

import (
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
	UserRepo UserRepository
}

func (m *MockUserRepository) Create(userInput models.User) error {
	args := m.Called(userInput)
	return args.Error(0)
}

// GetByEmail mocks the GetByEmail method
func (m *MockUserRepository) GetByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// UpdatePassword mocks the UpdatePassword method
func (m *MockUserRepository) UpdatePassword(resetPasswordInput models.LoginRequest) error {
	args := m.Called(resetPasswordInput)
	return args.Error(0)
}

func (m *MockUserRepository) CountCustomer() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
