package userservice

import (
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
	UserService UserService
}

func (m *MockService) User() UserService {
	args := m.Called()
	return args.Get(0).(UserService)
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(registerReq models.RegisterRequest) error {
	args := m.Called(registerReq)
	return args.Error(0)
}

// Login provides a mock function with given arguments
func (m *MockUserService) Login(input models.LoginRequest) (*models.User, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// CheckUserEmail provides a mock function with given arguments
func (m *MockUserService) CheckUserEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// ResetUserPassword provides a mock function with given arguments
func (m *MockUserService) ResetUserPassword(input models.LoginRequest) error {
	args := m.Called(input)
	return args.Error(0)
}
