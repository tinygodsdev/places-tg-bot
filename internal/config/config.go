package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/tinygodsdev/datasdk/pkg/storage/mongostorage"
)

type Config struct {
	Env           string `envconfig:"ENV" default:"dokku"`
	ServerType    string `envconfig:"SERVER_TYPE" required:"true"`
	TelegramToken string `envconfig:"TELEGRAM_TOKEN" required:"true"`

	mongostorage.Config
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
