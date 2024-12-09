package controllers

import (
	"project_auth_jwt/config"
	"project_auth_jwt/database"
	"project_auth_jwt/services"

	"go.uber.org/zap"
)

type Controller struct {
	User UserController
	Auth AuthController
}

func NewController(service services.Service, logger *zap.Logger, cacher database.Cacher, config config.Configuration) *Controller {
	return &Controller{
		User: *NewUserController(service),
		Auth: *NewAuthController(service),
	}
}
