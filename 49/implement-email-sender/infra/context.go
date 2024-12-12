package infra

import (
	"dashboard-ecommerce-team2/config"
	"dashboard-ecommerce-team2/controller"
	"dashboard-ecommerce-team2/database"
	"dashboard-ecommerce-team2/helper"
	"dashboard-ecommerce-team2/middleware"
	"dashboard-ecommerce-team2/repository"
	"dashboard-ecommerce-team2/service"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Cfg        config.Configuration
	DB         *gorm.DB
	Ctl        controller.Controller
	Log        *zap.Logger
	Cacher     database.Cacher
	Middleware middleware.Middleware
	Repo       *repository.Repository
}

var MockTest *ServiceContext

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.ReadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := helper.InitZapLogger()
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.InitDB(config)
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(config, 60*60)

	middleware := middleware.NewMiddleware(log, rdb)

	// instance repository
	repository := repository.NewRepository(db, log)

	// instance service
	service := service.NewService(repository, log)

	// instance controller
	Ctl := controller.NewController(service, log, rdb, config)

	MockTest = &ServiceContext{Cfg: config, DB: db, Ctl: *Ctl, Log: log, Cacher: rdb, Middleware: middleware, Repo: &repository}

	return &ServiceContext{Cfg: config, DB: db, Ctl: *Ctl, Log: log, Cacher: rdb, Middleware: middleware, Repo: &repository}, nil
}
