package dashboardservice

import (
	"bytes"
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
)

type MockDashboardService struct {
	mock.Mock
	DashboardService DashboardService
}

func (m *MockDashboardService) GetDashboardSummary() (*models.Summary, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Summary), args.Error(1)
}

func (m *MockDashboardService) CurrentMonthEarning() (*models.Revenue, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Revenue), args.Error(1)
}

func (m *MockDashboardService) GenerateRenevueChart() (*bytes.Buffer, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*bytes.Buffer), args.Error(1)
}

func (m *MockDashboardService) GetBestItemList() ([]models.BestProduct, error) {
	args := m.Called()
	return args.Get(0).([]models.BestProduct), args.Error(1)
}

// NewMockDashboardService is a helper function to create a mock instance
func NewMockDashboardService() *MockDashboardService {
	return &MockDashboardService{}
}
