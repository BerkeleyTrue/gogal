package app

import (
	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/handlers"
	"berkeleytrue/gogal/internal/infra"

	"github.com/gofiber/fiber/v2"
)

func Run(cfg *config.Config) {
	app := fiber.New()

	handlers.Public(app)

	infra.StartServerWithGracefulShutdown(app, &cfg.HTTP)
}
