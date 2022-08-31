package component

import (
	"PROJECT_NAME/config"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/johnpoint/go-bootstrap/core"
	"github.com/spf13/viper"
)

type Config struct {
	Path string
}

var _ core.Component = (*Config)(nil)

func (d *Config) Init(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(d.Path)
	viper.SetConfigFile(d.Path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(config.Config)
	if err != nil {
		panic(err)
	}
	return nil
}
