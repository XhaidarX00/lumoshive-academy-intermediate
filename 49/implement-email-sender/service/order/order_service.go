package orderservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderService interface {
	UpdateOrderStatus(id int, status string) error
	GetAllOrders(page, limit int) ([]models.Order, int64, error)
	GetOrderByID(id int) (*models.Order, error)
	DeleteOrder(id int) error
	GetOrderDetail(id int) (*models.Order, []models.OrderItem, error)
}

type orderService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// GetAllOrders implements OrderService.
func (o *orderService) GetAllOrders(page, limit int) ([]models.Order, int64, error) {
	o.Log.Info("Fetching paginated orders", zap.Int("page", page), zap.Int("limit", limit))

	// Panggil repository untuk mendapatkan data dengan pagination
	orders, totalItems, err := o.Repo.Order.GetAll(page, limit)
	if err != nil {
		o.Log.Error("Failed to fetch paginated orders", zap.Error(err))
		return nil, 0, err
	}

	return orders, totalItems, nil
}

// GetOrderByID implements OrderService.
func (o *orderService) GetOrderByID(id int) (*models.Order, error) {
	o.Log.Info("Fetching order by ID", zap.Int("id", id))
	order, err := o.Repo.Order.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			o.Log.Warn("Order not found", zap.Int("id", id))
			return nil, errors.New("order not found")
		}
		o.Log.Error("Failed to fetch order", zap.Error(err))
		return nil, err
	}
	return order, nil
}

// UpdateOrderStatus implements OrderService.
func (o *orderService) UpdateOrderStatus(id int, status string) error {
	o.Log.Info("Updating order status", zap.Int("id", id), zap.String("status", status))
	order, err := o.Repo.Order.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			o.Log.Warn("Order not found", zap.Int("id", id))
			return errors.New("order not found")
		}
		o.Log.Error("Failed to fetch order", zap.Error(err))
		return err
	}

	order.Status = status
	if err := o.Repo.Order.UpdateStatus(id, status); err != nil {
		o.Log.Error("Failed to update order status", zap.Error(err))
		return err
	}

	o.Log.Info("Order status updated successfully", zap.Int("id", id), zap.String("status", status))
	return nil
}

func (o *orderService) DeleteOrder(id int) error {
	o.Log.Info("Deleting order", zap.Int("id", id))
	err := o.Repo.Order.DeleteOrder(id)
	if err != nil {
		o.Log.Error("Failed to delete order", zap.Error(err))
		return err
	}
	o.Log.Info("Order deleted successfully", zap.Int("id", id))
	return nil
}

func (o *orderService) GetOrderDetail(id int) (*models.Order, []models.OrderItem, error) {
	o.Log.Info("Fetching order details", zap.Int("id", id))
	order, orderItems, err := o.Repo.Order.GetDetail(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			o.Log.Warn("Order not found", zap.Int("id", id))
			return nil, nil, errors.New("order not found")
		}
		o.Log.Error("Failed to fetch order details", zap.Error(err))
		return nil, nil, err
	}
	return order, orderItems, nil
}

func NewOrderService(repo repository.Repository, log *zap.Logger) OrderService {
	return &orderService{Repo: repo, Log: log}
}
