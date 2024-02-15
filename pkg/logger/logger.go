package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(string, ...zapcore.Field)
	Info(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
	Fatal(string, ...zapcore.Field)
}

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func NewSugar() *zap.SugaredLogger {
	return zap.Must(basicConfig().Build()).Sugar()
}

func NewCustomLogger() *zap.Logger {
	return zap.Must(basicConfig().Build())
}
