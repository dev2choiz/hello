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

var doLog func(l zapcore.Level, msg string, fields ...zap.Field)

func init() {
	encoder := newGcpEncoder()
	inst = zap.New(
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)

	defer inst.Sync() // flushes buffer, if any
	inst = inst.With(zap.Namespace("more"))
	if os.Getenv("LOGGING_MODE") == "console" {
		doLog = doConsoleLog
	} else {
		doLog = doZapLog
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

func doZapLog(l zapcore.Level, msg string, fields ...zap.Field) {
	if ce := inst.Check(l, msg); ce != nil {
		ce.Write(fields...)
		return
	}
}

func doConsoleLog(l zapcore.Level, msg string, fields ...zap.Field) {
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
