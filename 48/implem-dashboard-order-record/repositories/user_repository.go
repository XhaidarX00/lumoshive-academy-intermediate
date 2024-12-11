// repositories/user_repository.go
package repositories

import (
	"project_auth_jwt/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
}

type UserRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		r.Log.Error("Failed to create user", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	err := r.DB.Save(user).Error
	if err != nil {
		r.Log.Error("Failed to create user", zap.Error(err))
		return err
	}
	return nil
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		r.Log.Error("Failed to create user", zap.Error(result.Error))
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		r.Log.Error("Failed to create user", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepositoryInterface {
	return &UserRepository{DB: db, Log: log}
}
