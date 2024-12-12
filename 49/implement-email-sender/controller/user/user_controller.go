package usercontroller

import (
	"dashboard-ecommerce-team2/config"
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	utils "dashboard-ecommerce-team2/util"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	Service service.Service
	Log     *zap.Logger
	Cacher  database.Cacher
	Config  config.Configuration
}

func NewUserController(service service.Service, log *zap.Logger, cacher database.Cacher, config config.Configuration) *UserController {
	return &UserController{
		Service: service,
		Log:     log,
		Cacher:  cacher,
		Config:  config,
	}
}

// CreateUserController godoc
// @Summary      Create a new user
// @Description  Register a new user with a provided request body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        registerRequest  body     models.RegisterRequest  true  "User Registration Request Body"
// @Success      201             {object} helper.HTTPResponse   "User created successfully"
// @Failure      400             {object} helper.HTTPResponse   "Invalid request body"
// @Failure      500             {object} helper.HTTPResponse   "Failed to create user"
// @Router       /auth/register [post]
func (ctrl *UserController) CreateUserController(c *gin.Context) {
	var registerReq models.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	err := ctrl.Service.User.CreateUser(registerReq)
	if err != nil {
		ctrl.Log.Error("Failed to create user", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to create user", http.StatusInternalServerError)
		return
	}

	if err := ctrl.Service.SendOtp.SendEmailRegis(registerReq.Email, registerReq.Name); err != nil {
		ctrl.Log.Error("Error OtpSendercontroller", zap.Error(fmt.Errorf("Failed to send email: %v", err)))
		helper.ResponseError(c, "Error OtpSendercontroller", fmt.Sprintf("Failed to send email: %v", err), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, nil, "User created successfully", http.StatusCreated)
}

// LoginController godoc
// @Summary      User Login
// @Description  Authenticate a user and return a token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        loginRequest  body     models.LoginRequest  true  "Login Request Body"
// @Success      200           {object} helper.HTTPResponse   "User logged in successfully"
// @Failure      400           {object} helper.HTTPResponse   "Invalid request body"
// @Failure      401           {object} helper.HTTPResponse   "Failed to login user"
// @Router       /auth/login [post]
func (ctrl *UserController) LoginController(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}
	user, err := ctrl.Service.User.Login(loginReq)
	if err != nil {
		ctrl.Log.Error("Failed to login user", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to login user", http.StatusUnauthorized)
		return
	}

	userIDStr := helper.IntToString(user.ID)
	key := fmt.Sprintf("UserID-%s", userIDStr)
	token := helper.GenerateToken(userIDStr, ctrl.Config.SecretKey)
	loginResponse := utils.LoginResponse{
		ID:    key,
		Role:  user.Role,
		Token: token,
	}

	ctrl.Cacher.SaveToken(key, token)
	helper.ResponseOK(c, loginResponse, "User logged in successfully", http.StatusOK)
}

// CheckEmailUserController godoc
// @Summary      Check if email is already registered
// @Description  Verify if a user with the given email already exists in the system
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        email  body     models.CheckEmailRequest  true  "Email to check"
// @Success      200    {object} helper.HTTPResponse     "Email check result"
// @Failure      400    {object} helper.HTTPResponse     "Invalid request body"
// @Failure      500    {object} helper.HTTPResponse     "Failed to check user email"
// @Router       /auth/check-email [post]
func (ctrl *UserController) CheckEmailUserController(c *gin.Context) {
	request := models.CheckEmailRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	existedUser, err := ctrl.Service.User.CheckUserEmail(request.Email)
	if err != nil {
		ctrl.Log.Error("Failed to check user email", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to check user email", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, existedUser, "User email exists", http.StatusOK)
}

// ResetUserPasswordController godoc
// @Summary      Reset user password
// @Description  Reset the password for a user using a provided request body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        resetRequest  body     models.LoginRequest  true  "User password reset request body"
// @Success      200           {object} helper.HTTPResponse   "User password reset successfully"
// @Failure      400           {object} helper.HTTPResponse   "Invalid request body"
// @Failure      500           {object} helper.HTTPResponse   "Failed to reset user password"
// @Router       /auth/reset-password [PATCH]
func (ctrl *UserController) ResetUserPasswordController(c *gin.Context) {
	var resetReq models.LoginRequest
	if err := c.ShouldBindJSON(&resetReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	err := ctrl.Service.User.ResetUserPassword(resetReq)
	if err != nil {
		ctrl.Log.Error("Failed to reset user password", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to reset user password", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, nil, "User password reset successfully", http.StatusOK)
}

func (ctrl *UserController) ResetPassOtpSenderController(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		ctrl.Log.Error("Invalid request body", zap.Error(fmt.Errorf("Error OtpSendercontroller")))
		helper.ResponseError(c, "Error OtpSendercontroller", "Invalid request body", http.StatusBadRequest)
		c.Abort()
		return
	}

	existedUser, err := ctrl.Service.User.CheckUserEmail(email)
	if err != nil {
		ctrl.Log.Error("Failed to check user email", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to check user email", http.StatusInternalServerError)
		return
	}

	otp := ctrl.Service.SendOtp.GenerateOTP()

	if err := ctrl.Service.SendOtp.SendEmail(email, otp); err != nil {
		ctrl.Log.Error("Error OtpSendercontroller", zap.Error(fmt.Errorf("Failed to send email: %v", err)))
		helper.ResponseError(c, "Error OtpSendercontroller", fmt.Sprintf("Failed to send email: %v", err), http.StatusInternalServerError)
		c.Abort()
		return
	}

	keyOtp := fmt.Sprintf("email-otp-%s", otp)
	data, _ := json.Marshal(map[string]string{
		"email": existedUser.Email,
		"otp":   otp,
	})

	err = ctrl.Cacher.SetExpire(keyOtp, string(data), 1)
	if err != nil {
		ctrl.Log.Error("Error OtpSendercontroller", zap.Error(fmt.Errorf("Failed to create otp key: %v", err)))
		helper.ResponseError(c, "Error OtpSendercontroller", fmt.Sprintf("Failed to create otp key: %v", err), http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, keyOtp, "Otp sender successfully", http.StatusOK)
	fmt.Printf("OTP sent to %s\n", email)
}

func (ctrl *UserController) ValidateOTPHandler(c *gin.Context) {
	request := models.OTPDataKey{}
	if err := c.ShouldBindJSON(&request); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(fmt.Errorf("Error ValidateOTPHandler")))
		helper.ResponseError(c, "Error ValidateOTPHandler", "Invalid request body", http.StatusBadRequest)
		c.Abort()
		return
	}

	dataOtp, err := ctrl.Cacher.Get(request.Key)
	if err != nil {
		ctrl.Log.Error("Error ValidateOTPHandler", zap.Error(fmt.Errorf("Otp Expired duration")))
		helper.ResponseError(c, "Error ValidateOTPHandler", "Otp Expired duration", http.StatusInternalServerError)
		c.Abort()
		return
	}

	otpData := &models.OTPData{}
	err = json.Unmarshal([]byte(dataOtp), otpData)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	if otpData.OTP != request.OTP {
		ctrl.Log.Error("Error ValidateOTPHandler", zap.Error(fmt.Errorf("Otp Invalid")))
		helper.ResponseError(c, "Error ValidateOTPHandler", "Otp invalid", http.StatusInternalServerError)
		c.Abort()
		return
	}

	existedUser, err := ctrl.Service.User.CheckUserEmail(otpData.Email)
	if err != nil {
		ctrl.Log.Error("Failed to check user email", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to check user email", http.StatusInternalServerError)
		return
	}

	var resetReq models.LoginRequest
	resetReq.Email = otpData.Email
	resetReq.Password = existedUser.Password

	helper.ResponseOK(c, resetReq, "valid otp successfully", http.StatusOK)
}
