package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initialize() (logger *zap.Logger, err error) {
	if viper.GetString("server.mode") != "local" {
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}

		return logger, nil
	} else {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		return logger, nil
	}
}