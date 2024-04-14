package encoder

import (
	"go.uber.org/zap/zapcore"
	"time"
)

func DefaultCliEncoderConfig() (core zapcore.EncoderConfig) {
	return zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		TimeKey:      "ts",
		EncodeTime:   zapcore.TimeEncoderOfLayout("15:04:05"),
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
}
