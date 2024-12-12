package stockrepository

import (
	"dashboard-ecommerce-team2/models"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StockRepository interface {
	Update(stockHistory *models.StockRequest) error
	Delete(id int) error
	GetByID(id int) (*models.StockResponse, error)
}

type stockRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Delete implements StockRepository.
func (s *stockRepository) Delete(id int) error {
	var stockHistory models.Stock

	err := s.DB.First(&stockHistory, id).Error
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("history not found")
		}
		s.Log.Error("Failed to fetch history for deletion", zap.Error(err))
		return err
	}

	err = s.DB.Delete(&stockHistory).Error
	if err != nil {
		s.Log.Error("Failed to delete history", zap.Error(err))
		return err
	}
	return nil
}

// GetByID implements StockRepository.
func (s *stockRepository) GetByID(id int) (*models.StockResponse, error) {
	var stockHistory models.StockResponse
	var variantJSON []byte

	err := s.DB.Table("stocks").
		Select("stocks.id, stocks.product_id, stocks.type, stocks.quantity, stocks.created_at, stocks.updated_at, products.name as product_name, CASE WHEN categories.variant = '{}' THEN NULL ELSE categories.variant END AS variant").
		Joins("JOIN products ON products.id = stocks.product_id").
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("stocks.id = ?", id).
		First(&stockHistory).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("product not found")
		}
		s.Log.Error("Failed to fetch product by ID", zap.Error(err))
		return nil, err
	}
	if len(variantJSON) > 0 {
		if err := json.Unmarshal(variantJSON, &stockHistory.Variant); err != nil {
			s.Log.Error("Repository: failed to unmarshal variant JSON", zap.Error(err))
			return nil, err
		}
	}

	return &stockHistory, nil
}

// UpdateStock implements StockRepository.
func (s *stockRepository) Update(stockHistory *models.StockRequest) error {
	tx := s.DB.Begin()

	var currentStock int
	if err := tx.Table("products").
		Select("stock").
		Where("id = ?", stockHistory.ProductID).
		Scan(&currentStock).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch current stock: %w", err)
	}

	newStock := currentStock
	if stockHistory.Type == "in" {
		newStock += stockHistory.Quantity
	} else if stockHistory.Type == "out" {
		if stockHistory.Quantity > currentStock {
			tx.Rollback()
			return fmt.Errorf("insufficient stock")
		}
		newStock -= stockHistory.Quantity
	} else {
		tx.Rollback()
		return fmt.Errorf("invalid stock type")
	}

	if err := tx.Table("products").
		Where("id = ?", stockHistory.ProductID).
		Updates(map[string]interface{}{
			"stock":      newStock,
			"updated_at": time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update product stock: %w", err)
	}

	if err := tx.Table("stocks").Create(stockHistory).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create stock history: %w", err)
	}

	tx.Commit()
	return nil
}

func NewStockRepository(db *gorm.DB, log *zap.Logger) StockRepository {
	return &stockRepository{DB: db, Log: log}
}
