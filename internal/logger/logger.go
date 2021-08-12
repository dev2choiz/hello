package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var inst *zap.Logger
var sugarInst *zap.SugaredLogger

func init() {
	var config zap.Config
	if "dev" == os.Getenv("APP_ENV") {
		config = zap.NewDevelopmentConfig()
		config.DisableStacktrace = false
	} else {
		config = zap.NewProductionConfig()
	}
	config.Level.SetLevel(zap.InfoLevel)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	inst, _ = config.Build()
	defer inst.Sync() // flushes buffer, if any
	inst = inst.With(zap.Namespace("hello-api"))
	sugarInst = inst.Sugar()
}

func Info(msg string, fields ...zap.Field) {
	inst.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	inst.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	inst.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	inst.Fatal(msg, fields...)
}
