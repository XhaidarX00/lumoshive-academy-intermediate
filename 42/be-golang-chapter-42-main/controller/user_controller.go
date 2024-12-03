package controller

import (
	"golang-chapter-42/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewControllerUser() *Controller {
	return &Controller{}
}

// Get user all endpoint
// @Summary Get user all
// @Description Get details of all registered users.
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.ResponseWithData "All User"
// @Failure 404 {object} utils.Response "User not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router  /users [get]
func (c *Controller) GetAll(ctx *gin.Context) {

	// response
	ctx.JSON(200, gin.H{"status_code": 200, "message": "show all data users"})
}

// Get user by ID endpoint
// @Summary Get user by ID
// @Description Get details of a user by their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Security BasicAuth
// @Param   id     path     int     true  "User ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} utils.ResponseWithData "All User"
// @Failure 404 {object} utils.Response "User not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router  /users/{id} [get]
func (c *Controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	// response

	// Dummy implementation for example
	if id == "1" {
		ctx.JSON(200, gin.H{"status_code": 200, "message": "show data user by id"})
	} else {
		ctx.JSON(404, gin.H{"status_code": 404, "message": "User not found"})
	}
}

// store data user endpoint
// @Summary Add a new user
// @Description Add a new user with the provided JSON data.
// @Tags users
// @Security BasicAuth
// @Param user body model.User true "User data"
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} utils.Response "All User"
// @Failure 404 {object} utils.Response "User not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /users [post]
func (c *Controller) Store(ctx *gin.Context) {
	var newUser model.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	ctx.JSON(201, gin.H{"status_code": 201, "message": "success store data user"})
}

// update user endpoint
// @Summary Update user by ID
// @Description Update the details of a user by their ID.
// @Tags users
// @Security BasicAuth
// @Param id path int true "User ID"
// @Param user body model.User true "User object that needs to be updated"
// @Success 200 {object} utils.Response "All User"
// @Failure 404 {object} utils.Response "User not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /users/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	// Dummy implementation
	if id == "" {
		ctx.JSON(400, gin.H{"status_code": 400, "message": "Invalid updated user"})
	} else {
		ctx.JSON(200, gin.H{"status_code": 200, "message": "Success updated data user"})
	}
}

// delete user endpoint
// @Summary Delete user by ID
// @Description Delete a user by their ID.
// @Tags users
// @Security BasicAuth
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Success 200 {object} utils.Response "All User"
// @Failure 404 {object} utils.Response "User not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /users/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	// Dummy implementation:
	if id == "1" {
		ctx.JSON(200, gin.H{"status_code": 200, "message": "Success deleted data user"})
	} else {
		ctx.JSON(400, gin.H{"status_code": 400, "message": "Invalid deleted user"})
	}
}

// search data user by name endpoint
// @Summary Get users by name
// @Description Get a list of users that match the provided name.
// @Tags users
// @Security BasicAuth
// @Param name query string false "User Name"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} utils.ResponseWithData "All User"
// @Failure 404 {object} utils.Response "User not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /seacrh_users [get]
func (c *Controller) SearchByName(ctx *gin.Context) {
	name := ctx.Query("name")
	// Dummy implementation
	if name == "" {
		ctx.JSON(404, gin.H{"status_code": 404, "message": "User not found"})
	} else {
		ctx.JSON(200, gin.H{"status_code": 200, "message": "show data user by name"})
	}
}

// upload photo user endpoint
// @Summary Upload a file
// @Description Upload a file to the server.
// @Tags files
// @Security BasicAuth
// @Param file formData file true "File to upload"
// @Success 200 {object} utils.Response "All User"
// @Failure 404 {object} utils.Response "User not found"
// @Failure 500 {object} utils.Response "Internal server error"
// @Router /upload [post]
func (c *Controller) UploadPhoto(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"status_code": 400, "error": "Failed to upload file"})
		return
	}
	// Save the file
	ctx.SaveUploadedFile(file, file.Filename)
	ctx.JSON(200, gin.H{"status_code": 200, "message": "File uploaded successfully", "filename": file.Filename})
}
