package models

import "time"

type Report struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	CustomerName    string    `gorm:"size:100;not null" json:"customer_name" binding:"required"`
	ProductName     string    `gorm:"size:100;not null" json:"product_name" binding:"required"`
	Quantity        int       `gorm:"not null" json:"quantity" binding:"required" example:"1"`
	TotalAmount     float64   `gorm:"type:decimal(10,2);not null" json:"total_amount" example:"150.75"`
	PaymentMethod   string    `gorm:"size:20;not null" json:"payment_method" binding:"required" example:"credit_card"`
	ShippingAddress string    `gorm:"size:255" json:"shipping_address,omitempty" example:"123 Main St"`
	Status          string    `gorm:"size:20;not null;check:status IN ('pending','shipped','completed','canceled');default:'created'" json:"status" binding:"required" example:"pending"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
}
