package infra

import (
	"project_auth_jwt/config"
	"project_auth_jwt/controllers"
	"project_auth_jwt/database"
	"project_auth_jwt/helper"
	"project_auth_jwt/middlewares"
	"project_auth_jwt/repositories"
	"project_auth_jwt/services"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Cfg        config.Configuration
	DB         *gorm.DB
	Ctl        controllers.Controller
	Log        *zap.Logger
	Cacher     database.Cacher
	Middleware middlewares.Middleware
	Repo       *repositories.Repository
}

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

	// instance repository
	repository := repositories.NewRepository(db, log)

	// instance service
	jwtSvc := services.NewJWTService(config.SecretKey)
	service := services.NewService(repository, log, jwtSvc, config.SecretKey)

	middleware := middlewares.NewMiddleware(log, rdb, jwtSvc)

	// instance controller
	Ctl := controllers.NewController(service, log, rdb, config)

	return &ServiceContext{Cfg: config, DB: db, Ctl: *Ctl, Log: log, Cacher: rdb, Middleware: middleware, Repo: &repository}, nil
}
