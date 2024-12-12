package service

import (
	"dashboard-ecommerce-team2/repository"
	bannerservice "dashboard-ecommerce-team2/service/banner"
	categoryservice "dashboard-ecommerce-team2/service/category"
	dashboardservice "dashboard-ecommerce-team2/service/dashboard"
	orderservice "dashboard-ecommerce-team2/service/order"
	productservice "dashboard-ecommerce-team2/service/product"
	promotionservice "dashboard-ecommerce-team2/service/promotion"
	sendemail "dashboard-ecommerce-team2/service/send_email"
	stockservice "dashboard-ecommerce-team2/service/stock"
	userservice "dashboard-ecommerce-team2/service/user"

	"go.uber.org/zap"
)

type Service struct {
	Banner    bannerservice.BannerService
	Category  categoryservice.CategoryService
	Dashboard dashboardservice.DashboardService
	Order     orderservice.OrderService
	Product   productservice.ProductService
	Promotion promotionservice.PromotionService
	Stock     stockservice.StockService
	User      userservice.UserService
	SendOtp   sendemail.EmailSenderInterface
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Banner:    bannerservice.NewBannerService(repo, log),
		Category:  categoryservice.NewCategoryService(repo, log),
		Dashboard: dashboardservice.NewDashboardService(repo, log),
		Order:     orderservice.NewOrderService(repo, log),
		Product:   productservice.NewProductService(repo, log),
		Promotion: promotionservice.NewPromotionService(repo, log),
		Stock:     stockservice.NewStockService(repo, log),
		User:      userservice.NewUserService(repo, log),
		SendOtp:   sendemail.NewEmailSenderService(log),
	}
}
