package repository

import (
	"errors"
	"project-voucher-team3/models"
	"time"

	"gorm.io/gorm"
)

type ReedemRepository struct {
	DB *gorm.DB
}

func NewReedemRepository(db *gorm.DB) *ReedemRepository {
	return &ReedemRepository{db}
}

func (repo *ReedemRepository) GetUserRedeem(userID int, voucherFilter models.Voucher) ([]models.Redeem, error) {
	var redeems []models.Redeem

	err := repo.DB.Preload("Voucher", "voucher_type = ?", voucherFilter.VoucherType).
		Where("user_id = ?", userID).
		Find(&redeems).Error
	if err != nil {
		return nil, err
	}

	now := time.Now()
	activeRedeems := []models.Redeem{}
	for _, redeem := range redeems {
		if redeem.Voucher.ID != 0 && redeem.Voucher.StartDate.Before(now) && redeem.Voucher.EndDate.After(now) {
			activeRedeems = append(activeRedeems, redeem)
		}
	}

	return activeRedeems, nil
}

func (repo *ReedemRepository) RedeemVoucher(user *models.User, voucherID int) (models.Redeem, error) {
	var voucher models.Voucher
	if err := repo.DB.First(&voucher, voucherID).Error; err != nil {
		return models.Redeem{}, errors.New("voucher not found")
	}

	activeRedeems, err := repo.GetUserRedeem(user.ID, voucher)
	if err != nil {
		return models.Redeem{}, err
	}

	for _, voucher := range activeRedeems {
		if voucher.VoucherID == voucherID {
			return models.Redeem{}, errors.New("only one code in voucher for customers used")
		}
	}

	if voucher.MinRatePoint > user.Points {
		return models.Redeem{}, errors.New("not enough points to redeem")
	}

	redeem := models.Redeem{
		UserID:      user.ID,
		VoucherID:   voucherID,
		VoucherCode: voucher.VoucherCode,
		RedeemDate:  time.Now(),
		Voucher:     voucher,
	}

	if err := repo.DB.Create(&redeem).Error; err != nil {
		return models.Redeem{}, err
	}

	user.Points = user.Points - voucher.MinRatePoint
	return redeem, nil
}
