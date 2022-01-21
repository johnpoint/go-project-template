package depend

import (
	"PROJECT_NAME/config"
	"PROJECT_NAME/pkg/bootstrap"
	"context"
	"math/rand"
	"time"
)

type Config struct {
	Path string
}

var _ bootstrap.Component = (*Config)(nil)

func (d *Config) Init(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())
	return config.Config.SetPath(d.Path).ReadConfig()
}