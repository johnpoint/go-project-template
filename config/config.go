package config

import (
	"encoding/json"
	"io"
	"os"
)

var Config = &ServiceConfig{}

type ServiceConfig struct {
	configPath       string
	HttpServerListen string         `json:"http_server_listen"`
	Environment      string         `json:"environment"`
	MongoDBConfig    *MongoDBConfig `json:"mongo_db_config"`
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

	if err = json.Unmarshal(cfgByte, c); err != nil {
		return err
	}

	return nil
}

type MongoDBConfig struct {
	URL      string `json:"url"`
	Database string `json:"database"`
}
