package stockservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type StockService interface {
	UpdateProductStock(stockHistory *models.StockRequest) error
	GetProductStockDetail(id int) (*models.StockResponse, error)
	DeleteProductStock(id int) error
}

type stockService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// DeleteProductStock implements StockService.
func (s *stockService) DeleteProductStock(id int) error {
	err := s.Repo.Stock.Delete(id)
	if err != nil {
		return errors.New("failed to delete product")
	}
	return nil
}

// GetProductStockDetail implements StockService.
func (s *stockService) GetProductStockDetail(id int) (*models.StockResponse, error) {
	stockHistory, err := s.Repo.Stock.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %w", err)
	}
	return stockHistory, nil
}

// UpdateProductStock implements StockService.
func (s *stockService) UpdateProductStock(stockHistory *models.StockRequest) error {
	if err := s.Repo.Stock.Update(stockHistory); err != nil {
		return fmt.Errorf("failed to update stock: %w", err)
	}
	return nil
}

func NewStockService(repo repository.Repository, log *zap.Logger) StockService {
	return &stockService{Repo: repo, Log: log}
}
