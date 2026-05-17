package main

import (
	"fmt"

	"github.com/1DelFin1/testMicroservice/config"
	"github.com/1DelFin1/testMicroservice/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg) // test

	app.Run(cfg)
}
