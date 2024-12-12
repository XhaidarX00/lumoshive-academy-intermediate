package controller

import (
	"dashboard-ecommerce-team2/config"
	bannercontroller "dashboard-ecommerce-team2/controller/banner"
	categorycontroller "dashboard-ecommerce-team2/controller/category"
	dashboardcontroller "dashboard-ecommerce-team2/controller/dashboard"
	ordercontroller "dashboard-ecommerce-team2/controller/order"
	productcontroller "dashboard-ecommerce-team2/controller/product"
	promotioncontroller "dashboard-ecommerce-team2/controller/promotion"
	stockcontroller "dashboard-ecommerce-team2/controller/stock"
	usercontroller "dashboard-ecommerce-team2/controller/user"
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/service"

	"go.uber.org/zap"
)

type Controller struct {
	Banner    bannercontroller.BannerController
	Category  categorycontroller.CategoryController
	Dashboard dashboardcontroller.DashboardController
	Order     ordercontroller.OrderController
	Product   productcontroller.ProductController
	Promotion promotioncontroller.PromotionController
	Stock     stockcontroller.StockController
	User      usercontroller.UserController
}

func NewController(service service.Service, logger *zap.Logger, cacher database.Cacher, config config.Configuration) *Controller {
	return &Controller{
		Banner:    *bannercontroller.NewBannerController(service, logger),
		Category:  *categorycontroller.NewCategoryController(service, logger),
		Dashboard: *dashboardcontroller.NewDashboardController(service, logger),
		Order:     *ordercontroller.NewOrderController(service, logger),
		Product:   *productcontroller.NewProductController(service, logger),
		Promotion: *promotioncontroller.NewPromotionController(service, logger),
		Stock:     *stockcontroller.NewStockController(service, logger),
		User:      *usercontroller.NewUserController(service, logger, cacher, config),
	}
}
