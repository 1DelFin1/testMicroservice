package main

import (
	"fmt"

	"github.com/1DelFin1/testMicroservice/config"
	"github.com/1DelFin1/testMicroservice/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := logger.SetupLogger(cfg.App.Env)
	log.Info("Logger initialized")
}
