package repositories

import (
	"log"
	"project_auth_jwt/models"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReportRepositoryInterface interface {
	GetOrdersFromLastMinutes(minutes int) ([]models.Report, error)
	GetReports(minutes int) ([]models.Report, error)
}

type ReportRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewReportRepository(db *gorm.DB, log *zap.Logger) ReportRepositoryInterface {
	return &ReportRepository{
		DB:  db,
		Log: log,
	}
}

func (r *ReportRepository) GetOrdersFromLastMinutes(minutes int) ([]models.Report, error) {
	var orders []models.Order
	var report []models.Report

	// Validasi input untuk memastikan minutes positif
	if minutes <= 0 {
		r.Log.Warn("Invalid minutes value, must be greater than 0")
		return nil, nil
	}

	// Hitung waktu sebelumnya berdasarkan input minutes
	timeAgo := time.Now().Add(-time.Duration(minutes) * time.Minute)

	// Query untuk mengambil data dalam waktu yang dinamis
	err := r.DB.Where("created_at >= ?", timeAgo).Find(&orders).Error
	if err != nil {
		r.Log.Error("Failed to fetch orders from last minutes", zap.Int("minutes", minutes), zap.Error(err))
		return nil, err
	}

	return report, nil
}

func (r *ReportRepository) GetReports(minutes int) ([]models.Report, error) {
	var orders []models.Order
	var users []models.User
	var reports []models.Report

	// Query untuk mengambil data orders
	orderQuery := r.DB.Find(&orders)
	if minutes > 0 {
		timeAgo := time.Now().Add(-time.Duration(minutes) * time.Minute)
		orderQuery = r.DB.Where("created_at >= ?", timeAgo).Find(&orders)
	}

	err := orderQuery.Error
	if err != nil {
		return nil, err
	}

	// Query untuk mengambil data users
	err = r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	log.Printf("%v\n", orders)
	// log.Printf("%v\n%v\n", orders, users)

	// Gabungkan data secara manual
	userMap := make(map[int]string)
	for _, user := range users {
		userMap[user.ID] = user.Name
	}

	for _, order := range orders {
		reports = append(reports, models.Report{
			ID:              uint(order.ID),
			CustomerName:    userMap[order.UserID],
			ShippingAddress: order.ShippingAddress,
			PaymentMethod:   order.PaymentMethod,
			Status:          order.Status,
			TotalAmount:     order.TotalAmount,
			CreatedAt:       order.CreatedAt,
		})
	}

	return reports, nil
}
