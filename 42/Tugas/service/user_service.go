package service

import (
	"project-voucher-team3/models"
	"project-voucher-team3/repository"
)

type UserService interface {
	CreateUser(user models.User) error
	GetUser(id int) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUser(id int) (models.User, error) {
	return s.repo.GetUser(id)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
