// package log zap.go.
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger LoggerWrapper

type LoggerWrapper = *zap.Logger

func init() {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		
	})
	core := zapcore.NewCore(encoder)
	tee := zapcore.NewTee(core)
	logger = zap.New(tee)
}
