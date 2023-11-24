package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	loc, _ := time.LoadLocation("UTC")
	enc.AppendString(t.In(loc).Format("2006-01-02T15:04:05.000Z0700"))
}

func productionConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "sso.log"}
	return config
}

func developmentConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return config
}

func MustLogger(env string) *zap.Logger {
	var config zap.Config
	switch env {
	case "production":
		config = productionConfig()
	default:
		config = developmentConfig()
	}
	config.EncoderConfig.EncodeTime = customTimeEncoder
	return zap.Must(config.Build())
}
