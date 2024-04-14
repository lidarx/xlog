package encoder

import "go.uber.org/zap/zapcore"

func DefaultFileEncoderConfig() (core zapcore.EncoderConfig) {
	return zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		TimeKey:      "ts",
		EncodeTime:   zapcore.TimeEncoderOfLayout("01-02 15:04:05"),
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
}
