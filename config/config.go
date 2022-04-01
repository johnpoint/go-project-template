package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

var Config = &ServiceConfig{}

type ServiceConfig struct {
	configPath       string
	ServiceName      string         `yaml:"service_name"`
	HttpServerListen string         `yaml:"http_server_listen"`
	Environment      string         `yaml:"environment"`
	MongoDBConfig    *MongoDBConfig `yaml:"mongo_db_config"`
}

func (c *ServiceConfig) SetPath(path string) *ServiceConfig {
	c.configPath = path
	return c
}

func (c *ServiceConfig) ReadConfig() error {
	f, err := os.Open(c.configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	cfgByte, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(cfgByte, c); err != nil {
		return err
	}

	return nil
}

type MongoDBConfig struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
}
