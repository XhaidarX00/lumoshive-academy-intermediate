package log

import (
	"project-voucher-team3/config"

	"go.uber.org/zap"
)

func InitZapLogger(cfg config.Config) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	if cfg.AppDebug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}
