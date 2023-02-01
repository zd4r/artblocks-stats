package main

import (
	"log"

	"github.com/zd4r/artblocks-stats/cmd/api/config"
	"github.com/zd4r/artblocks-stats/internal/api/app"
)

func main() {
	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// Run app
	app.Run(cfg)
}
