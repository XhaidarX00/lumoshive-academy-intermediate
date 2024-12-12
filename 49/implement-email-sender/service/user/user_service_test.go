package userservice

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	userrepository "dashboard-ecommerce-team2/repository/user"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func setupUserService(mockRepository userrepository.UserRepository) *userService {
	mockLogger, _ := zap.NewProduction()
	return &userService{
		Repo: repository.Repository{User: mockRepository},
		Log:  mockLogger,
	}
}

func TestCheckUserEmail(t *testing.T) {
	mockRepo := new(userrepository.MockUserRepository)
	userService := setupUserService(mockRepo)

	t.Run("Valid Email - User Found", func(t *testing.T) {
		email := "test@example.com"
		expectedUser := &models.User{
			Email: "test@example.com",
			Name:  "Test User",
			Role:  "admin",
		}

		// Setup mock expectations

		// Create instance of the user service with the mock repository
		mockRepo.On("GetByEmail", email).Return(expectedUser, nil).Once()

		// Call the service method
		user, err := userService.CheckUserEmail(email)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Valid Email - User Not Found", func(t *testing.T) {
		email := "nonexistent@example.com"
		expectedError := errors.New("user not found")

		// Setup mock expectations
		mockRepo.On("GetByEmail", email).Return((*models.User)(nil), expectedError).Once()

		// Call the service method
		user, err := userService.CheckUserEmail(email)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Email - Error from Repository", func(t *testing.T) {
		email := "invalid-email"
		expectedError := errors.New("invalid email format")

		// Setup mock expectations
		mockRepo.On("GetByEmail", email).Return((*models.User)(nil), expectedError).Once()

		// Call the service method
		user, err := userService.CheckUserEmail(email)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(userrepository.MockUserRepository)
	userService := setupUserService(mockRepo)

	t.Run("Successful user creation", func(t *testing.T) {
		userInput := models.RegisterRequest{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "John Doe",
		}

		// Create a new user with a hashed password
		hashedPassword := helper.HashPassword(userInput.Password) // hash the password
		expectedUser := models.User{
			Email:    userInput.Email,
			Password: hashedPassword, // hashed password
			Role:     "staff",
			Name:     userInput.Name,
		}

		// Setup mock expectations with a custom matcher for Password (ignoring hash)
		mockRepo.On("Create", mock.MatchedBy(func(user models.User) bool {
			return user.Email == expectedUser.Email &&
				user.Role == expectedUser.Role &&
				user.Name == expectedUser.Name
		})).Return(nil).Once()

		// Call the service method
		err := userService.CreateUser(userInput)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error while creating user", func(t *testing.T) {
		userInput := models.RegisterRequest{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "John Doe",
		}
		expectedError := errors.New("unable to create user")

		// Setup mock expectations
		mockRepo.On("Create", mock.Anything).Return(expectedError).Once()

		// Call the service method
		err := userService.CreateUser(userInput)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestResetUserPassword(t *testing.T) {
	mockRepo := new(userrepository.MockUserRepository)
	userService := setupUserService(mockRepo)

	t.Run("Successful password reset", func(t *testing.T) {
		// Sample input for the password reset
		input := models.LoginRequest{
			Email:    "test@example.com",
			Password: "newPassword123",
		}

		// Hash the password as done in the service method
		hashedPassword := helper.HashPassword(input.Password)

		// Prepare the expected reset password input with hashed password
		resetPassword := models.LoginRequest{
			Email:    input.Email,
			Password: hashedPassword, // hashed password
		}

		// Setup mock expectations
		mockRepo.On("UpdatePassword", mock.MatchedBy(func(loginReq models.LoginRequest) bool {
			// Check if email matches and password is hashed correctly (ignore actual hash comparison)
			return loginReq.Email == resetPassword.Email
		})).Return(nil).Once()

		// Call the service method
		err := userService.ResetUserPassword(input)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error updating password - Repository error", func(t *testing.T) {
		// Sample input for the password reset
		input := models.LoginRequest{
			Email:    "test@example.com",
			Password: "newPassword123",
		}

		// Hash the password as done in the service method
		hashedPassword := helper.HashPassword(input.Password)

		// Prepare the expected reset password input with hashed password
		resetPassword := models.LoginRequest{
			Email:    input.Email,
			Password: hashedPassword, // hashed password
		}

		// Setup mock expectations for a failure from the repository
		mockRepo.On("UpdatePassword", mock.MatchedBy(func(loginReq models.LoginRequest) bool {
			return loginReq.Email == resetPassword.Email
		})).Return(errors.New("failed to update password")).Once()

		// Call the service method
		err := userService.ResetUserPassword(input)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "failed to update password", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockRepo := new(userrepository.MockUserRepository)
	userService := setupUserService(mockRepo)

	t.Run("Successful login - Valid credentials", func(t *testing.T) {
		// Sample valid input for login
		input := models.LoginRequest{
			Email:    "test@example.com",
			Password: "correctPassword123",
		}

		// Expected user returned by the mock repository
		expectedUser := &models.User{
			Email:    input.Email,
			Password: helper.HashPassword("correctPassword123"), // Hash the correct password
			Role:     "staff",
			Name:     "Test User",
		}

		// Mock GetByEmail to return the expected user without error
		mockRepo.On("GetByEmail", input.Email).Return(expectedUser, nil).Once()

		// Call the service method
		user, err := userService.Login(input)

		// Assertions: Check only the email
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.Email, user.Email) // Check email only
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		// Sample input for login
		input := models.LoginRequest{
			Email:    "nonexistent@example.com",
			Password: "password123",
		}

		// Mock GetByEmail to return an error (user not found)
		mockRepo.On("GetByEmail", input.Email).Return(nil, errors.New("user not found")).Once()

		// Call the service method
		user, err := userService.Login(input)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user not found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error - Invalid password", func(t *testing.T) {
		// Sample valid input for login
		input := models.LoginRequest{
			Email:    "test@example.com",
			Password: "wrongPassword123", // Incorrect password
		}

		// Expected user returned by the mock repository
		expectedUser := &models.User{
			Email:    input.Email,
			Password: helper.HashPassword("correctPassword123"), // Hash the correct password
			Role:     "staff",
			Name:     "Test User",
		}

		// Mock GetByEmail to return the expected user without error
		mockRepo.On("GetByEmail", input.Email).Return(expectedUser, nil).Once()

		// Call the service method
		user, err := userService.Login(input)

		// Assertions: Check only the email
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid user password", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error - Repository error", func(t *testing.T) {
		// Sample valid input for login
		input := models.LoginRequest{
			Email:    "test@example.com",
			Password: "correctPassword123",
		}

		// Mock GetByEmail to return an error (e.g., database error)
		mockRepo.On("GetByEmail", input.Email).Return(nil, errors.New("database error")).Once()

		// Call the service method
		user, err := userService.Login(input)

		// Assertions: Check that email is not returned, and an error occurred
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "database error", err.Error())
		mockRepo.AssertExpectations(t)
	})

}
