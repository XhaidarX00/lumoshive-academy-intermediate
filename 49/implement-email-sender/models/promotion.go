package models

import (
	"time"

	"gorm.io/gorm"
)

// Promotion represents the promotion data model
type Promotion struct {
	ID          int             `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name        string          `json:"name" gorm:"size:255;not null" binding:"required" example:"Holiday Sale"`
	ProductName JSONB           `json:"product_name" gorm:"type:jsonb;default:'[]'" binding:"omitempty" swaggertype:"array,string" example:"[\"Smartphone\", \"Laptop\"]"`
	Type        JSONB           `json:"type" gorm:"type:jsonb;default:'[]'" binding:"omitempty" swaggertype:"array,string" example:"[\"discount\", \"bundle\"]"`
	Description string          `json:"description" gorm:"size:500" binding:"omitempty" example:"Special holiday discounts"`
	Discount    JSONB           `json:"discount" gorm:"type:jsonb;default:'[]'" binding:"omitempty" swaggertype:"array,string" example:"[{\"value\":10,\"type\":\"percentage\"}]"`
	StartDate   time.Time       `json:"start_date" binding:"required" example:"2024-12-01"`
	EndDate     time.Time       `json:"end_date" binding:"required,gtfield=StartDate" example:"2024-12-31"`
	Quota       int             `json:"quota" gorm:"default:0" example:"100"`
	Status      bool            `json:"status" gorm:"default:false" example:"true"`
	Published   bool            `json:"published" gorm:"default:false" example:"true"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

// PromotionSeed returns a list of seed data for promotions
func PromotionSeed() []Promotion {
	return []Promotion{
		{
			Name:        "Winter Sale",
			ProductName: JSONB{"Coat", "Gloves"},
			Type:        JSONB{"discount"},
			Description: "Get up to 50% off winter essentials",
			Discount:    JSONB{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000},
			StartDate:   time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			Quota:       200,
			Status:      true,
			Published:   true,
		},
		{
			Name:        "Holiday Discount",
			ProductName: JSONB{"Smartphone", "Laptop"},
			Type:        JSONB{"bundle"},
			Description: "Bundle offer: Buy one smartphone, get a laptop for 30% off",
			Discount:    JSONB{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000},
			StartDate:   time.Date(2024, 12, 10, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2024, 12, 25, 23, 59, 59, 0, time.UTC),
			Quota:       150,
			Status:      true,
			Published:   true,
		},
		{
			Name:        "New Year Flash Sale",
			ProductName: JSONB{"Washing Machine", "Refrigerator"},
			Type:        JSONB{"discount"},
			Description: "Up to 40% off on select home appliances",
			Discount:    JSONB{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000},
			StartDate:   time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 1, 5, 23, 59, 59, 0, time.UTC),
			Quota:       100,
			Status:      true,
			Published:   true,
		},
		{
			Name:        "Back to School Sale",
			ProductName: JSONB{"School Bag", "Stationery"},
			Type:        JSONB{"bundle"},
			Description: "Buy a school bag and get a free set of stationery",
			Discount:    JSONB{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000},
			StartDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2024, 8, 15, 23, 59, 59, 0, time.UTC),
			Quota:       300,
			Status:      true,
			Published:   true,
		},
	}
}
