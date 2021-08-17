package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
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
	encoder := newGcpEncoder()
	inst = zap.New(
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)

	defer inst.Sync() // flushes buffer, if any
	inst = inst.With(zap.Namespace("more"))
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

type gcpEncoder struct {
	zapcore.Encoder
}

func newGcpEncoder() zapcore.Encoder {
	conf := zap.NewProductionEncoderConfig()
		conf.LevelKey = "severity"
		conf.MessageKey = "message"
		conf.EncodeTime = zapcore.ISO8601TimeEncoder
		conf.TimeKey = "time"
		conf.CallerKey = ""
		conf.EncodeLevel = func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(zapLvlToString(l))
	}

	return gcpEncoder{
		Encoder: zapcore.NewJSONEncoder(conf),
	}
}

func (enc gcpEncoder) Clone() zapcore.Encoder {
	return gcpEncoder{
		Encoder: enc.Encoder.Clone(),
	}
}

func (enc gcpEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	file := ent.Caller.TrimmedPath()
	idx := strings.LastIndexByte(file, ':')
	if idx != -1 {
		file = file[:idx]
	}

	chunk := []byte(fmt.Sprintf(
		",\"logging.googleapis.com/sourceLocation\":{\"file\":\"%s\",\"line\":\"%d\"}",
		file, ent.Caller.Line))

	buf, err := enc.Encoder.EncodeEntry(ent, fields)
	cop := buf.Bytes()
	final := make([]byte, 0, len(cop) + len(chunk))
	final = append(final, cop[:len(cop) - 2]...)
	final = append(final, chunk...)
	final = append(final, cop[len(cop) - 2:]...)
	buf.Reset()
	_, _ = buf.Write(final)

	return buf, err
}
