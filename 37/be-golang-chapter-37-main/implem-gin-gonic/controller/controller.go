package controller

import (
	"golang-chapter-37/implem-gin-gonic/repository"

	"gorm.io/gorm"
)

type Controller struct {
	Product  ProductController
	Customer CustomerController
	Shipping ShippingController
}

func NewController(db *gorm.DB, cs repository.CustomerRepository, slr repository.SellerRepository) Controller {
	return Controller{
		Product:  *NewProductController(db),
		Customer: *NewCustomerController(db),
		Shipping: *NewShippingController(db, cs, slr),
	}
}
