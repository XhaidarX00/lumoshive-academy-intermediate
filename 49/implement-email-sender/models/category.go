package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int             `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name      string          `json:"name" gorm:"size:50;unique;not null" binding:"required" example:"Electronics"`
	Variant   json.RawMessage `json:"variant" gorm:"type:jsonb;default:'{}'" example:"{\"color\":\"red\", \"size\":\"medium\"}"`
	Icon      string          `json:"icon" form:"icon" gorm:"size:255;not null" binding:"omitempty" example:"/icon/category.png"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func CategorySeed() []Category {
	return []Category{
		{
			Name:      "Electronics",
			Icon:      "/icon/electronics.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Fashion",
			Icon:      "/icon/fashion.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Home & Kitchen",
			Icon:      "/icon/home-kitchen.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Books",
			Icon:      "/icon/books.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Toys & Games",
			Icon:      "/icon/toys-games.png",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
