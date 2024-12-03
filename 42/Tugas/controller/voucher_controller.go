package controller

import (
	"net/http"
	"project-voucher-team3/models"
	"project-voucher-team3/service"
	"project-voucher-team3/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoucherController struct {
	service service.VoucherService
	logger  *zap.Logger
}

func NewVoucherController(service service.VoucherService, logger *zap.Logger) *VoucherController {
	return &VoucherController{service, logger}
}

// @Tags Voucher
// @Summary Validate a voucher
// @Description Validate the provided voucher and return its status.
// @Success 200 {object} utils.ResponseSuccess "Validation result"
// @Failure 400 {object} utils.ResponseFailed "Invalid voucher details"
// @Failure 500 {object} utils.ResponseFailed "Internal server error"
// @Security TokenAuth
// @Security IDKeyAuth
// @Param token header string true "Authentication token"
// @Param IDKey header string true "Unique user key"
// @Router /voucher/validate [get]
func (ctrl *VoucherController) ValidateVoucherController(c *gin.Context) {
	var voucherInput models.VoucherDTO
	if err := c.ShouldBindJSON(&voucherInput); err != nil {
		ctrl.logger.Error("Failed to bind voucher data", zap.Error(err))
		utils.ResponseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	validateResult, err := ctrl.service.ValidateVoucher(voucherInput)
	if err != nil {
		ctrl.logger.Error("Invalid voucher data", zap.Error(err))
		utils.ResponseError(c, "INVALID_DATA", err.Error(), http.StatusBadRequest)
		return
	}

	ctrl.logger.Info("Voucher data validated successfully")
	utils.ResponseOK(c, validateResult, "Voucher data validated successfully")
}

// @Tags Voucher
// @Summary Create a new voucher
// @Description Create a new voucher with the specified details.
// @Param voucher body map[string]interface{} true "Voucher data"
// @Success 201 {object} utils.ResponseSuccess "Voucher created"
// @Failure 400 {object} utils.ResponseFailed "Invalid voucher data"
// @Failure 500 {object} utils.ResponseFailed "Internal server error"
// @Security TokenAuth
// @Security IDKeyAuth
// @Param token header string true "Authentication token"
// @Param IDKey header string true "Unique user key"
// @Router /voucher [post]
func (ctrl *VoucherController) CreateVoucher(c *gin.Context) {
	var voucher models.Voucher
	var err error

	if err = c.ShouldBindJSON(&voucher); err != nil {
		ctrl.logger.Error("Failed to create voucher", zap.Error(err))
		utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	// voucher.StartDate, err = utils.TimeDateParse(voucher.StartDate.String())
	// if err != nil {
	// 	ctrl.logger.Error("Failed to create voucher", zap.Error(err))
	// 	utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// voucher.EndDate, err = utils.TimeDateParse(voucher.EndDate.String())
	// if err != nil {
	// 	ctrl.logger.Error("Failed to create voucher", zap.Error(err))
	// 	utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
	// 	return
	// }

	if err := ctrl.service.CreateVoucher(&voucher); err != nil {
		ctrl.logger.Error("Failed to create voucher", zap.Error(err))
		utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Create voucher successfully")
	utils.ResponseOK(c, voucher, "Create voucher successfully")
}

// @Tags Voucher
// @Summary Delete a voucher
// @Description Delete a voucher by its ID.
// @Param id path string true "Voucher ID"
// @Success 204 {object} utils.ResponseSuccess "Voucher deleted"
// @Failure 404 {object} utils.ResponseFailed "Voucher not found"
// @Failure 500 {object} utils.ResponseFailed "Internal server error"
// @Security TokenAuth
// @Security IDKeyAuth
// @Param token header string true "Authentication token"
// @Param IDKey header string true "Unique user key"
// @Router /voucher/{id} [delete]
func (ctrl *VoucherController) DeleteVoucher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Failed to delete voucher", zap.Error(err))
		utils.ResponseError(c, "DELETE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.DeleteVoucher(id); err != nil {
		ctrl.logger.Error("Data not found", zap.Error(err))
		utils.ResponseError(c, "DATA_NOT_FOUND", err.Error(), http.StatusNotFound)
		return
	}

	ctrl.logger.Info("Voucher deleted successfully")
	utils.ResponseOK(c, nil, "Voucher deleted successfully")
}

// @Tags Voucher
// @Summary Update a voucher
// @Description Update details of a voucher by its ID.
// @Param id path string true "Voucher ID"
// @Param voucher body map[string]interface{} true "Updated voucher data"
// @Success 200 {object} utils.ResponseSuccess "Voucher updated"
// @Failure 404 {object} utils.ResponseFailed "Voucher not found"
// @Failure 500 {object} utils.ResponseFailed "Internal server error"
// @Security TokenAuth
// @Security IDKeyAuth
// @Param token header string true "Authentication token"
// @Param IDKey header string true "Unique user key"
// @Router /voucher/{id} [put]
func (ctrl *VoucherController) UpdateVoucher(c *gin.Context) {
	var voucher models.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		ctrl.logger.Error("Failed to update voucher", zap.Error(err))
		utils.ResponseError(c, "UPDATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.UpdateVoucher(&voucher); err != nil {
		ctrl.logger.Error("Failed to update voucher", zap.Error(err))
		utils.ResponseError(c, "UPDATE_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Voucher update successfully")
	utils.ResponseOK(c, voucher, "Voucher update successfully")
}

// @Tags Voucher
// @Summary List all vouchers
// @Description Retrieve a list of all vouchers.
// @Success 200 {array} utils.ResponseSuccess "List of vouchers"
// @Failure 500 {object} utils.ResponseFailed "Internal server error"
// @Security TokenAuth
// @Security IDKeyAuth
// @Router /voucher [get]
func (ctrl *VoucherController) GetVouchers(c *gin.Context) {
	filter := make(map[string]interface{})
	if c.Query("voucher_code") != "" {
		filter["voucher_code"] = c.Query("voucher_code")
	}
	if c.Query("voucher_type") != "" {
		filter["voucher_type"] = c.Query("voucher_type")
	}

	vouchers, err := ctrl.service.GetVouchers(filter)
	if err != nil {
		ctrl.logger.Error("Failed to get voucher", zap.Error(err))
		utils.ResponseError(c, "GET_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	var response []models.VoucherWithStatus
	for _, voucher := range vouchers {
		response = append(response, models.VoucherWithStatus{
			Voucher:  voucher,
			IsActive: voucher.IsActive(),
		})
	}

	ctrl.logger.Info("Get voucher successfully")
	utils.ResponseOK(c, response, "Get voucher successfully")
}

// @Tags Voucher
// @Summary Get vouchers by minimum rate point
// @Description Retrieve vouchers with a minimum specified rate point.
// @Param ratePoint path int true "Minimum Rate Point"
// @Success 200 {array} utils.ResponseSuccess "List of vouchers"
// @Failure 400 {object} utils.ResponseFailed "Invalid rate point"
// @Failure 500 {object} utils.ResponseFailed "Internal server error"
// @Security TokenAuth
// @Security IDKeyAuth
// @Param token header string true "Authentication token"
// @Param IDKey header string true "Unique user key"
// @Router /voucher/point/{ratePoint} [get]
func (ctrl *VoucherController) GetVoucherWithMinRatePoint(c *gin.Context) {
	ratePoint, err := strconv.Atoi(c.Param("ratePoint"))
	if err != nil {
		ctrl.logger.Error("Failed to parse ratePoint", zap.Error(err))
		utils.ResponseError(c, "GET_VOUCHER_ERROR", "Invalid ratePoint parameter", http.StatusBadRequest)
		return
	}

	vouchers, err := ctrl.service.GetVoucherWithMinRatePoint(ratePoint)
	if err != nil {
		ctrl.logger.Error("Failed to get vouchers", zap.Error(err))
		utils.ResponseError(c, "GET_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	if len(vouchers) == 0 {
		ctrl.logger.Error("No vouchers found")
		utils.ResponseError(c, "GET_VOUCHER_ERROR", "No vouchers found", http.StatusNotFound)
		return
	}

	ctrl.logger.Info("Get vouchers successfully")
	utils.ResponseOK(c, vouchers, "Get vouchers successfully")
}
