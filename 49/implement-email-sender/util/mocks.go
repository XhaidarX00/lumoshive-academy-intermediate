package utils

import (
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
)

type Service struct {
	Banner *BannerService
}

type BannerService struct {
	mock.Mock
}

func (m *BannerService) CreateBanner(banner *models.Banner) error {
	args := m.Called(banner)
	return args.Error(0)
}

func (m *BannerService) GetBannerByID(id int) (*models.Banner, error) {
	args := m.Called(id)
	if banner, ok := args.Get(0).(*models.Banner); ok {
		return banner, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *BannerService) UpdateBanner(banner models.Banner) error {
	args := m.Called(banner)
	return args.Error(0)
}

func (m *BannerService) DeleteBanner(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
