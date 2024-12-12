package orderrepository

import (
	"dashboard-ecommerce-team2/models"
	"math"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository interface {
	UpdateStatus(id int, status string) error
	GetByID(id int) (*models.Order, error)
	GetAll(page, limit int) ([]models.Order, int64, error)
	CountOrder() (int, error)
	CountTotalPriceOrder() (float64, error)
	GetEarningEachMonth() ([]models.Revenue, error)
	GetDetail(id int) (*models.Order, []models.OrderItem, error)
	DeleteOrder(id int) error
}

type orderRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// GetEarningEachMonth implements OrderRepository.
func (o *orderRepository) GetEarningEachMonth() ([]models.Revenue, error) {
	var revenues []models.Revenue
	currentYear := time.Now().Year()

	// Perform the query to calculate total earnings for each month of the current year
	err := o.DB.Table("(VALUES (1), (2), (3), (4), (5), (6), (7), (8), (9), (10), (11), (12)) AS months(month)").
		Select("TO_CHAR(DATE_TRUNC('month', TO_DATE(CAST(month AS TEXT), 'MM')), 'Month') AS month, COALESCE(SUM(total_amount), 0) AS total_earning").
		Joins(`LEFT JOIN orders ON 
				EXTRACT(MONTH FROM orders.created_at) = months.month 
				AND EXTRACT(YEAR FROM orders.created_at) = ? 
				AND orders.status = ?`, currentYear, "completed").
		Group("months.month").
		Order("months.month ASC").
		Scan(&revenues).Error

	if err != nil {
		return nil, err
	}

	for _, revenue := range revenues {
		totalEarning := revenue.TotalEarning * ((100 - 10) / 100)
		revenue.TotalEarning = math.Round(totalEarning*100) / 100
	}

	return revenues, nil
}

// CountTotalPriceOrder implements OrderRepository.
func (o *orderRepository) CountTotalPriceOrder() (float64, error) {
	var totalPrice float64
	err := o.DB.Table("orders").
		Select("COALESCE(SUM(total_amount), 0)").
		Where("status = ? AND EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", "completed", time.Now().Month(), time.Now().Year()).
		Scan(&totalPrice).Error

	if err != nil {
		return 0, err
	}
	newTotal := math.Round(totalPrice*100) / 100
	return newTotal, nil
}

// CountOrder implements OrderRepository.
func (o *orderRepository) CountOrder() (int, error) {
	var count int64
	err := o.DB.Model(&models.Order{}).
		Where("status = ? AND EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", "completed", time.Now().Month(), time.Now().Year()).
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// GetAll implements OrderRepository.
func (o *orderRepository) GetAll(page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var totalItems int64

	offset := (page - 1) * limit

	if err := o.DB.Model(&models.Order{}).Count(&totalItems).Error; err != nil {
		o.Log.Error("Failed to count orders", zap.Error(err))
		return nil, 0, err
	}

	if err := o.DB.Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		o.Log.Error("Failed to retrieve paginated orders", zap.Error(err))
		return nil, 0, err
	}

	return orders, totalItems, nil
}

// 2 param 3 return
// kondisi :
// 1. if error nil, 0, err
// 2. if error nil, 0, err
// if success 3. []models.Order, int64, nil

// GetByID implements OrderRepository.
func (o *orderRepository) GetByID(id int) (*models.Order, error) {
	order := models.Order{}
	if err := o.DB.First(&order, id).Error; err != nil {
		o.Log.Error("Failed to find order", zap.Error(err))
		return nil, err
	}
	return &order, nil
}

// UpdateStatus implements OrderRepository.
func (o *orderRepository) UpdateStatus(id int, status string) error {
	order := models.Order{}
	if err := o.DB.First(&order, id).Error; err != nil {
		o.Log.Error("Failed to find order", zap.Error(err))
		return err
	}
	order.Status = status
	return o.DB.Save(&order).Error
}

func (o *orderRepository) DeleteOrder(id int) error {
	var orderItems []models.OrderItem
	if err := o.DB.Where("order_id = ?", id).Delete(&orderItems).Error; err != nil {
		o.Log.Error("Failed to delete order items", zap.Error(err))
		return err
	}

	if err := o.DB.Delete(&models.Order{}, id).Error; err != nil {
		o.Log.Error("Failed to delete order", zap.Error(err))
		return err
	}
	return nil
}

func (o *orderRepository) GetDetail(id int) (*models.Order, []models.OrderItem, error) {
	var order models.Order
	if err := o.DB.First(&order, id).Error; err != nil {
		o.Log.Error("Failed to find order", zap.Error(err))
		return nil, nil, err
	}

	var orderItems []models.OrderItem
	if err := o.DB.Where("order_id = ?", id).Find(&orderItems).Error; err != nil {
		o.Log.Error("Failed to retrieve order items", zap.Error(err))
		return nil, nil, err
	}

	return &order, orderItems, nil
}

func NewOrderRepository(db *gorm.DB, log *zap.Logger) OrderRepository {
	return &orderRepository{DB: db, Log: log}
}
