package controller

import (
	"golang-chapter-37/implem-gin-gonic/model"
	"golang-chapter-37/implem-gin-gonic/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerController struct {
	service service.CustomerService
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		service: service.NewCustomerService(db),
	}
}

func (c *CustomerController) GetCustomers(ctx *gin.Context) {
	customers, err := c.service.GetAllCustomers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer model.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateCustomer(&customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, customer)
}
