package dashboardcontroller

import (
	"bytes"
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/models"
	"dashboard-ecommerce-team2/service"
	dashboardservice "dashboard-ecommerce-team2/service/dashboard"

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var _ dashboardservice.DashboardService = (*dashboardservice.MockDashboardService)(nil)

func setupDashboardController(mockService dashboardservice.DashboardService) *DashboardController {
	mockLogger, _ := zap.NewProduction()

	return &DashboardController{
		Service: service.Service{Dashboard: mockService},
		Log:     mockLogger,
	}
}

func setupTest(route string) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, route, nil)
	c.Request.Header.Set("Content-Type", "application/json")

	return w, c
}

func TestGetSummaryController(t *testing.T) {
	t.Run("Success get summary data", func(t *testing.T) {
		// Expected summary data
		summary := &models.Summary{
			TotalUser:    27,
			TotalSales:   199.9,
			TotalOrder:   127,
			TotalProduct: 99,
		}

		// Setup mock service
		mockService := &dashboardservice.MockDashboardService{}
		mockService.On("GetDashboardSummary").Return(summary, nil).Once()

		// Initialize the controller with the mock service
		ctrl := setupDashboardController(mockService)
		w, c := setupTest("/dashboard/summary")

		// Call the controller method
		ctrl.GetSummaryController(c)

		// Verify the response
		assert.Equal(t, http.StatusOK, w.Code)

		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "dashboard summary successfully retrieved", response.Message)

		// Verify mock service expectations
		mockService.AssertExpectations(t)
	})

	// sub test failure get summary
	t.Run("Error getting summary data", func(t *testing.T) {
		// Mock the error response
		mockService := &dashboardservice.MockDashboardService{}
		mockService.On("GetDashboardSummary").Return(nil, errors.New("Error getting summary")).Once()

		// Initialize the controller with the mock service
		ctrl := setupDashboardController(mockService)
		w, c := setupTest("/dashboard/summary")

		// Call the controller method
		ctrl.GetSummaryController(c)

		// Verify the response
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Parse the response body
		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)

		// Expected error response
		assert.Equal(t, "Error getting summary", response.Message)

		// Verify mock service expectations
		mockService.AssertExpectations(t)
	})
}

func TestCurrentMonthEarningController(t *testing.T) {
	// sub test success get current month earning
	t.Run("Success getting current month earning", func(t *testing.T) {
		thisMonthEarning := &models.Revenue{
			Month:        "December",
			TotalEarning: 127.9,
		}

		mockService := &dashboardservice.MockDashboardService{}
		mockService.On("CurrentMonthEarning").Return(thisMonthEarning, nil).Once()

		ctrl := setupDashboardController(mockService)
		w, c := setupTest("/dashboard/current-month-earning")

		// Call the controller method
		ctrl.CurrentMonthEarningController(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "current month earnings successfully retrieved", response.Message)

		// Verify mock service expectations
		mockService.AssertExpectations(t)
	})

	// sub test failure get current month earning
	t.Run("Failure getting current month earning", func(t *testing.T) {
		mockService := &dashboardservice.MockDashboardService{}
		mockService.On("CurrentMonthEarning").Return(nil, errors.New("Error getting current month earning")).Once()

		ctrl := setupDashboardController(mockService)
		w, c := setupTest("/dashboard/current-month-earning")

		// Call the controller method
		ctrl.CurrentMonthEarningController(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response helper.HTTPResponse
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "Error getting current month earning", response.Message)

		// Verify mock service expectations
		mockService.AssertExpectations(t)
	})
}

func TestRenevueChartController(t *testing.T) {
	t.Run("Success generating revenue chart", func(t *testing.T) {
		// Mock the binary data for the chart
		mockChartData := []byte("<svg>chart</svg>")
		mockService := &dashboardservice.MockDashboardService{}
		mockService.On("GenerateRenevueChart").Return(bytes.NewBuffer(mockChartData), nil).Once()

		// Initialize the controller with the mock service
		ctrl := setupDashboardController(mockService)
		w, c := setupTest("/dashboard/revenue-chart")

		// Call the controller method
		ctrl.RenevueChartController(c)

		// Verify the response
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "text/html", w.Header().Get("Content-Type"))
		assert.Equal(t, string(mockChartData), w.Body.String())

		// Verify mock service expectations
		mockService.AssertExpectations(t)
	})

	t.Run("Error generating revenue chart", func(t *testing.T) {
		// Mock the error response
		mockService := &dashboardservice.MockDashboardService{}
		mockService.On("GenerateRenevueChart").Return(nil, errors.New("Error generating chart")).Once()

		// Initialize the controller with the mock service
		ctrl := setupDashboardController(mockService)
		w, c := setupTest("/dashboard/revenue-chart")

		// Call the controller method
		ctrl.RenevueChartController(c)

		// Verify the response
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		response := helper.HTTPResponse{}
		err := json.NewDecoder(w.Body).Decode(&response)
		assert.NoError(t, err)

		// Expected error response
		assert.Equal(t, "Error generating revenue chart", response.Message)

		// Verify mock service expectations
		mockService.AssertExpectations(t)
	})
}
