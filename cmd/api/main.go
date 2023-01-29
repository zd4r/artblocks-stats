package main

import (
	"log"

	"github.com/zd4rova/artblocks-stats/cmd/api/config"
	"github.com/zd4rova/artblocks-stats/internal/api/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// Run
	app.Run(cfg)
}
