package repository

import (
	"voucher_system/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(email string) (models.User, error)
}

type userRepository struct {
	DB  *gorm.DB
	log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{DB: db, log: log}
}

func (r *userRepository) Login(email string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}
