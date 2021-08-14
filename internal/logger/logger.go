package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var inst *zap.Logger
var strOutLogger = log.New(os.Stdout, "", log.LstdFlags)
var strErrLogger = log.New(os.Stderr, "", log.LstdFlags)

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
	doLog(zap.InfoLevel, msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	doLog(zap.InfoLevel, msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	doLog(zap.WarnLevel, msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	doLog(zap.ErrorLevel, msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	doLog(zap.FatalLevel, msg, fields...)
}

func Debugf(msg string, args ...interface{}) {
	doLog(zap.DebugLevel, fmt.Sprintf(msg, args...))
}

func Infof(msg string, args ...interface{}) {
	doLog(zap.InfoLevel, fmt.Sprintf(msg, args...))
}

func Warnf(msg string, args ...interface{}) {
	doLog(zap.WarnLevel, fmt.Sprintf(msg, args...))
}

func Errorf(msg string, args ...interface{}) {
	doLog(zap.ErrorLevel, fmt.Sprintf(msg, args...))
}

func Fatalf(msg string, args ...interface{}) {
	doLog(zap.FatalLevel, fmt.Sprintf(msg, args...))
}

func doLog(l zapcore.Level, msg string, fields ...zap.Field) {
	if ce := inst.Check(l, msg); ce != nil {
		ce.Write(fields...)
		return
	}

	var logger *log.Logger
	if l >= zap.WarnLevel {
		logger = strErrLogger
	} else {
		logger = strOutLogger
	}
	_ = logger.Output(2, fmt.Sprintf("[%s] %s", zapLvlToString(l), msg))
	switch l {
	case zap.FatalLevel, zap.PanicLevel, zap.DPanicLevel :
		os.Exit(1)
	}
}
