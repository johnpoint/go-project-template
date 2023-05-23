package config

var Config = &ServiceConfig{}

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type ServiceConfig struct {
	ServiceName string `mapstructure:"service_name"`
	Environment string `mapstructure:"environment"`

	CORS            []string `mapstructure:"cors"`
	ApiServerListen string   `mapstructure:"api_server_listen"`
}
