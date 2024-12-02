package model

import "time"

type CustomerAddress struct {
	AddressID     int       `json:"address_id"`
	UserID        int       `json:"user_id"`
	RecipientName string    `json:"recipient_name"`
	PhoneNumber   string    `json:"phone_number"`
	AddressLine   string    `json:"address_line"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	PostalCode    string    `json:"postal_code"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
}

type SellerAddress struct {
	AddressID     int       `json:"address_id"`
	UserID        int       `json:"user_id"`
	OrderID       int       `json:"order_id"`
	RecipientName string    `json:"recipient_name"`
	PhoneNumber   string    `json:"phone_number"`
	AddressLine   string    `json:"address_line"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	PostalCode    string    `json:"postal_code"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	CreatedAt     time.Time `json:"created_at"`
}
