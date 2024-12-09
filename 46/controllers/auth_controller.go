// controllers/auth_controller.go
package controllers

import (
	"project_auth_jwt/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.Service
}

func NewAuthController(authService services.Service) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.authService.Auth.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.authService.Auth.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
