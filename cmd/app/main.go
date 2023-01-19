package main

import (
	"log"

	"github.com/zd4rova/artblocks-holders/config"
	"github.com/zd4rova/artblocks-holders/internal/app"
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
