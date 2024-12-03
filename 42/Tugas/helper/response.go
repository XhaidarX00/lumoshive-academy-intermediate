package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, message string, statusCode int) {
	response := Response{
		Status:  true,
		Message: message,
	}
	c.JSON(statusCode, response)
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	respon := Response{
		Status:  false,
		Message: message,
	}
	c.JSON(statusCode, respon)
}

func SuccessResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	respon := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, respon)
}
