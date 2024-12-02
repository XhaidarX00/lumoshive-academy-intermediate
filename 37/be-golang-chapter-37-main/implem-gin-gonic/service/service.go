package service

import (
	"golang-chapter-37/implem-gin-gonic/repository"

	"gorm.io/gorm"
)

type AllService struct {
	customer CustomerService
	shipping ShippingService
	product  ProductService
}

func NewAllService(db *gorm.DB, cs repository.CustomerRepository, slr repository.SellerRepository) *AllService {
	return &AllService{
		customer: NewCustomerService(db),
		shipping: NewShippingService(db, cs, slr),
		product:  NewProductService(db),
	}
}
