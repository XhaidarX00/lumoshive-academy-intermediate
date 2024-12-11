package controllers

import (
	"net/http"
	"project_auth_jwt/helper"
	"project_auth_jwt/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReportController struct {
	ReportService services.Service
	Log           *zap.Logger
}

func NewReportController(reportService services.Service, log *zap.Logger) *ReportController {
	return &ReportController{
		ReportService: reportService,
		Log:           log,
	}
}

func (ctrl *ReportController) GenerateReport(c *gin.Context) {
	minute, err := strconv.Atoi(c.Param("minute"))
	if err != nil {
		ctrl.Log.Error("Invalid minutes", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid minutes", http.StatusBadRequest)
		c.Abort()
		return
	}

	err = ctrl.ReportService.Report.GenerateReport(minute)
	if err != nil {
		ctrl.Log.Error("Failed Generate Report", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed Generate Report", http.StatusInternalServerError)
		c.Abort()
		return
	}

	helper.ResponseOK(c, nil, "Report generated successfully", http.StatusCreated)
}
