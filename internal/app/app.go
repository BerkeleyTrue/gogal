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
  is_dev := cfg.Release == "development"

  engine := html.New("./web/views", ".html")
  engine.Reload(is_dev)


	app := fiber.New(fiber.Config{
    Views: engine,
  })

	return app;
}
