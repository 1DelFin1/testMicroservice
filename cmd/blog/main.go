package main

import (
	"fmt"

	"github.com/1DelFin1/testMicroservice/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
