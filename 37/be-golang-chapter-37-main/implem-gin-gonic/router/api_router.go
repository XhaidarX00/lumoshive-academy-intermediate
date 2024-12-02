package router

import (
	"golang-chapter-37/implem-gin-gonic/controller"

	"github.com/gin-gonic/gin"
)

func APIRouter(router *gin.Engine, ctl controller.Controller) {
	router.GET("/customers", ctl.Customer.GetCustomers)
	router.POST("/customers", ctl.Customer.CreateCustomer)

	router.GET("/products", ctl.Product.GetProducts)
	router.POST("/products", ctl.Product.CreateProduct)

	// router.GET()
}
