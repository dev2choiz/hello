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

// lvlMap map zap log level to stack-drive
var lvlMap = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func zapLvlToString(l zapcore.Level) string {
	if ret, ok := lvlMap[l]; ok {
		return ret
	}
	panic("unknown zap log level")
}

func Debug(msg string, fields ...zap.Field) {
	instance.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	instance.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	instance.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	instance.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	instance.Fatal(msg, fields...)
}

func Debugf(msg string, args ...interface{}) {
	instance.Debugf(fmt.Sprintf(msg, args...))
}

func Infof(msg string, args ...interface{}) {
	instance.Infof(fmt.Sprintf(msg, args...))
}

func Warnf(msg string, args ...interface{}) {
	instance.Warnf(fmt.Sprintf(msg, args...))
}

func Errorf(msg string, args ...interface{}) {
	instance.Errorf(fmt.Sprintf(msg, args...))
}

func Fatalf(msg string, args ...interface{}) {
	instance.Fatalf(fmt.Sprintf(msg, args...))
}
