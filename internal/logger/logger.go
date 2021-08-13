package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var inst *zap.Logger
var sugarInst *zap.SugaredLogger

var lvlMap = map[zapcore.Level]string{
	zapcore.DebugLevel: "DEBUG",
	zapcore.InfoLevel: "INFO",
	zapcore.WarnLevel: "WARNING",
	zapcore.ErrorLevel: "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel: "ALERT",
	zapcore.FatalLevel: "EMERGENCY",
}

func init() {
	var config zap.Config
	if false && "dev" == os.Getenv("APP_ENV") {
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
	inst = inst.WithOptions(zap.AddCallerSkip(2))
	inst = inst.With(zap.Namespace("more"))
	sugarInst = inst.Sugar()
}

func configureForStackDriver(conf *zap.Config) {
	conf.EncoderConfig.LevelKey = "severity"
	conf.EncoderConfig.MessageKey = "message"
	conf.EncoderConfig.TimeKey = "time"
	conf.EncoderConfig.CallerKey = "logging.googleapis.com/sourceLocation"
	conf.EncoderConfig.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(zapLvlToString(l))
	}
}

// zapLvlToString Turn zap level to string according to stackdriver
func zapLvlToString(l zapcore.Level) string {
	if ret, ok := lvlMap[l]; ok {
		return ret
	}
	panic("unknown zap log level")
}

func Debug(msg string, fields ...zap.Field) {
	zapLog(zap.InfoLevel, msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLog(zap.InfoLevel, msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLog(zap.WarnLevel, msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLog(zap.ErrorLevel, msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zapLog(zap.FatalLevel, msg, fields...)
}

func Debugf(msg string, args ...interface{}) {
	zapLog(zap.DebugLevel, fmt.Sprintf(msg, args...))
}

func Infof(msg string, args ...interface{}) {
	zapLog(zap.InfoLevel, fmt.Sprintf(msg, args...))
}

func Warnf(msg string, args ...interface{}) {
	zapLog(zap.WarnLevel, fmt.Sprintf(msg, args...))
}

func Errorf(msg string, args ...interface{}) {
	zapLog(zap.ErrorLevel, fmt.Sprintf(msg, args...))
}

func Fatalf(msg string, args ...interface{}) {
	zapLog(zap.FatalLevel, fmt.Sprintf(msg, args...))
}

func zapLog(l zapcore.Level, msg string, fields ...zap.Field) {
	/*if ce := inst.Check(l, msg); ce != nil {
		ce.Write(fields...)
		return
	}*/

	log.Printf("[%s] %s", zapLvlToString(l), msg)
	switch l {
	case zap.FatalLevel, zap.PanicLevel, zap.DPanicLevel :
		os.Exit(1)
	}
}
