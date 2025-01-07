package main

import (
	"log"

	"yalp/internal/config"
	"yalp/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	app.Run(cfg)

}
