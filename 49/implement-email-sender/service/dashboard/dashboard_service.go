package dashboardservice

import (
	"bytes"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/repository"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"go.uber.org/zap"
)

type DashboardService interface {
	GetDashboardSummary() (*models.Summary, error)
	CurrentMonthEarning() (*models.Revenue, error)
	GenerateRenevueChart() (*bytes.Buffer, error)
	GetBestItemList() ([]models.BestProduct, error)
}

type dashboardService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// GetDashboardSummary implements DashboardService.
func (d *dashboardService) GetDashboardSummary() (*models.Summary, error) {
	countCustomer, err := d.Repo.User.CountCustomer()
	if err != nil {
		return nil, err
	}

	countTotalOrder, err := d.Repo.Order.CountOrder()
	if err != nil {
		return nil, err
	}
	countTotalProduct, err := d.Repo.Product.CountProduct()
	if err != nil {
		return nil, err
	}
	totalPriceOrder, err := d.Repo.Order.CountTotalPriceOrder()
	if err != nil {
		return nil, err
	}
	summary := &models.Summary{
		TotalUser:    countCustomer,
		TotalOrder:   countTotalOrder,
		TotalProduct: countTotalProduct,
		TotalSales:   totalPriceOrder,
	}
	return summary, nil
}

// CurrentMonthEarning implements DashboardService.
func (d *dashboardService) CurrentMonthEarning() (*models.Revenue, error) {
	// Fetch earnings for each month
	monthlyEarning, err := d.Repo.Order.GetEarningEachMonth()
	if err != nil {
		d.Log.Error("Error getting monthly earning", zap.Error(err))
		return nil, err
	}

	// Ensure monthlyEarning is not nil or empty
	if monthlyEarning == nil || len(monthlyEarning) == 0 {
		d.Log.Warn("No monthly earning data available")
		return nil, nil
	}

	// Trim spaces from month names to ensure comparison works
	for i := range monthlyEarning {
		monthlyEarning[i].Month = strings.TrimSpace(monthlyEarning[i].Month)
	}

	d.Log.Info("Monthly earnings", zap.Any("Earnings", monthlyEarning))

	// Get the current month's name
	currentMonth := time.Now().Format("January")

	// Initialize a variable to store the current month's earnings
	var thisMonthEarning *models.Revenue

	// Ensure the list has an entry for each month from January to December
	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	// Create a map for easy lookup
	earningsMap := make(map[string]float64)
	for _, earning := range monthlyEarning {
		earningsMap[earning.Month] = earning.TotalEarning
	}

	// Iterate over months to find the current month and ensure all months are accounted for
	for _, month := range months {
		// If there's no data for a specific month, set its earnings to 0
		if _, exists := earningsMap[month]; !exists {
			earningsMap[month] = 0
		}

		if month == currentMonth {
			thisMonthEarning = &models.Revenue{
				Month:        month,
				TotalEarning: earningsMap[month],
			}
			break
		}
	}

	// Handle case where there's no data for the current month
	if thisMonthEarning == nil {
		d.Log.Warn("No earning data found for the current month", zap.String("Month", currentMonth))
		return nil, nil
	}

	return thisMonthEarning, nil
}

// GetBestItemList implements DashboardService.
func (d *dashboardService) GetBestItemList() ([]models.BestProduct, error) {
	return d.Repo.Product.CountEachProduct()
}

// RenevueChart implements DashboardService.
func (d *dashboardService) GenerateRenevueChart() (*bytes.Buffer, error) {
	monthlyEarning, err := d.Repo.Order.GetEarningEachMonth()
	if err != nil {
		d.Log.Error("Error getting monthly earning", zap.Error(err))
		return nil, err
	}
	d.Log.Info("Monthly earning", zap.Any("Earning", monthlyEarning))
	// Prepare data for the chart
	xValues := make([]string, len(monthlyEarning))
	yValues := make([]opts.LineData, len(monthlyEarning))

	for i, entry := range monthlyEarning {
		xValues[i] = entry.Month
		yValues[i] = opts.LineData{Value: entry.TotalEarning}
	}

	line := charts.NewLine()

	// Set chart options
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Monthly Revenue"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Month"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Revenue"}),
	)
	// Set X-axis and add series
	line.SetXAxis(xValues).
		AddSeries("Revenue", yValues).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)})) // Properly cast 'true' to opts.Bool

	// Render the chart into a buffer
	buffer := new(bytes.Buffer)
	if err := line.Render(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

func NewDashboardService(repo repository.Repository, log *zap.Logger) DashboardService {
	return &dashboardService{Repo: repo, Log: log}
}
