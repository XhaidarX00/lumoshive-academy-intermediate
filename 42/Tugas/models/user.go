package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Points    int       `gorm:"type:int;default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relationships
	Redeems []Redeem `gorm:"foreignKey:UserID"`
}

type Auth struct {
	ID    string
	Token string
}

type Login struct {
	Email    string `example:"jane.smith@example.com"`
	Password string `example:"securepassword"`
}
