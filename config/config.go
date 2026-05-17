package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App app
		PG  pg
		Log log
	}

	app struct {
		Env  string `env:"APP_ENV,required"`
		Name string `env:"APP_NAME"`
	}

	log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	pg struct {
		Host     string `env:"PG_HOST,required"`
		Port     string `env:"PG_PORT,required"`
		Database string `env:"PG_DATABASE,required"`
		User     string `env:"PG_USER,required"`
		Pass     string `env:"PG_PASSWORD,required"`
		URL      string `env:"PG_URL,required"`
	}
)

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("can`t create config: %w", err)
	}

	return &cfg, nil
}
