package depend

import (
	"PROJECT_NAME/config"
	"PROJECT_NAME/pkg/bootstrap"
	"context"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

type Config struct {
	Path string
}

var _ bootstrap.Component = (*Config)(nil)

func (d *Config) Init(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())
	viper.AddConfigPath(d.Path)
	viper.SetConfigType("yaml")
	err := viper.Unmarshal(config.Config)
	if err != nil {
		return err
	}
	return nil
}
