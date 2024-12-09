package repository

import (
	"voucher_system/repository/banner"
	managementvoucher "voucher_system/repository/management_voucher"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	User    UserRepository
	Manage  managementvoucher.ManagementVoucherInterface
	Voucher VoucherRepository
	Redeem  RedeemRepository
	History HistoryRepository
	Banner  banner.BannerRepositoryInterface
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		User:    NewUserRepository(db, log),
		Manage:  managementvoucher.NewManagementVoucherRepo(db, log),
		Voucher: NewVoucherRepository(db, log),
		Redeem:  NewRedeemRepository(db, log),
		History: NewHistoryRepository(db, log),
		Banner:  banner.NewBannerRepository(db, log),
	}
}
