package bannerservice

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"

	"go.uber.org/zap"
)

type BannerService interface {
	CreateBanner(bannerInput *models.Banner) error
	UpdateBanner(bannerInput *models.Banner) error
	GetBannerByID(id int) (*models.Banner, error)
	DeleteBanner(id int) error
}

type bannerService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// jika tidak ada logic bisa diskip

// CreateBanner implements BannerService.
func (b *bannerService) CreateBanner(bannerInput *models.Banner) error {
	return b.Repo.Banner.Create(bannerInput)
}

// DeleteBanner implements BannerService.
func (b *bannerService) DeleteBanner(id int) error {
	return b.Repo.Banner.Delete(id)
}

// GetBannerByID implements BannerService.
func (b *bannerService) GetBannerByID(id int) (*models.Banner, error) {
	return b.Repo.Banner.GetByID(id)
}

// UpdateBanner implements BannerService.
func (b *bannerService) UpdateBanner(bannerInput *models.Banner) error {
	return b.Repo.Banner.Update(bannerInput)
}

func NewBannerService(repo repository.Repository, log *zap.Logger) BannerService {
	return &bannerService{Repo: repo, Log: log}
}
