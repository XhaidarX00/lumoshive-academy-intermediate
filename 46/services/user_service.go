// services/user_service.go (baru)
package services

import (
	"project_auth_jwt/models"
	"project_auth_jwt/repositories"

	"go.uber.org/zap"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type userService struct {
	Repo repositories.Repository
	Log  *zap.Logger
}

func NewUserService(userRepo repositories.Repository, log *zap.Logger) UserServiceInterface {
	return &userService{
		Repo: userRepo,
		Log:  log,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.Repo.User.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.User.FindByID(id)
}
