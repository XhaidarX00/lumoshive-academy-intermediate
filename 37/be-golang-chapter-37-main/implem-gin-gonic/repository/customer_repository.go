package repository

import (
	"fmt"
	"golang-chapter-37/implem-gin-gonic/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll() ([]model.Customer, error)
	Save(customer *model.Customer) error
	GetCustomerAddressesByUserID(userID int) ([]model.CustomerAddress, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := r.db.Find(&customers).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching customers: %w", err)
	}
	return customers, nil
}

func (r *customerRepository) Save(customer *model.Customer) error {
	if err := r.db.Create(customer).Error; err != nil {
		return fmt.Errorf("error saving customer: %w", err)
	}
	return nil
}

func (r *customerRepository) GetCustomerAddressesByUserID(userID int) ([]model.CustomerAddress, error) {
	var addresses []model.CustomerAddress

	query := `
        SELECT 
            address_id, user_id, recipient_name, phone_number, address_line, 
            city, province, postal_code, latitude, longitude, is_default, created_at
        FROM customer_addresses
        WHERE user_id = ?
    `

	result := r.db.Raw(query, userID).Scan(&addresses)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching customer addresses: %w", result.Error)
	}

	return addresses, nil
}
