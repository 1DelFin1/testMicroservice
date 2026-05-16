package main

import (
	"fmt"

	"github.com/1DelFin1/testMicroservice/config"
	"github.com/1DelFin1/testMicroservice/internal/lib/logger"
	"github.com/1DelFin1/testMicroservice/internal/storage/pg"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := logger.SetupLogger(cfg.App.Env)
	log.Info("Logger initialized")

	conn, err := pg.GetPGConn(cfg)
	defer pg.Shutdown(log, conn)
	if err != nil {
		log.Error("unable to connect to database: %w", err)
	}
	log.Info("PG connection initialized")
}
