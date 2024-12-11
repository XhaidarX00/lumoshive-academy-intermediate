package services

import (
	"project_auth_jwt/repositories"
	"project_auth_jwt/utils"
	"time"

	"go.uber.org/zap"
)

type ReportServiceInterface interface {
	GenerateReport(minutes int) error
}

type ReportService struct {
	ReportRepo repositories.Repository
	Log        *zap.Logger
}

func NewReportService(ReportRepo repositories.Repository, log *zap.Logger) ReportServiceInterface {
	return &ReportService{
		ReportRepo: ReportRepo,
		Log:        log,
	}
}

func (s *ReportService) GenerateReport(minutes int) error {
	orders, err := s.ReportRepo.Report.GetReports(minutes)
	if err != nil {
		s.Log.Error("Failed to generate report", zap.Error(err))
		return err
	}

	timestamp := time.Now().Format("20060102_150405")
	filePath := "reports/report_" + timestamp + ".xlsx"

	return utils.WriteReportsToExcel(filePath, orders)
}
