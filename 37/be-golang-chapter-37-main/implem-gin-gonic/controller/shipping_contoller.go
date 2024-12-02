package controller

import (
	"golang-chapter-37/implem-gin-gonic/model"
	"golang-chapter-37/implem-gin-gonic/repository"
	"golang-chapter-37/implem-gin-gonic/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShippingController struct {
	shippingService service.ShippingService
}

func NewShippingController(db *gorm.DB, cs repository.CustomerRepository, slr repository.SellerRepository) *ShippingController {
	return &ShippingController{shippingService: service.NewShippingService(db, cs, slr)}
}

func (sc *ShippingController) GetShippingServices(c *gin.Context) {
	var services []model.ShippingService

	err := sc.shippingService.GetShippingServices(&services)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"services": services})
}
