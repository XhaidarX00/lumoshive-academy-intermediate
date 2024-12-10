package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type JWTServiceInterface interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}

type JWTService struct {
	SecretKey []byte
	Log       *zap.Logger
}

func NewJWTService(SecretKey string) JWTServiceInterface {
	return &JWTService{
		SecretKey: []byte(SecretKey),
	}
}

func (s *JWTService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.SecretKey)
}

func (s *JWTService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.SecretKey, nil
	})

	if err != nil {
		s.Log.Error("Failed to validate token", zap.Error(err))
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	s.Log.Error("Failed to ", zap.Error(err))
	return nil, errors.New("invalid token")
}
