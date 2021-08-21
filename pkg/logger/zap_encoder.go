package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"strings"
)

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
