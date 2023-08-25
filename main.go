package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Println("Hello world from port " + port + "!")
  app := fiber.New()
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World 👋!")
  })

  app.Listen(":" + port)
}
