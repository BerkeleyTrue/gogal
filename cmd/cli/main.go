package main

import (
	"log"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app"
)

// TODO: add cli commands

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
