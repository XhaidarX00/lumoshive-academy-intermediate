// services/auth_service.go
package services

import (
	"errors"
	"time"

	"project_auth_jwt/models"
	"project_auth_jwt/repositories"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(username, email, password string) error
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

func (s *AuthService) Register(username, email, password string) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.Repo.User.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.User.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT
	token, err := s.JWTService.GenerateToken(uint(user.ID))
	if err != nil {
		return "", err
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	err = s.Repo.User.Update(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
