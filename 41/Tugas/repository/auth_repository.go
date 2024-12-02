package repository

import (
	"fmt"
	"project-voucher-team3/database"
	"project-voucher-team3/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	RDB database.Cacher
	DB  *gorm.DB
}

func NewAuthRepository(rdb database.Cacher, db *gorm.DB) *AuthRepository {
	return &AuthRepository{rdb, db}
}

func (a AuthRepository) Login(user *models.User, auth *models.Auth) error {
	err := a.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return err
	}

	key := fmt.Sprintf("TOKEN_%s", user.Email)
	token := uuid.NewString()
	err = a.RDB.Set(key, token)
	if err != nil {
		return err
	}

	auth.ID = key
	auth.Token = token
	return nil
}
