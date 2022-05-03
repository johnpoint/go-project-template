package bootstrap

import "go.uber.org/zap"

type log interface {
	Info(name string, fields ...zap.Field)
	Error(name string, fields ...zap.Field)
}

var logger log

type DefaultLogger struct {
	l *zap.Logger
}

func NewDefaultLogger() log {
	logger, _ := zap.NewProduction()
	return &DefaultLogger{
		l: logger,
	}
}

var _ log = (*DefaultLogger)(nil)

func (d *DefaultLogger) Info(name string, fields ...zap.Field) {
	d.l.Info(name, fields...)
}

func (d *DefaultLogger) Error(name string, fields ...zap.Field) {
	d.l.Error(name, fields...)
}
