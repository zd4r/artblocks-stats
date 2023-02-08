package config

import (
	"fmt"
	"log"
	"os"

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
		Level      string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
		Structured bool   `yaml:"structured"   env:"STRUCTURED" env-default:"false"`
	}

	PG struct {
		PostgresDB       string `env-required:"true"   env:"POSTGRES_DB"`
		PostgresUser     string `env-required:"true"   env:"POSTGRES_USER"`
		PostgresPassword string `env-required:"true"   env:"POSTGRES_PASSWORD"`
		PostgresHost     string `env-required:"true"   env:"PG_HOST"`
		MaxOpenConns     int    `yaml:"max_open_conns" env:"PG_MAX_OPEN_CONNS" env-default:"25"`
		MaxIdleConns     int    `yaml:"max_idle_conns" env:"PG_MAX_IDLE_CONNS" env-default:"25"`
		MaxIdleTime      string `yaml:"max_idle_time"  env:"PG_MAX_IDLE_TIME"  env-default:"15m"`
	}
)

// New creates new config based on config.yml and env
func New() (*Config, error) {
	cfg := Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = cleanenv.ReadConfig(dir+"/config.yml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &cfg, nil
}
