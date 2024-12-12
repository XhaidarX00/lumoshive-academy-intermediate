package dashboardcontroller

import (
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DashboardController struct {
	Service service.Service
	Log     *zap.Logger
}

func NewDashboardController(service service.Service, log *zap.Logger) *DashboardController {
	return &DashboardController{
		Service: service,
		Log:     log,
	}
}

// @Summary Get dashboard summary
// @Description Retrieve a summary of the dashboard
// @Tags Dashboard
// @Produce json
// @Success 200 {object} helper.HTTPResponse "dashboard summary successfully retrieved"
// @Failure 500 {object} helper.HTTPResponse "Error getting summary"
// @Router /dashboard/summary [get]
func (ctrl *DashboardController) GetSummaryController(c *gin.Context) {
	summary, err := ctrl.Service.Dashboard.GetDashboardSummary()
	if err != nil {
		ctrl.Log.Error("Error getting summary", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error getting summary", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, summary, "dashboard summary successfully retrieved", http.StatusOK)
}

// @Summary Get current month earnings
// @Description Retrieve the earnings for the current month
// @Tags Dashboard
// @Produce json
// @Success 200 {object} helper.HTTPResponse "current month earnings successfully retrieved"
// @Failure 500 {object} helper.HTTPResponse"Error getting earnings"
// @Router /dashboard/current-month-earning [get]
func (ctrl *DashboardController) CurrentMonthEarningController(c *gin.Context) {
	earnings, err := ctrl.Service.Dashboard.CurrentMonthEarning()
	if err != nil {
		ctrl.Log.Error("Error getting current month earning", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error getting current month earning", http.StatusInternalServerError)
		return
	}

	helper.ResponseOK(c, earnings, "current month earnings successfully retrieved", http.StatusOK)
}

// @Summary Get revenue chart
// @Description Generate a revenue chart for the dashboard
// @Tags Dashboard
// @Produce text/html
// @Success 200 {string} string "<!DOCTYPE html><html><head><meta charset='utf-8'><title>Awesome go-echarts</title><script src='https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js'></script></head><body><div class='container'><div class='item' id='PvBdsRyYJxot' style='width:900px;height:500px;'></div></div><script type='text/javascript'>\"use strict\";let goecharts_PvBdsRyYJxot = echarts.init(document.getElementById('PvBdsRyYJxot'), 'white', { renderer: 'canvas' });let option_PvBdsRyYJxot = {\"color\":[\"#5470c6\",\"#91cc75\",\"#fac858\",\"#ee6666\",\"#73c0de\",\"#3ba272\",\"#fc8452\",\"#9a60b4\",\"#ea7ccc\"],\"legend\":{},\"series\":[{\"name\":\"Revenue\",\"type\":\"line\",\"smooth\":true,\"data\":[{\"value\":150.75},{\"value\":300},{\"value\":500.5},{\"value\":175.25},{\"value\":250},{\"value\":100.75},{\"value\":400.5},{\"value\":300.25},{\"value\":275},{\"value\":125.5},{\"value\":500},{\"value\":350}]}],\"title\":{\"text\":\"Monthly Revenue\"},\"toolbox\":{},\"tooltip\":{},\"xAxis\":[{\"name\":\"Month\",\"data\":[\"January  \",\"February \",\"March    \",\"April    \",\"May      \",\"June     \",\"July     \",\"August   \",\"September\",\"October  \",\"November \",\"December \"]}],\"yAxis\":[{\"name\":\"Revenue\"}]}goecharts_PvBdsRyYJxot.setOption(option_PvBdsRyYJxot);</script><style>.container {margin-top:30px; display: flex;justify-content: center;align-items: center;}.item {margin: auto;}</style></body></html>"
// @Failure 500 {object} helper.HTTPResponse "Error generating revenue chart"
// @Router /dashboard/revenue-chart [get]
func (ctrl *DashboardController) RenevueChartController(c *gin.Context) {
	revenueChart, err := ctrl.Service.Dashboard.GenerateRenevueChart()
	if err != nil {
		ctrl.Log.Error("Error generating revenue chart", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error generating revenue chart", http.StatusInternalServerError)
		return
	}
	c.Header("Content-Type", "text/html")
	c.Writer.Write(revenueChart.Bytes())
}

// @Summary Get best product list
// @Description Retrieve a list of the best-selling products
// @Tags Dashboard
// @Produce json
// @Success 200 {object} helper.HTTPResponse "best product list successfully retrieved"
// @Failure 500 {object} helper.HTTPResponse "Error getting best product list"
// @Router /dashboard/best-item-list [get]
func (ctrl *DashboardController) GetBestProductListController(c *gin.Context) {
	bestProduct, err := ctrl.Service.Dashboard.GetBestItemList()
	if err != nil {
		ctrl.Log.Error("Error getting best product list", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Error getting best product list", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, bestProduct, "best product list successfully retrieved", http.StatusOK)
}
