package models

import "time"

type Redeem struct {
	Base
	UserID      int       `json:"user_id" gorm:"type:integer"`
	VoucherID   int       `gorm:"type:integer;index" json:"-"` // Foreign key to Voucher
	VoucherCode string    `gorm:"type:varchar(255)" json:"voucher_code"`
	RedeemDate  time.Time `gorm:"type:timestamp" json:"redeem_date"`

	// Relationships
	Voucher Voucher `gorm:"foreignKey:VoucherID" json:"voucher"`
}
