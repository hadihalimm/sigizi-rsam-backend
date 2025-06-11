package config

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string{"NUL"}
	cfg.ErrorOutputPaths = []string{"NUL"}
	logger, err := cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	Logger = logger.Sugar()
}

func WithRequestContext(logger *zap.SugaredLogger, c *gin.Context) *zap.SugaredLogger {
	return logger.With(
		"userID", c.GetUint("userID"),
		"ip", c.ClientIP(),
		"path", c.FullPath(),
		"method", c.Request.Method,
	)
}
