package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var inst *zap.Logger
var sugarInst *zap.SugaredLogger

func init() {
	dev := "dev" == os.Getenv("APP_ENV")
	_ = dev
	if false {
		inst, _ := zap.NewDevelopment()
		defer inst.Sync() // flushes buffer, if any
		sugarInst = inst.Sugar()
	} else {
		config := zap.NewProductionConfig()
		config.Level.SetLevel(zap.InfoLevel)
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		inst, _ := config.Build()
		defer inst.Sync() // flushes buffer, if any
		sugarInst = inst.With(zap.Namespace("hello-api")).Sugar()
	}
}

func Info(args ...interface{}) {
	sugarInst.Info(args...)
}

func RInfo(msg string, fields ...zap.Field) {
	inst.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	sugarInst.Infof(template, args...)
}

func Warn(args ...interface{}) {
	sugarInst.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugarInst.Warnf(template, args...)
}

func Error(args ...interface{}) {
	sugarInst.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugarInst.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	sugarInst.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugarInst.Fatalf(template, args...)
}
