package main

import (
	"fmt"

	"github.com/1DelFin1/testMicroservice/config"
	"github.com/1DelFin1/testMicroservice/internal/lib/logger"
	"github.com/1DelFin1/testMicroservice/internal/lib/sl"
	"github.com/1DelFin1/testMicroservice/internal/storage/pg"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

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
