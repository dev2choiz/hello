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
	configureForStackDriver(&config)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	inst, _ = config.Build()
	defer inst.Sync() // flushes buffer, if any
	inst = inst.WithOptions(zap.AddCallerSkip(1))
	inst = inst.With(zap.Namespace("more"))
	sugarInst = inst.Sugar()
}

func configureForStackDriver(conf *zap.Config) {
	conf.EncoderConfig.LevelKey = "severity"
	conf.EncoderConfig.MessageKey = "message"
	conf.EncoderConfig.TimeKey = "time"
	conf.EncoderConfig.CallerKey = "logging.googleapis.com/sourceLocation"
	conf.EncoderConfig.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
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
