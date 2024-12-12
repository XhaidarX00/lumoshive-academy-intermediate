package promotionrepository

import (
	"dashboard-ecommerce-team2/models"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PromotionRepository interface {
	Create(promotionInput *models.Promotion) error
	Update(promotionInput *models.Promotion) error
	Delete(id int) error
	GetAll() ([]models.Promotion, error)
	GetByID(id int) (*models.Promotion, error)
}

type promotionRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements PromotionRepository.
func (p *promotionRepository) Create(promotionInput *models.Promotion) error {
	err := p.DB.Create(promotionInput).Error
	if err != nil {
		p.Log.Error("Error from repo creating promotions:", zap.Error(err))
		return err
	}
	return nil
}

// Delete implements PromotionRepository.
func (p *promotionRepository) Delete(id int) error {
	var promotion models.Promotion
	err := p.DB.First(&promotion, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.Log.Warn("Promotion ID not found:", zap.Int("ID", id))
			return fmt.Errorf("Promotion with ID %d not found", id)
		}

		p.Log.Error("Error finding promotion before delete:", zap.Error(err))
		return err
	}

	// Hapus promotion jika ditemukan
	err = p.DB.Delete(&models.Promotion{}, id).Error
	if err != nil {
		p.Log.Error("Error deleting promotion:", zap.Error(err))
		return err
	}

	return nil
}

// GetAll implements PromotionRepository.
func (p *promotionRepository) GetAll() ([]models.Promotion, error) {
	promotions := []models.Promotion{}
	if err := p.DB.Find(&promotions).Error; err != nil {
		p.Log.Error("Error get all promotion:", zap.Error(err))
		return []models.Promotion{}, err
	}

	return promotions, nil
}

// GetByID implements PromotionRepository.
func (p *promotionRepository) GetByID(id int) (*models.Promotion, error) {
	var promotion models.Promotion

	// Mencari promotion berdasarkan ID
	err := p.DB.First(&promotion, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.Log.Warn("Promotion ID not found", zap.Int("ID", id))
			return nil, fmt.Errorf("promotion with ID %d not found", id)
		}

		p.Log.Error("Error retrieving promotion by ID", zap.Int("ID", id), zap.Error(err))
		return nil, err
	}

	return &promotion, nil
}

// Update implements PromotionRepository.
func (p *promotionRepository) Update(promotionInput *models.Promotion) error {
	err := p.DB.Save(promotionInput).Error
	if err != nil {
		p.Log.Error("Error from repo updating promption:", zap.Error(err))
		return err
	}

	return nil
}

func NewPromotionRepository(db *gorm.DB, log *zap.Logger) PromotionRepository {
	return &promotionRepository{DB: db, Log: log}
}
