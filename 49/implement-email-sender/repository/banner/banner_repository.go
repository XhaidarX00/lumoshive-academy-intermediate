package bannerrepository

import (
	"dashboard-ecommerce-team2/models"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BannerRepository interface {
	Create(bannerInput *models.Banner) error
	Update(bannerInput *models.Banner) error
	Delete(id int) error
	GetByID(id int) (*models.Banner, error)
}

type BannerRepo struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements BannerRepository.
func (b *BannerRepo) Create(bannerInput *models.Banner) error {
	if bannerInput.Image == "" {
		return errors.New("invalid image url")
	}

	err := b.DB.Create(bannerInput).Error
	if err != nil {
		b.Log.Error("Error from repo creating banner:", zap.Error(err))
		return errors.New("failed create banner")
	}

	return nil
}

// Delete implements BannerRepository.
func (b *BannerRepo) Delete(id int) error {
	var banner models.Banner
	err := b.DB.First(&banner, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.Warn("Banner ID not found:", zap.Int("ID", id))
			return fmt.Errorf("banner with ID %d not found", id)
		}

		b.Log.Error("Error finding banner before delete:", zap.Error(err))
		return errors.New("Error deleting banner")
	}

	// Hapus banner jika ditemukan
	err = b.DB.Delete(&models.Banner{}, id).Error
	if err != nil {
		b.Log.Error("Error deleting banner:", zap.Error(err))
		return errors.New("Error deleting banner")
	}

	return nil
}

// GetByID implements BannerRepository.
func (b *BannerRepo) GetByID(id int) (*models.Banner, error) {
	var banner models.Banner

	// Mencari banner berdasarkan ID
	err := b.DB.First(&banner, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.Warn("Banner ID not found", zap.Int("ID", id))
			return nil, fmt.Errorf("banner with ID %d not found", id)
		}

		b.Log.Error("Error retrieving banner by ID", zap.Int("ID", id), zap.Error(err))
		return nil, err
	}

	return &banner, nil
}

// Update implements BannerRepository.
func (b *BannerRepo) Update(bannerInput *models.Banner) error {
	err := b.DB.First(bannerInput, bannerInput.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.Warn("Banner ID not found", zap.Int("ID", bannerInput.ID))
			return fmt.Errorf("banner with ID %d not found", bannerInput.ID)
		}

		b.Log.Error("Error retrieving banner by ID", zap.Int("ID", bannerInput.ID), zap.Error(err))
		return err
	}

	bannerInput.Published = !bannerInput.Published
	err = b.DB.Save(bannerInput).Error
	if err != nil {
		b.Log.Error("Error from repo updating banner:", zap.Error(err))
		return err
	}

	return nil
}

func NewBannerRepository(db *gorm.DB, log *zap.Logger) BannerRepository {
	return &BannerRepo{DB: db, Log: log}
}
