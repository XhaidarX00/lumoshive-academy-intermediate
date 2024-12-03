package repository

import (
	"project-voucher-team3/database"

	"gorm.io/gorm"
)

type Repository struct {
	User    UserRepository
	Voucher VoucherRepository
	Reedem  ReedemRepository
	Auth    AuthRepository
}

func NewRepository(db *gorm.DB, rdb database.Cacher) Repository {
	return Repository{
		User:    *NewUserRepository(db),
		Voucher: *NewVoucherRepository(db),
		Reedem:  *NewReedemRepository(db),
		Auth:    *NewAuthRepository(rdb, db),
	}
}
