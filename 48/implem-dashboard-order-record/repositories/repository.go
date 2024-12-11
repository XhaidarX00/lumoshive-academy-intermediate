package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	User   UserRepositoryInterface
	Order  OrderRepositoryInterface
	Report ReportRepositoryInterface
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		User:   NewUserRepository(db, log),
		Order:  NewOrderRepository(db, log),
		Report: NewReportRepository(db, log),
	}
}