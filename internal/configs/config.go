package configs

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	TelegramToken string `toml:"telegram_token"`
}

func (cfg *Config) decodeFromToml() error {
	_, err := toml.DecodeFile("internal/configs/config.toml", cfg)
	if err != nil {
		return err
	}

	return nil
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := cfg.decodeFromToml()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
