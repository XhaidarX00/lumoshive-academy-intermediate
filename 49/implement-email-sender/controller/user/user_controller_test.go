package usercontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"dashboard-ecommerce-team2/config"
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	userservice "dashboard-ecommerce-team2/service/user"
)

var _ userservice.UserService = (*userservice.MockUserService)(nil)

// Helper function to setup user controller with mock services
func setupUserController(mockService userservice.UserService) *UserController {
	mockLogger, _ := zap.NewProduction()
	mockConfig := config.Configuration{}
	mockCacher := database.Cacher{}

	return &UserController{
		Log:     mockLogger,
		Service: service.Service{User: mockService},
		Cacher:  mockCacher,
		Config:  mockConfig,
	}
}

// Helper function to setup test request context
func setupTest(t *testing.T, requestBody interface{}) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Prepare request body
	jsonBody, err := json.Marshal(requestBody)
	assert.NoError(t, err)
	c.Request, _ = http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	return w, c
}

func TestCreateUserController(t *testing.T) {
	// Subtest for Successful User Creation
	t.Run("Successful User Creation", func(t *testing.T) {
		requestBody := models.RegisterRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}

		expectedStatus := http.StatusCreated
		expectedMessage := "User created successfully"
		mockServiceError := error(nil) // no error, successful case

		// Initialize mock service
		mockUserService := new(userservice.MockUserService)
		mockUserService.On("CreateUser", requestBody).Return(mockServiceError).Once()

		// Setup the mock service and controller
		w, c := setupTest(t, requestBody)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.CreateUserController(c)

		// Assertions
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)

		// Verify that the mock service method was called
		mockUserService.AssertExpectations(t)
	})

	// Subtest for Invalid Request Body
	t.Run("Invalid Request Body", func(t *testing.T) {
		requestBody := models.RegisterRequest{
			Name:     "",
			Email:    "invalid-email",
			Password: "",
		}

		expectedStatus := http.StatusBadRequest
		expectedMessage := "Invalid request body"

		// Initialize mock service (not necessary to mock CreateUser here)
		mockUserService := new(userservice.MockUserService)

		// Setup the mock service and controller
		w, c := setupTest(t, requestBody)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.CreateUserController(c)

		// Assertions
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)

		// Verify that the mock service method was not called (as we expect early return due to validation failure)
		mockUserService.AssertExpectations(t)
	})

	// Subtest for Service Creation Error
	t.Run("Service Creation Error", func(t *testing.T) {
		requestBody := models.RegisterRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}

		expectedStatus := http.StatusInternalServerError
		expectedMessage := "Failed to create user"
		mockServiceError := errors.New("database error")

		// Initialize mock service
		mockUserService := new(userservice.MockUserService)
		mockUserService.On("CreateUser", requestBody).Return(mockServiceError).Once()

		// Setup the mock service and controller
		w, c := setupTest(t, requestBody)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.CreateUserController(c)

		// Assertions
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)

		// Verify that the mock service method was called
		mockUserService.AssertExpectations(t)
	})
}

func TestCheckEmailUserController(t *testing.T) {
	// Subtest for Valid Email
	t.Run("Valid Email", func(t *testing.T) {
		requestBody := models.CheckEmailRequest{
			Email: "test@example.com",
		}

		existedUser := &models.User{
			Email: "test@example.com",
			Name:  "Test User",
			Role:  "admin",
		}
		expectedStatus := http.StatusOK
		expectedMessage := "User email exists"
		mockUserService := new(userservice.MockUserService)
		mockUserService.On("CheckUserEmail", requestBody.Email).Return(existedUser, nil).Once()

		// Setup the mock service and controller
		w, c := setupTest(t, requestBody)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.CheckEmailUserController(c)

		// Assertions
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)
	})

	// Subtest for Invalid Email
	t.Run("Invalid Email", func(t *testing.T) {
		requestBody := models.CheckEmailRequest{
			Email: "invalid_email",
		}
		expectedStatus := http.StatusBadRequest
		expectedMessage := "Invalid request body"

		// Initialize mock service (not necessary to mock CheckUserEmail here)
		mockUserService := new(userservice.MockUserService)
		// Setup the mock service and controller
		w, c := setupTest(t, requestBody)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.CheckEmailUserController(c)

		// Assertions
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)

		// Verify that the mock service method was not called (as we expect early return due to validation failure)
		mockUserService.AssertExpectations(t)
	})

	// Subtest for failure to check user email
	t.Run("Failure to check user email", func(t *testing.T) {
		requestBody := models.CheckEmailRequest{
			Email: "test@example.com",
		}

		expectedStatus := http.StatusInternalServerError
		expectedMessage := "Failed to check user email"
		mockUserService := new(userservice.MockUserService)
		mockUserService.On("CheckUserEmail", requestBody.Email).Return(nil, errors.New("User not found")).Once()

		// Setup the mock service and controller
		w, c := setupTest(t, requestBody)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.CheckEmailUserController(c)

		// Assertions
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)

		// Verify that the mock service method was called
		mockUserService.AssertExpectations(t)
	})
}

func TestResetUserPasswordController(t *testing.T) {
	// sub test for success reset password
	t.Run("Success reset password", func(t *testing.T) {
		resetRequest := models.LoginRequest{
			Email:    "test@example.com",
			Password: "newpassword",
		}

		expectedStatus := http.StatusOK
		expectedMessage := "User password reset successfully"
		mockUserService := new(userservice.MockUserService)
		mockUserService.On("ResetUserPassword", resetRequest).Return(nil).Once()

		w, c := setupTest(t, resetRequest)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.ResetUserPasswordController(c)

		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)

		// Verify that the mock service method was called
		mockUserService.AssertExpectations(t)
	})

	// sub test invalid request
	t.Run("Invalid request body", func(t *testing.T) {
		resetRequest := models.LoginRequest{
			Email:    "",
			Password: "newpassword",
		}
		expectedStatus := http.StatusBadRequest
		expectedMessage := "Invalid request body"

		// Initialize mock service (not necessary to mock CheckUserEmail here)
		mockUserService := new(userservice.MockUserService)
		// Setup the mock service and controller
		w, c := setupTest(t, resetRequest)
		userController := setupUserController(mockUserService)

		// Run the controller's method
		userController.ResetUserPasswordController(c)

		// assert
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)

		// Verify that the mock service method was not called (as we expect early return due to validation failure)
		mockUserService.AssertExpectations(t)
	})

	// sub test for failure to reset password
	t.Run("Failure to reset password", func(t *testing.T) {
		resetRequest := models.LoginRequest{
			Email:    "test@example.com",
			Password: "newpassword",
		}
		expectedStatus := http.StatusInternalServerError
		expectedMessage := "Failed to reset user password"
		mockUserService := new(userservice.MockUserService)
		// Setup the mock service and controller
		mockUserService.On("ResetUserPassword", resetRequest).Return(errors.New("Failed to reset user password")).Once()
		w, c := setupTest(t, resetRequest)
		userController := setupUserController(mockUserService)
		// Run the controller's method
		userController.ResetUserPasswordController(c)
		// assert
		assert.Equal(t, expectedStatus, w.Code)
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, response.Message)
		// Verify that the mock service method was called
		mockUserService.AssertExpectations(t)
	})
}
