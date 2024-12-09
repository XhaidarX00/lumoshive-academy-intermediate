package models

import (
	"voucher_system/helper"
)

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name     string `json:"name,omitempty" gorm:"type:varchar(255);not null" binding:"required"`
	Email    string `json:"email,omitempty" gorm:"type:varchar(255);unique;not null" binding:"required,email"`
	Password string `json:"password,omitempty" gorm:"type:varchar(255);not null" binding:"required,min=8"`
}

func UserSeed() []User {
	return []User{
		{Name: "John Doe", Email: "john.doe@example.com", Password: helper.HashPassword("password1234")},
		{Name: "Jane Smith", Email: "jane.smith@example.com", Password: helper.HashPassword("password1245")},
		{Name: "Alice Johnson", Email: "alice.johnson@example.com", Password: helper.HashPassword("password1256")},
		{Name: "Bob Brown", Email: "bob.brown@example.com", Password: helper.HashPassword("password1278")},
		{Name: "Charlie Davis", Email: "charlie.davis@example.com", Password: helper.HashPassword("password1298")},
	}
}

// type Banner struct {
// 	Image       string
// 	Title       string
// 	Type        []string
// 	PathPage    string
// 	ReleaseDate *time.Time
// 	EndDate     *time.Time
// 	Published   bool
// }
