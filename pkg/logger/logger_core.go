package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

type LoggerInterface interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
}

type Logger struct {
	logger *zap.Logger
	doLog func(l zapcore.Level, msg string, fields ...zap.Field)
}

var instance LoggerInterface

func NewLogger() LoggerInterface {
	if instance == nil {
		l := &Logger{}
		l.init()
		instance = l
	}
	return instance
}

func (l *Logger) init() {
	encoder := newGcpEncoder()
	l.logger = zap.New(
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)

	defer l.logger.Sync() // flushes buffer, if any
	l.logger = l.logger.With(zap.Namespace("more"))
	if os.Getenv("LOGGING_MODE") == "console" {
		l.doLog = l.doConsoleLog
	} else {
		l.doLog = l.doZapLog
	}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.doLog(zap.InfoLevel, msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.doLog(zap.InfoLevel, msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.doLog(zap.WarnLevel, msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.doLog(zap.ErrorLevel, msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.doLog(zap.FatalLevel, msg, fields...)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.doLog(zap.DebugLevel, fmt.Sprintf(msg, args...))
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.doLog(zap.InfoLevel, fmt.Sprintf(msg, args...))
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.doLog(zap.WarnLevel, fmt.Sprintf(msg, args...))
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.doLog(zap.ErrorLevel, fmt.Sprintf(msg, args...))
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.doLog(zap.FatalLevel, fmt.Sprintf(msg, args...))
}

func (l *Logger) doZapLog(lvl zapcore.Level, msg string, fields ...zap.Field) {
	if ce := l.logger.Check(lvl, msg); ce != nil {
		ce.Write(fields...)
		return
	}
}

func (l *Logger) doConsoleLog(lvl zapcore.Level, msg string, fields ...zap.Field) {
	var logger *log.Logger
	if lvl >= zap.WarnLevel {
		logger = strErrLogger
	} else {
		logger = strOutLogger
	}
	_ = logger.Output(2, fmt.Sprintf("[%s] %s", zapLvlToString(lvl), msg))
	switch lvl {
	case zap.FatalLevel, zap.PanicLevel, zap.DPanicLevel:
		os.Exit(1)
	}
}
