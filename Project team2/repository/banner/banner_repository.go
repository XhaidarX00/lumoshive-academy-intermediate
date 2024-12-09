package banner

import (
	"errors"
	"fmt"
	"time"
	"voucher_system/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BannerRepositoryInterface interface {
	Create(banner *models.Banner) error
	Update(banner *models.Banner) error
	Delete(id uint) error
	GetBannerByID(id uint) (*models.Banner, error)
	ListBanner(page, limit int) ([]models.Banner, int64, error)
	GetPublishedBanners() ([]models.Banner, error)
}

type BannerRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewBannerRepository(db *gorm.DB, log *zap.Logger) BannerRepositoryInterface {
	return &BannerRepository{
		DB:  db,
		Log: log,
	}
}

func (b *BannerRepository) Create(banner *models.Banner) error {
	err := b.DB.Create(banner).Error
	if err != nil {
		b.Log.Error("Error from repo creating banner:", zap.Error(err))
		return err
	}

	return nil
}

func (b *BannerRepository) Update(banner *models.Banner) error {
	err := b.DB.Save(banner).Error
	if err != nil {
		b.Log.Error("Error from repo updating banner:", zap.Error(err))
		return err
	}

	return nil
}

func (b *BannerRepository) Delete(id uint) error {
	var banner models.Banner
	err := b.DB.First(&banner, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.Warn("Banner ID not found:", zap.Uint("ID", id))
			return fmt.Errorf("banner with ID %d not found", id)
		}

		b.Log.Error("Error finding banner before delete:", zap.Error(err))
		return err
	}

	// Hapus banner jika ditemukan
	err = b.DB.Delete(&models.Banner{}, id).Error
	if err != nil {
		b.Log.Error("Error deleting banner:", zap.Error(err))
		return err
	}

	return nil
}

func (b *BannerRepository) GetBannerByID(id uint) (*models.Banner, error) {
	var banner models.Banner

	// Mencari banner berdasarkan ID
	err := b.DB.First(&banner, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.Warn("Banner ID not found", zap.Uint("ID", id))
			return nil, fmt.Errorf("banner with ID %d not found", id)
		}

		b.Log.Error("Error retrieving banner by ID", zap.Uint("ID", id), zap.Error(err))
		return nil, err
	}

	return &banner, nil
}

func (b *BannerRepository) ListBanner(page, limit int) ([]models.Banner, int64, error) {
	var banners []models.Banner
	var count int64

	// Count total records
	b.DB.Model(&models.Banner{}).Count(&count)

	// Paginate results
	err := b.DB.
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&banners).Error

	if err != nil {
		b.Log.Error("Error from repo ListBanner:", zap.Error(err))
		return nil, 0, err
	}

	return banners, count, nil
}

func (b *BannerRepository) GetPublishedBanners() ([]models.Banner, error) {
	var banners []models.Banner
	err := b.DB.
		Where("published = ? AND release_date <= ? AND end_date >= ?",
			true, time.Now(), time.Now()).
		Find(&banners).Error

	if err != nil {
		b.Log.Error("Error from repo :", zap.Error(err))
		return nil, err
	}

	return banners, err
}
