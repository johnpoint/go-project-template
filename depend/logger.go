package depend

import (
	"PROJECT_NAME/config"
	"PROJECT_NAME/pkg/bootstrap"
	"PROJECT_NAME/pkg/log"
	"context"
	"go.uber.org/zap/zapcore"
)

type Logger struct{}

var _ bootstrap.Component = (*Logger)(nil)

func (d *Logger) Init(ctx context.Context) error {
	var options = make([]log.Option, 0)
	switch config.Config.Environment {
	case config.EnvProd:
		options = append(options,
			log.WithJSONEncoding(),
			log.WrapLevelEncoder(zapcore.CapitalLevelEncoder),
		)
	default:
		options = append(options,
			log.WithConsoleEncoding(),
			log.WrapLevelEncoder(zapcore.CapitalColorLevelEncoder),
			log.WrapTimeEncoder(zapcore.ISO8601TimeEncoder),
		)
	}

	log.OverrideLoggerWithOption(map[string]interface{}{
		"service-name": config.Config.ServiceName,
	}, options...)
	return nil
}
