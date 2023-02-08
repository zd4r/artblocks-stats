package main

import (
	"log"

	"github.com/zd4r/artblocks-stats/internal/api/app"
	"github.com/zd4r/artblocks-stats/pkg/config"
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
