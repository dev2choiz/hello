package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type LoggerMock struct {}

func NewLoggerTest() LoggerInterface {
	if instance == nil {
		instance = &LoggerMock{}
	}
	return instance
}

func (l *LoggerMock) Debug(msg string, fields ...zap.Field) {
	doLog(zap.InfoLevel, msg, fields...)
}

func (l *LoggerMock) Info(msg string, fields ...zap.Field) {
	doLog(zap.InfoLevel, msg, fields...)
}

func (l *LoggerMock) Warn(msg string, fields ...zap.Field) {
	doLog(zap.WarnLevel, msg, fields...)
}

func (l *LoggerMock) Error(msg string, fields ...zap.Field) {
	doLog(zap.ErrorLevel, msg, fields...)
}

func (l *LoggerMock) Fatal(msg string, fields ...zap.Field) {
	doLog(zap.FatalLevel, msg, fields...)
}

func (l *LoggerMock) Debugf(msg string, args ...interface{}) {
	doLog(zap.DebugLevel, fmt.Sprintf(msg, args...))
}

func (l *LoggerMock) Infof(msg string, args ...interface{}) {
	doLog(zap.InfoLevel, fmt.Sprintf(msg, args...))
}

func (l *LoggerMock) Warnf(msg string, args ...interface{}) {
	doLog(zap.WarnLevel, fmt.Sprintf(msg, args...))
}

func (l *LoggerMock) Errorf(msg string, args ...interface{}) {
	doLog(zap.ErrorLevel, fmt.Sprintf(msg, args...))
}

func (l *LoggerMock) Fatalf(msg string, args ...interface{}) {
	doLog(zap.FatalLevel, fmt.Sprintf(msg, args...))
}
