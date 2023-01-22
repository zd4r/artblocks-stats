package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		URL          string `env-required:"true"   env:"PG_URL"`
		MaxOpenConns int    `yaml:"max_open_conns" env:"PG_MAX_OPEN_CONNS" env-default:"25"`
		MaxIdleConns int    `yaml:"max_idle_conns" env:"PG_MAX_IDLE_CONNS" env-default:"25"`
		MaxIdleTime  string `yaml:"max_idle_time"  env:"PG_MAX_IDLE_TIME"  env-default:"15m"`
	}
)

func NewConfig() (*Config, error) {
	cfg := Config{}

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &cfg, nil
}
