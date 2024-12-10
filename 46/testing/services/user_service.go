// services/user_service.go (baru)
package services

import (
	"project_auth_jwt/models"
	"project_auth_jwt/repositories"

	"go.uber.org/zap"
)

// type UserServiceInterface interface {
// 	GetAllUsers() ([]models.User, error)
// 	GetUserByID(id uint) (*models.User, error)
// }

type UserService struct {
	Repo repositories.Repository
	Log  *zap.Logger
}

// func NewUserService(userRepo repositories.Repository, log *zap.Logger) UserServiceInterface {
// 	return &UserService{
// 		Repo: userRepo,
// 		Log:  log,
// 	}
// }

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.User.FindAll()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.User.FindByID(id)
}
