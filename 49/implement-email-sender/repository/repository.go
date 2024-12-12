package repository

import (
	bannerrepository "dashboard-ecommerce-team2/repository/banner"
	categoryrepository "dashboard-ecommerce-team2/repository/category"
	orderrepository "dashboard-ecommerce-team2/repository/order"
	productrepository "dashboard-ecommerce-team2/repository/product"
	promotionrepository "dashboard-ecommerce-team2/repository/promotion"
	stockrepository "dashboard-ecommerce-team2/repository/stock"
	userrepository "dashboard-ecommerce-team2/repository/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Banner    bannerrepository.BannerRepository
	Category  categoryrepository.CategoryRepository
	Order     orderrepository.OrderRepository
	Product   productrepository.ProductRepository
	Promotion promotionrepository.PromotionRepository
	Stock     stockrepository.StockRepository
	User      userrepository.UserRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		Banner:    bannerrepository.NewBannerRepository(db, log),
		Category:  categoryrepository.NewCategoryRepository(db, log),
		Order:     orderrepository.NewOrderRepository(db, log),
		Product:   productrepository.NewProductRepository(db, log),
		Promotion: promotionrepository.NewPromotionRepository(db, log),
		Stock:     stockrepository.NewStockRepository(db, log),
		User:      userrepository.NewUserRepository(db, log),
	}
}
