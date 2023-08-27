package app

import (
	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/fx"
)

var Module = fx.Options(
  fx.Provide(NewApp),
  fx.Provide(controllers.NewService),
  fx.Invoke(controllers.Register),
)

func NewApp(cfg *config.Config) *fiber.App {
  engine := html.New("./web/views", ".html")

	app := fiber.New(fiber.Config{
    Views: engine,
  })

	return app;
}
