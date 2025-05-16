package main

import (
	"log"

	"github.com/nadduli/ShipSwift/config"
	"github.com/nadduli/ShipSwift/internal/api"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded successfully %v\n", err)
	}

	api.StartServer(cfg)
}
