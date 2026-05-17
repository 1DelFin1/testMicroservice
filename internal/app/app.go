package app

import (
	"github.com/1DelFin1/testMicroservice/config"
	"github.com/1DelFin1/testMicroservice/internal/lib/logger"
	"github.com/1DelFin1/testMicroservice/internal/lib/sl"
	"github.com/1DelFin1/testMicroservice/internal/storage/pg"
)

func Run(cfg *config.Config) {
	log := logger.SetupLogger(cfg.App.Env)
	log.Info("Logger initialized")

	pool, err := pg.New(cfg)
	if err != nil {
		log.Error("unable to connect to database", sl.Err(err))
		return
	}
	defer pool.Close(log)
	log.Info("PG pool initialized")
}
