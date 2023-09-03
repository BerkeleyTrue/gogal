package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/controllers"
	"berkeleytrue/gogal/internal/app/services"
)

var Module = fx.Options(
  fx.Provide(services.NewImageService),
  fx.Provide(services.NewGalleryService),
	fx.Provide(controllers.NewController),
	fx.Invoke(SetupStatic),
	fx.Invoke(controllers.RegisterRoutes),
)

func SetupStatic(app *fiber.App, cfg *config.Config) {
	directory := cfg.Directory

	app.Static("/images", directory)
	app.Static("/", "./web/public")
}
