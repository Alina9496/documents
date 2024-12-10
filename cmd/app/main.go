package main

import (
	"log"

	"github.com/Alina9496/documents/config"
	"github.com/Alina9496/documents/internal/app"
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