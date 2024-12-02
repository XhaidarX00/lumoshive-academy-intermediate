package service

import (
	"project-voucher-team3/models"
	"project-voucher-team3/repository"
)

type AuthService interface {
	Login(user *models.User, auth *models.Auth) error
}

type authService struct {
	Repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (a authService) Login(user *models.User, auth *models.Auth) error {
	return a.Repo.Login(user, auth)
}
