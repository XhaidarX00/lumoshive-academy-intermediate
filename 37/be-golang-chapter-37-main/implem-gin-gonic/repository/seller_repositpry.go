package repository

import (
	"fmt"
	"golang-chapter-37/implem-gin-gonic/model"

	"gorm.io/gorm"
)

type SellerRepository interface {
	GetSellerAddressesByUserID(userID int) (model.SellerAddress, error)
}

type sellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) SellerRepository {
	return &sellerRepository{db: db}
}

func (r *sellerRepository) GetSellerAddressesByUserID(userID int) (model.SellerAddress, error) {
	var addresses model.SellerAddress

	query := `
        SELECT 
			address_id, user_id, recipient_name, phone_number, address_line, 
			city, province, postal_code, latitude, longitude, created_at
		FROM seller_addresses
		WHERE user_id = ?
    `

	// Gunakan GORM untuk eksekusi query
	result := r.db.Raw(query, userID).Scan(&addresses)
	if result.Error != nil {
		return model.SellerAddress{}, fmt.Errorf("error fetching Seller addresses: %w", result.Error)
	}

	return addresses, nil
}
