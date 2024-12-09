package banner

import (
	"voucher_system/models"
	"voucher_system/repository/banner"

	"go.uber.org/zap"
)

type BannerServiceInterface interface {
	Create(banner *models.Banner) error
	Update(banner *models.Banner) error
}

type BannerService struct {
	Repo banner.BannerRepositoryInterface
	Log  *zap.Logger
}

func NewBannerService(repo banner.BannerRepositoryInterface, log *zap.Logger) BannerServiceInterface {
	return &BannerService{
		Repo: repo,
		Log:  log,
	}
}

func (b *BannerService) Create(banner *models.Banner) error {
	return b.Repo.Create(banner)
}

func (b *BannerService) Update(banner *models.Banner) error {
	return b.Repo.Update(banner)
}
