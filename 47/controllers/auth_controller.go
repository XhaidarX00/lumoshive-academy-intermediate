// controllers/auth_controller.go
package controllers

import (
	"net/http"
	"project_auth_jwt/helper"
	"project_auth_jwt/models"
	"project_auth_jwt/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthService services.Service
	Log         *zap.Logger
}

func NewAuthController(authService services.Service, log *zap.Logger) *AuthController {
	return &AuthController{
		AuthService: authService,
		Log:         log,
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		c.Abort()
		return
	}

	user, err := ctrl.AuthService.Auth.Register(req)
	if err != nil {
		ctrl.Log.Error("Failed Register", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed Register", http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, user, "User registered successfully", http.StatusCreated)
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		c.Abort()
		return
	}

	token, err := ctrl.AuthService.Auth.Login(req.Email, req.Password)
	if err != nil {
		ctrl.Log.Error("Token invalid", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Token invalid", http.StatusUnauthorized)
		c.Abort()
		return
	}

	req.Token = token
	helper.ResponseOK(c, req, "Login successfully", http.StatusOK)
}
