package config

var Config = &ServiceConfig{}

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type ServiceConfig struct {
	ServiceName      string         `mapstructure:"service_name"`
	HttpServerListen string         `mapstructure:"http_server_listen"`
	Environment      string         `mapstructure:"environment"`
	MongoDBConfig    *MongoDBConfig `mapstructure:"mongo_db_config"`
}

type MongoDBConfig struct {
	URL      string `mapstructure:"url"`
	Database string `mapstructure:"database"`
}
