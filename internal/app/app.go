package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/controllers"
)

var Module = fx.Options(
	fx.Provide(controllers.NewService),
	fx.Invoke(SetupRoutes),
	fx.Invoke(controllers.Register),
)

func SetupRoutes(app *fiber.App, cfg *config.Config) {
	directory := cfg.Directory

	app.Static("/images", directory)
	app.Static("/", "./web/public")
}
