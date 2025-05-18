package config

import (
	"go.uber.org/zap"
)

type ctxLoggerKeyType string

const ctxLoggerKey ctxLoggerKeyType = "logger"

var Logger *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialized logger: " + err.Error())
	}

	Logger = logger.Sugar()
}
