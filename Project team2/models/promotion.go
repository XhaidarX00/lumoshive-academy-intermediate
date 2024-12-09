package models

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	ProductName datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"product_name"`
	Type        datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"type"`
	Description string         `gorm:"type:text" json:"description"`
	Discount    datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"discount"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	Quota       int            `gorm:"default:0" json:"quota"`
	Status      bool           `gorm:"default:false" json:"status"`
	Published   bool           `gorm:"default:false" json:"published"`
}

// TableName overrides the table name
func (Promotion) TableName() string {
	return "promotions"
}

// BeforeCreate validates date range
func (p *Promotion) BeforeCreate(tx *gorm.DB) (err error) {
	if p.EndDate.Before(p.StartDate) {
		return fmt.Errorf("end date must be after start date")
	}
	return
}
