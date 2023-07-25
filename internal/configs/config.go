package configs

import (
	"github.com/Nikkoz/mp.gateway/internal/configs/types/environment"
	"github.com/Nikkoz/mp.gateway/internal/configs/types/logger"
	"github.com/caarlos0/env/v7"
)

type (
	Config struct {
		App  App  `envPrefix:"APP_"`
		Auth Auth `envPrefix:"AUTH_"`
		Log  Log  `envPrefix:"LOG_"`
		Db   Db   `envPrefix:"DB_"`
		Http Http `envPrefix:"HTTP_"`
		Grpc Grpc `envPrefix:"GRPC_"`
	}

	App struct {
		Name        string                  `env:"NAME,required"`
		Version     string                  `env:"VERSION,required"`
		Environment environment.Environment `env:"ENV" envDefault:"local"`
	}

	Auth struct {
		Token string `env:"TOKEN,required"`
	}

	Log struct {
		Level logger.LogLevel `env:"LEVEL" envDefault:"debug"`
	}

	Db struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Port     uint16 `env:"PORT" envDefault:"5432"`
		Name     string `env:"NAME,required"`
		User     string `env:"USER,required"`
		Password string `env:"PASSWORD,required"`
		SslMode  bool   `env:"USE_SSL" envDefault:"false"`
	}

	Http struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port uint16 `env:"PORT" envDefault:"8080"`
	}

	Grpc struct {
		Host string `env:"HOST,required"`
		Port uint16 `env:"PORT,required"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
