package controller

import (
	"project-voucher-team3/service"

	"go.uber.org/zap"
)

type Controller struct {
	User    UserController
	Redeem  RedeemController
	Voucher VoucherController
	Auth    AuthController
}

func NewController(service service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		User:    *NewUserController(service.User, logger),
		Redeem:  *NewRedeemController(service.Reedem, service.User, logger),
		Voucher: *NewVoucherController(service.Voucher, logger),
		Auth:    *NewAuthController(service.Auth, logger),
	}
}
