package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/fx"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/controllers"
)

var Module = fx.Options(
  fx.Provide(NewApp),
  fx.Provide(controllers.NewService),
  fx.Invoke(controllers.Register),
)

func NewApp(cfg *config.Config) *fiber.App {
  isDev := cfg.Release == "development"
  directory := cfg.Directory

  engine := html.New("./web/views", ".html")
  engine.Reload(isDev)


	app := fiber.New(fiber.Config{
    Views: engine,
  })

  app.Static("/images", directory)
  app.Static("/", "./web/public")

	return app;
}
