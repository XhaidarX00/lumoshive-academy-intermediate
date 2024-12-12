package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int         `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	CategoryID  int         `json:"category_id" form:"category_id" gorm:"not null" binding:"required,gt=0" example:"1"`
	Name        string      `json:"name" form:"name" gorm:"size:100;not null" binding:"required,min=3" example:"Smartphone"`
	CodeProduct string      `json:"code_product" form:"code_product" gorm:"size:100;unique;not null" binding:"required" example:"SPH-001"`
	Images      StringArray `json:"images" form:"file" gorm:"type:jsonb;default:'[]'" binding:"omitempty,dive,url" example:"[\"/images/smartphone1.png\"]" swaggerignore:"true"`
	Description string      `json:"description" form:"description" binding:"omitempty" example:"Latest smartphone with advanced features"`
	Stock       int         `json:"stock" form:"stock" gorm:"not null" binding:"required,gt=0" example:"50"`
	Price       float64     `json:"price" form:"price" gorm:"type:decimal(10,2);not null" binding:"required,gt=0" example:"699.99"`
	Published   bool        `json:"published" form:"published" gorm:"type:boolean;default:false" example:"true"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}
type ProductWithCategory struct {
	ID           int         `json:"id"`
	CategoryName string      `json:"category_name"`
	Name         string      `json:"name"`
	CodeProduct  string      `json:"code_product"`
	Images       StringArray `json:"images"`
	Description  string      `json:"description"`
	Stock        int         `json:"stock"`
	Price        float64     `json:"price"`
	Published    bool        `json:"published"`
}

type ProductID struct {
	ID           int             `json:"id"`
	CategoryName string          `json:"category_name"`
	Variant      json.RawMessage `json:"variant"`
	Name         string          `json:"name"`
	CodeProduct  string          `json:"code_product"`
	Images       StringArray     `json:"images"`
	Description  string          `json:"description"`
	Stock        int             `json:"stock"`
	Price        float64         `json:"price"`
	Published    bool            `json:"published"`
}

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}
	byteValue, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringArray")
	}
	return json.Unmarshal(byteValue, s)
}

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (v *Product) BeforeSave(tx *gorm.DB) (err error) {
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()

	return nil
}

type Stock struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id" binding:"required,gt=0"`
	Type      string    `gorm:"size:10;check:type IN ('in','out')" json:"type"`
	Quantity  int       `json:"quantity" binding:"required,gt=0"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (v *Stock) BeforeSave(tx *gorm.DB) (err error) {
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()
	return nil
}

type StockRequest struct {
	ProductID uint      `json:"product_id" binding:"required,gt=0"`
	Type      string    `json:"type" binding:"required,oneof=in out"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type StockResponse struct {
	ID           uint            `json:"id"`
	ProductID    uint            `json:"product_id"`
	Type         string          `json:"description"`
	Quantity     int             `json:"quantity"`
	ProductName  string          `json:"product_name"`
	Variant      json.RawMessage `json:"variant"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type BestProduct struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Total     int    `json:"total"`
}

func ProductSeed() []Product {
	return []Product{
		{
			CategoryID:  1,
			Name:        "Smartphone",
			CodeProduct: "SPH-001",
			Images:      []string{"/images/smartphone1.png", "/images/smartphone2.png"},
			Description: "Latest smartphone with advanced features",
			Stock:       50,
			Price:       699.99,
			Published:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			CategoryID:  2,
			Name:        "Laptop",
			CodeProduct: "LAP-002",
			Images:      []string{"/images/laptop1.png", "/images/laptop2.png"},
			Description: "Powerful laptop with high performance",
			Stock:       30,
			Price:       1299.99,
			Published:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			CategoryID:  1,
			Name:        "Smartwatch",
			CodeProduct: "SW-003",
			Images:      []string{"/images/smartwatch1.png", "/images/smartwatch2.png"},
			Description: "Smartwatch with health monitoring features",
			Stock:       100,
			Price:       199.99,
			Published:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			CategoryID:  3,
			Name:        "Headphones",
			CodeProduct: "HP-004",
			Images:      []string{"/images/headphones1.png", "/images/headphones2.png"},
			Description: "Wireless headphones with noise cancellation",
			Stock:       150,
			Price:       89.99,
			Published:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
}

func StockSeed() []Stock {
	return []Stock{
		{
			ProductID: 1,
			Type:      "in",
			Quantity:  10,
		},
		{
			ProductID: 2,
			Type:      "in",
			Quantity:  10,
		},
		{
			ProductID: 3,
			Type:      "out",
			Quantity:  10,
		},
		{
			ProductID: 4,
			Type:      "in",
			Quantity:  10,
		},
	}
}
