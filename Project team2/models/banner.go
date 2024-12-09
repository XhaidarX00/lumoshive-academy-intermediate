package models

import (
	"time"
	"voucher_system/helper"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Banner struct {
	gorm.Model
	Image       string         `gorm:"type:varchar(255);not null" json:"image"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Type        datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"type"`
	Description string         `json:"description,omitempty" gorm:"type:text;not null"`
	PathPage    string         `gorm:"type:varchar(255);not null" json:"path_page"`
	ReleaseDate *time.Time     `json:"release_date"`
	EndDate     *time.Time     `json:"end_date"`
	Published   bool           `gorm:"default:false" json:"published"`
}

// TableName overrides the table name
func (Banner) TableName() string {
	return "banners"
}

// BeforeCreate sets default values
func (b *Banner) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if b.ReleaseDate == nil {
		b.ReleaseDate = &now
	}
	if b.EndDate == nil {
		endDate := now.AddDate(0, 1, 0)
		b.EndDate = &endDate
	}
	return
}

func BannerSeed() []Banner {
	now := time.Now()
	return []Banner{
		{
			Image:       "summer_sale.jpg",
			Title:       "Summer Sale 2024",
			Type:        datatypes.JSON(`[{"type": "seasonal", "category": "sale"}]`),
			Description: "Get up to 50% off on summer products!",
			PathPage:    "/summer-sale",
			Published:   true,
			ReleaseDate: &now,
			EndDate:     helper.PointerToTime(now.AddDate(0, 1, 0)), // 1 month later
		},
		{
			Image:       "black_friday.jpg",
			Title:       "Black Friday Deals",
			Type:        datatypes.JSON(`[{"type": "holiday", "category": "discount"}]`),
			Description: "Exclusive Black Friday discounts on all items!",
			PathPage:    "/black-friday",
			Published:   true,
			ReleaseDate: helper.PointerToTime(now.AddDate(0, -2, 0)), // 2 months ago
			EndDate:     helper.PointerToTime(now.AddDate(0, -1, 0)), // 1 month ago
		},
		{
			Image:       "new_year.jpg",
			Title:       "New Year Celebration",
			Type:        datatypes.JSON(`[{"type": "event", "category": "celebration"}]`),
			Description: "Celebrate the new year with amazing offers!",
			PathPage:    "/new-year",
			Published:   false,
			ReleaseDate: helper.PointerToTime(now.AddDate(0, 1, 0)), // 1 month later
			EndDate:     helper.PointerToTime(now.AddDate(0, 2, 0)), // 2 months later
		},
		{
			Image:       "spring_collection.jpg",
			Title:       "Spring Collection Launch",
			Type:        datatypes.JSON(`[{"type": "seasonal", "category": "new collection"}]`),
			Description: "Discover the new spring collection now!",
			PathPage:    "/spring-collection",
			Published:   true,
			ReleaseDate: helper.PointerToTime(now.AddDate(0, -3, 0)), // 3 months ago
			EndDate:     helper.PointerToTime(now.AddDate(0, 0, -7)), // 7 days ago
		},
		{
			Image:       "winter_clearance.jpg",
			Title:       "Winter Clearance",
			Type:        datatypes.JSON(`[{"type": "seasonal", "category": "clearance"}]`),
			Description: "Last chance to grab winter items at a discount!",
			PathPage:    "/winter-clearance",
			Published:   false,
			ReleaseDate: helper.PointerToTime(now.AddDate(0, -2, -10)), // 2 months and 10 days ago
			EndDate:     helper.PointerToTime(now.AddDate(0, -1, 0)),   // 1 month ago
		},
	}
}
