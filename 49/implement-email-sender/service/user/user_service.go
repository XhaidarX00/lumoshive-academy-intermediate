package userservice

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	"errors"

	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(userInput models.RegisterRequest) error
	Login(input models.LoginRequest) (*models.User, error)
	CheckUserEmail(email string) (*models.User, error)
	ResetUserPassword(input models.LoginRequest) error
}

type userService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CheckUserEmail implements UserService.
func (u *userService) CheckUserEmail(email string) (*models.User, error) {
	return u.Repo.User.GetByEmail(email)
}

// CreateUser implements UserService.
func (u *userService) CreateUser(userInput models.RegisterRequest) error {
	newUserInput := models.User{
		Email:    userInput.Email,
		Password: helper.HashPassword(userInput.Password),
		Role:     "staff",
		Name:     userInput.Name,
	}
	return u.Repo.User.Create(newUserInput)
}

// Login implements UserService.
func (u *userService) Login(input models.LoginRequest) (*models.User, error) {
	user, err := u.Repo.User.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	// Check if the user's password matches the input password
	if !helper.CheckPassword(input.Password, user.Password) {
		return nil, errors.New("invalid user password")
	}
	return user, nil
}

// ResetUserPassword implements UserService.
func (u *userService) ResetUserPassword(input models.LoginRequest) error {
	resetPassword := models.LoginRequest{
		Email:    input.Email,
		Password: helper.HashPassword(input.Password),
	}
	return u.Repo.User.UpdatePassword(resetPassword)
}

func NewUserService(repo repository.Repository, log *zap.Logger) UserService {
	return &userService{Repo: repo, Log: log}
}
