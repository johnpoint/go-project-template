package config

import "github.com/spf13/viper"

var Config = &ServiceConfig{}

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type ServiceConfig struct {
	ServiceName string `mapstructure:"service_name"`
	Environment string `mapstructure:"environment"`
	LogLevel    string `mapstructure:"log_level"`

	CORS            []string `mapstructure:"cors"`
	ApiServerListen string   `mapstructure:"api_server_listen"`
}

func InitConfig(path string) {
	if path == "" {
		path = "config_local.yaml"
	}
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(Config)
	if err != nil {
		panic(err)
	}
}
