package config

var Config = &ServiceConfig{}

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type ServiceConfig struct {
	ServiceName      string `mapstructure:"service_name"`
	HttpServerListen string `mapstructure:"http_server_listen"`
	Environment      string `mapstructure:"environment"`
}
