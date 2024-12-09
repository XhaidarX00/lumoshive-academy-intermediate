package controller

import (
	"net/http"
	"project-voucher-team3/models"
	"project-voucher-team3/service"
	"project-voucher-team3/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	service service.AuthService
	logger  *zap.Logger
}

func NewAuthController(service service.AuthService, logger *zap.Logger) *AuthController {
	return &AuthController{service, logger}
}

// Login user endpoint
// @Summary Login User
// @Description Login users and create token temporary.
// @Tags Users
// @Param user body models.Login true "Login object that needs to Login auth"
// @Success 200 {object} utils.ResponseSuccess "All User"
// @Failure 404 {object} utils.ResponseFailed "User not found"
// @Failure 500 {object} utils.ResponseFailed "Internal server error"
// @Router /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ctrl.logger.Error("Failed to bind user data", zap.Error(err))
		utils.ResponseError(c, "Login Error", err.Error(), http.StatusBadRequest)
		return
	}

	var auth models.Auth
	err := ctrl.service.Login(&user, &auth)
	if err != nil {
		utils.ResponseError(c, "Login Error", err.Error(), http.StatusUnauthorized)
		c.Abort()
		return
	}

	utils.ResponseOK(c, auth, "Login Succesfully")
}