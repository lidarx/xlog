package encoder

import (
	"go.uber.org/zap/zapcore"
	"time"
)

func DefaultRedisEncoderConfig() (core zapcore.EncoderConfig) {
	return zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "ts",
		NameKey:      "name",
		EncodeTime:   zapcore.TimeEncoderOfLayout("01/02 15:04:05"),
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
}
