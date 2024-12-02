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

type RedeemController struct {
	service service.RedeemService
	user    service.UserService
	logger  *zap.Logger
}

func NewRedeemController(service service.RedeemService, user service.UserService, logger *zap.Logger) *RedeemController {
	return &RedeemController{service, user, logger}
}

func (ctrl *RedeemController) GetUserRedeemVoucherController(c *gin.Context) {
	userID := 1

	voucherType := c.Param("vourcher-type")
	if voucherType == "" {
		ctrl.logger.Error("voucher type is empty")
		utils.ResponseError(c, "EMPTY_PARAM", "voucher type is empty", http.StatusBadRequest)
		return
	}
	voucherFilter := models.Voucher{
		VoucherType: voucherType,
	}
	userRedeem, err := ctrl.service.GetAllUserRedeems(userID, voucherFilter)
	if err != nil {
		ctrl.logger.Error("Failed to get user redeem vouchers", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	if len(userRedeem) == 0 {
		ctrl.logger.Info("User has no redeem vouchers")
		utils.ResponseOK(c, userRedeem, "user has no redeem voucher")
		return
	}
	utils.ResponseOK(c, userRedeem, "user redeem successfully retrieved")
}

func (ctrl *RedeemController) RedeemVoucher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	VoucherId, err := strconv.Atoi(c.Param("voucher_id"))
	if err != nil {
		ctrl.logger.Error("Invalid Voucher Id", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ctrl.user.GetUser(id)
	if err != nil {
		ctrl.logger.Error("User not found", zap.Error(err))
		utils.ResponseError(c, "NOT_FOUND", "User not found", http.StatusNotFound)
		return
	}

	reedem, err := ctrl.service.RedeemVoucher(&user, VoucherId)
	if err != nil {
		ctrl.logger.Error("Reedem voucher error", zap.Error(err))
		utils.ResponseError(c, "REEDEM_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	err = ctrl.user.UpdateUser(user)
	if err != nil {
		ctrl.logger.Error("Error update point user", zap.Error(err))
		utils.ResponseError(c, "ERR0R_UPDATE_POINT_USER", err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Reedem voucher successfully")
	utils.ResponseOK(c, reedem, "Reedem voucher successfully")
}
