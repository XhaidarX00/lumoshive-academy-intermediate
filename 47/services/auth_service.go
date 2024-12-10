// services/auth_service.go
package services

import (
	"errors"
	"time"

	"project_auth_jwt/helper"
	"project_auth_jwt/models"
	"project_auth_jwt/repositories"

	"go.uber.org/zap"
)

type AuthServiceInterface interface {
	Register(user models.RegisterRequest) (models.User, error)
	Login(email, password string) (string, error)
}

type AuthService struct {
	JWTService JWTServiceInterface
	Repo       repositories.Repository
	Log        *zap.Logger
}

func NewAuthService(userRepo repositories.Repository, jwtService JWTServiceInterface, log *zap.Logger) AuthServiceInterface {
	return &AuthService{
		JWTService: jwtService,
		Repo:       userRepo,
		Log:        log,
	}
}

func (s *AuthService) Register(regis models.RegisterRequest) (models.User, error) {
	// Hash password
	hashedPassword, err := helper.HashPassword(regis.Password)
	if err != nil {
		s.Log.Error("Failed to register", zap.Error(err))
		return models.User{}, err
	}

	var user = models.User{
		Name:     regis.Name,
		Email:    regis.Email,
		Password: hashedPassword,
		Role:     regis.Role,
	}
	err = s.Repo.User.Create(&user)
	if err != nil {
		s.Log.Error("Failed to Register", zap.Error(err))
		return models.User{}, err
	}

	result, err := s.Repo.User.FindByEmail(regis.Email)
	if err != nil {
		s.Log.Error("Failed to Register", zap.Error(err))
		return models.User{}, err
	}

	return *result, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.User.FindByEmail(email)
	if err != nil {
		s.Log.Error("Failed to login", zap.Error(err))
		return "", errors.New("user not found")
	}

	// Cek password
	err = helper.CompareHashAndPassword(user.Password, password)
	if err != nil {
		s.Log.Error("Failed to login", zap.Error(err))
		return "", errors.New("invalid password")
	}

	// Generate JWT
	token, err := s.JWTService.GenerateToken(uint(user.ID))
	if err != nil {
		s.Log.Error("Failed to login", zap.Error(err))
		return "", err
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	err = s.Repo.User.Update(user)
	if err != nil {
		s.Log.Error("Failed to login", zap.Error(err))
		return "", err
	}

	return token, nil
}
