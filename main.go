package main

import (
	"log"

	"berkeleytrue/gogal/internal/app/handlers"
	"berkeleytrue/gogal/internal/infra"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	handlers.Public(app)

	infra.StartServerWithGracefulShutdown(app)
}
