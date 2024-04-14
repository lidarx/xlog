package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Apply(cores ...zapcore.Core) {
	zap.ReplaceGlobals(zap.New(zapcore.NewTee(cores...)))
}
