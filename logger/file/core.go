package file

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"xlog/encoder"
)

func NewCore(level zapcore.Level, w io.Writer, encoderConfig zapcore.EncoderConfig, extra ...zapcore.Field) (core zapcore.Core) {
	return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(w), zap.NewAtomicLevelAt(level)).With(extra)
}

func NewDefaultCore(level zapcore.Level, w io.Writer, extra ...zapcore.Field) (core zapcore.Core) {
	return NewCore(level, w, encoder.DefaultCliEncoderConfig(), extra...)
}
