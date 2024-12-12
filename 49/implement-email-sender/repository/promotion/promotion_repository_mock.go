package promotionrepository

import (
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
)

type PromotionRepoMock struct {
	mock.Mock
}

func (p *PromotionRepoMock) Create(promotionInput *models.Promotion) error {
	return nil
}

func (p *PromotionRepoMock) Update(promotionInput *models.Promotion) error {
	return nil
}

func (p *PromotionRepoMock) Delete(id int) error {
	return nil
}

func (p *PromotionRepoMock) GetAll() ([]models.Promotion, error) {
	args := p.Called()

	if banner := args.Get(0); banner != nil {
		return banner.([]models.Promotion), args.Error(1)
	}

	return nil, args.Error(1)
}

func (p *PromotionRepoMock) GetByID(id int) (*models.Promotion, error) {
	return nil, nil
}
