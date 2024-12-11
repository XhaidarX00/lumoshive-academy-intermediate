package services

import (
	"project_auth_jwt/repositories"

	"go.uber.org/zap"
)

type Service struct {
	User   UserServiceInterface
	JWT    JWTServiceInterface
	Auth   AuthServiceInterface
	Report ReportServiceInterface
}

func NewService(repo repositories.Repository, log *zap.Logger, jwtService JWTServiceInterface, secretKey string) Service {
	return Service{
		User:   NewUserService(repo, log),
		JWT:    NewJWTService(secretKey),
		Auth:   NewAuthService(repo, jwtService, log),
		Report: NewReportService(repo, log),
	}
}
