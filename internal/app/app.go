package app

import (
	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/handlers"
	"berkeleytrue/gogal/internal/infra"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
)

func Run(cfg *config.Config) {
  engine := pug.New("./web/views", ".pug")

	app := fiber.New(fiber.Config{
    Views: engine,
  })

	handlers.Public(app)

	infra.StartServerWithGracefulShutdown(app, &cfg.HTTP)
}
