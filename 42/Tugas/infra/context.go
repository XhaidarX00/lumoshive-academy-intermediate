package infra

import (
	"project-voucher-team3/config"
	"project-voucher-team3/controller"
	"project-voucher-team3/database"
	"project-voucher-team3/log"
	"project-voucher-team3/middleware"
	"project-voucher-team3/repository"
	"project-voucher-team3/service"
	"project-voucher-team3/utils"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl controller.Controller
	Log *zap.Logger
	Mid *middleware.Middleware
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := log.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	config2, err := utils.ReadConfig()
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(config2, 1*60)

	// instance repository
	repository := repository.NewRepository(db, rdb)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := controller.NewController(service, log)

	midd := middleware.NewMiddleware(rdb)

	return &ServiceContext{Cfg: config, Ctl: *Ctl, Log: log, Mid: &midd}, nil
}
