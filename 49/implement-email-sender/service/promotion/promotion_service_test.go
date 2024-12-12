package promotionservice_test

import (
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	promotionrepository "dashboard-ecommerce-team2/repository/promotion"
	promotionservice "dashboard-ecommerce-team2/service/promotion"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetAllPromotionService(t *testing.T) {
	MockPromotion := new(promotionrepository.PromotionRepoMock)
	MockRepo := &repository.Repository{
		Promotion: MockPromotion,
	}

	logger := zap.NewNop()
	service := promotionservice.NewPromotionService(*MockRepo, logger)

	// GetAll
	t.Run("Succes Get All Promotion", func(t *testing.T) {
		promotions := []models.Promotion{
			{ID: 1, Name: "Promo A"},
			{ID: 2, Name: "Promo B"},
		}

		MockPromotion.On("GetAll").Once().Return(promotions, nil)

		result, err := service.GetAllPromotions()

		assert.NoError(t, err)
		assert.Greater(t, len(result), 1)
	})

	t.Run("Failed Get All Promotion", func(t *testing.T) {
		promotions := []models.Promotion{}

		// Konfigurasikan mock untuk mengembalikan error
		MockPromotion.On("GetAll").Return(promotions, errors.New("failed to get promotions"))

		result, err := service.GetAllPromotions()

		// Verifikasi bahwa error terjadi
		assert.Error(t, err, "Error should not be nil")
		assert.Equal(t, "failed to get promotions", err.Error())
		assert.Equal(t, 0, len(result))
	})

}
