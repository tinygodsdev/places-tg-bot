package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerType string `envconfig:"SERVER_TYPE" required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
