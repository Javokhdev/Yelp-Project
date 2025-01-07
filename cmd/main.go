package main

import (
	"log"

	"github.com/Javokhdev/Yelp-Project/config"
	"github.com/Javokhdev/Yelp-Project/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
