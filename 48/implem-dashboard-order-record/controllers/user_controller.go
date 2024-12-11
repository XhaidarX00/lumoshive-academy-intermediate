// controllers/user_controller.go
package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"project_auth_jwt/helper"
	"project_auth_jwt/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserController struct {
	Svc services.Service
	Log *zap.Logger
}

func NewUserController(userService services.Service, log *zap.Logger) *UserController {
	return &UserController{
		Svc: userService,
		Log: log,
	}
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.Svc.User.GetAllUsers()
	if err != nil {
		ctrl.Log.Error("Failed get users", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed get users", http.StatusInternalServerError)
		c.Abort()
		return
	}

	// Hindari menampilkan password
	var userResponses []gin.H
	for _, user := range users {
		userResponses = append(userResponses, gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"last_login": user.LastLogin,
		})
	}

	helper.ResponseOK(c, userResponses, "Get users succesfully", http.StatusOK)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	// Parse ID dari parameter URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		ctrl.Log.Error("Invalid user ID", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid user ID", http.StatusBadRequest)
		c.Abort()
		return
	}

	user, err := ctrl.Svc.User.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctrl.Log.Error("User not found", zap.Error(err))
			helper.ResponseError(c, err.Error(), "User not found", http.StatusNotFound)
			c.Abort()
		} else {
			ctrl.Log.Error("Failed get user", zap.Error(err))
			helper.ResponseError(c, err.Error(), "Failed get user", http.StatusInternalServerError)
			c.Abort()
		}

		return
	}

	// Hindari menampilkan password
	userResponse := gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"last_login": user.LastLogin,
	}

	helper.ResponseOK(c, userResponse, "Get users succesfully", http.StatusOK)
}

func (ctrl *UserController) GetUserByEmail(c *gin.Context) {
	// Parse ID dari parameter URL
	email := c.Param("email")
	if email == "" {
		ctrl.Log.Error("Invalid user email", zap.String("message", "email should be not nil"))
		helper.ResponseError(c, "email should be not nil", "Invalid user email", http.StatusBadRequest)
		c.Abort()
		return
	}

	user, err := ctrl.Svc.User.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctrl.Log.Error("User not found", zap.Error(err))
			helper.ResponseError(c, err.Error(), "User not found", http.StatusNotFound)
			c.Abort()
		} else {
			ctrl.Log.Error("Failed get user", zap.Error(err))
			helper.ResponseError(c, err.Error(), "Failed get user", http.StatusInternalServerError)
			c.Abort()
		}

		return
	}

	// Hindari menampilkan password
	userResponse := gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"last_login": user.LastLogin,
	}

	helper.ResponseOK(c, userResponse, "Get users succesfully", http.StatusOK)
}
