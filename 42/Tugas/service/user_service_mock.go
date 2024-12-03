package service

import (
	"project-voucher-team3/models"

	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUser(id uint) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
