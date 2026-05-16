package config

import (
	logger "log"

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

func MustLoad() *Config {
	_ = godotenv.Load()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		logger.Fatalf("failed to read config: %s", err)
	}

	return &cfg
}
