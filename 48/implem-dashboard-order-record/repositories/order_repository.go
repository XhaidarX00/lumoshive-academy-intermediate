package repositories

import (
	"project_auth_jwt/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	GetAllOrders() ([]models.Order, error)
}

type OrderRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewOrderRepository(db *gorm.DB, log *zap.Logger) OrderRepositoryInterface {
	return &OrderRepository{
		DB:  db,
		Log: log,
	}
}

func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Find(&orders).Error
	if err != nil {
		r.Log.Error("Failed to create order", zap.Error(err))
		return nil, err
	}

	return orders, nil
}
