package models

import "time"

type Order struct {
	ID              int       `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	UserID          int       `json:"user_id" gorm:"not null" binding:"required" example:"1"`
	TotalAmount     float64   `json:"total_amount" gorm:"type:decimal(10,2);not null" binding:"required" example:"150.75"`
	PaymentMethod   string    `json:"payment_method" gorm:"size:20;not null" binding:"required" example:"credit_card"`
	ShippingAddress string    `json:"shipping_address" gorm:"size:255" binding:"omitempty" example:"123 Main St"`
	Status          string    `json:"status" gorm:"size:20;check:status IN ('pending','shipped','completed','canceled');default:'created'" binding:"required" example:"pending"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}

type OrderItem struct {
	ID        int     `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	OrderID   int     `json:"order_id" gorm:"not null" binding:"required" example:"1"`
	ProductID int     `json:"product_id" gorm:"not null" binding:"required" example:"101"`
	Quantity  int     `json:"quantity" gorm:"not null" binding:"required" example:"2"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);not null" binding:"required" example:"75.00"`
	Total     float64 `json:"total" gorm:"type:decimal(10,2);not null" binding:"required"`
}

type OrderDetailResponse struct {
	Order *Order      `json:"order"`
	Items []OrderItem `json:"items"`
}

func OrderSeed() []Order {
	return []Order{
		{
			UserID:          1,
			TotalAmount:     150.75,
			PaymentMethod:   "credit_card",
			ShippingAddress: "123 Main St",
			Status:          "pending",
			CreatedAt:       time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.January, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          2,
			TotalAmount:     300.00,
			PaymentMethod:   "paypal",
			ShippingAddress: "456 Elm St",
			Status:          "shipped",
			CreatedAt:       time.Date(2024, time.February, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.February, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          3,
			TotalAmount:     500.50,
			PaymentMethod:   "debit_card",
			ShippingAddress: "789 Oak St",
			Status:          "completed",
			CreatedAt:       time.Date(2024, time.March, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.March, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          4,
			TotalAmount:     175.25,
			PaymentMethod:   "debit_card",
			ShippingAddress: "123 Pine St",
			Status:          "canceled",
			CreatedAt:       time.Date(2024, time.April, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.April, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          5,
			TotalAmount:     250.00,
			PaymentMethod:   "bank_transfer",
			ShippingAddress: "321 Birch St",
			Status:          "pending",
			CreatedAt:       time.Date(2024, time.May, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.May, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          6,
			TotalAmount:     100.75,
			PaymentMethod:   "cash",
			ShippingAddress: "654 Maple St",
			Status:          "shipped",
			CreatedAt:       time.Date(2024, time.June, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.June, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          7,
			TotalAmount:     400.50,
			PaymentMethod:   "credit_card",
			ShippingAddress: "987 Cedar St",
			Status:          "completed",
			CreatedAt:       time.Date(2024, time.July, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.July, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          8,
			TotalAmount:     300.25,
			PaymentMethod:   "debit_card",
			ShippingAddress: "213 Ash St",
			Status:          "canceled",
			CreatedAt:       time.Date(2024, time.August, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.August, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          9,
			TotalAmount:     275.00,
			PaymentMethod:   "paypal",
			ShippingAddress: "546 Spruce St",
			Status:          "pending",
			CreatedAt:       time.Date(2024, time.September, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.September, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          10,
			TotalAmount:     125.50,
			PaymentMethod:   "cash",
			ShippingAddress: "879 Willow St",
			Status:          "shipped",
			CreatedAt:       time.Date(2024, time.October, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.October, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          11,
			TotalAmount:     500.00,
			PaymentMethod:   "bank_transfer",
			ShippingAddress: "210 Poplar St",
			Status:          "completed",
			CreatedAt:       time.Date(2024, time.November, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.November, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			UserID:          12,
			TotalAmount:     350.00,
			PaymentMethod:   "credit_card",
			ShippingAddress: "432 Sycamore St",
			Status:          "completed",
			CreatedAt:       time.Date(2024, time.December, 15, 12, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2024, time.December, 15, 12, 0, 0, 0, time.UTC),
		},
	}
}

func OrderItemSeed() []OrderItem {
	return []OrderItem{
		{
			OrderID:   1,
			ProductID: 1, // Smartphone
			Quantity:  2,
			Price:     699.99,
			Total:     1399.98,
		},
		{
			OrderID:   2,
			ProductID: 2, // Laptop
			Quantity:  3,
			Price:     1299.99,
			Total:     3899.97,
		},
		{
			OrderID:   3,
			ProductID: 3, // Smartwatch
			Quantity:  5,
			Price:     199.99,
			Total:     999.95,
		},
		{
			OrderID:   4,
			ProductID: 4, // Headphones
			Quantity:  1,
			Price:     89.99,
			Total:     89.99,
		},
		{
			OrderID:   5,
			ProductID: 1, // Smartphone
			Quantity:  3,
			Price:     699.99,
			Total:     2099.97,
		},
		{
			OrderID:   6,
			ProductID: 2, // Laptop
			Quantity:  2,
			Price:     1299.99,
			Total:     2599.98,
		},
		{
			OrderID:   7,
			ProductID: 3, // Smartwatch
			Quantity:  4,
			Price:     199.99,
			Total:     799.96,
		},
		{
			OrderID:   8,
			ProductID: 4, // Headphones
			Quantity:  3,
			Price:     89.99,
			Total:     269.97,
		},
		{
			OrderID:   9,
			ProductID: 1, // Smartphone
			Quantity:  1,
			Price:     699.99,
			Total:     699.99,
		},
		{
			OrderID:   10,
			ProductID: 2, // Laptop
			Quantity:  1,
			Price:     1299.99,
			Total:     1299.99,
		},
		{
			OrderID:   11,
			ProductID: 3, // Smartwatch
			Quantity:  3,
			Price:     199.99,
			Total:     599.97,
		},
		{
			OrderID:   12,
			ProductID: 4, // Headphones
			Quantity:  2,
			Price:     89.99,
			Total:     179.98,
		},
	}
}
