package cli

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"xlog/encoder"
)

func NewCore(level zapcore.Level, encoderConfig zapcore.EncoderConfig, extra ...zapcore.Field) (core zapcore.Core) {
	return zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.Lock(os.Stdout), zap.NewAtomicLevelAt(level)).With(extra)
}

func NewDefaultCore(level zapcore.Level, extra ...zapcore.Field) (core zapcore.Core) {
	return NewCore(level, encoder.DefaultCliEncoderConfig(), extra...)
}
