// controllers/user_controller.go
package controllers

import (
	"errors"
	"strconv"

	"project_auth_jwt/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Svc services.Service
}

func NewUserController(userService services.Service) *UserController {
	return &UserController{
		Svc: userService,
	}
}

func (ctrl *UserController) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	c.JSON(200, gin.H{"user_id": userID, "message": "Profile accessed"})
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.Svc.User.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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

	c.JSON(200, userResponses)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	// Parse ID dari parameter URL
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := ctrl.Svc.User.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	// Hindari menampilkan password
	c.JSON(200, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"last_login": user.LastLogin,
	})
}
