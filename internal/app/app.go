package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/fx"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/controllers"
	"berkeleytrue/gogal/internal/app/services"
	"berkeleytrue/gogal/web/public"
)

var Module = fx.Options(
  fx.Provide(services.NewImageService),
  fx.Provide(services.NewGalleryService),
	fx.Provide(controllers.NewController),
	fx.Invoke(controllers.RegisterRoutes),
	fx.Invoke(SetupStatic),
)

func SetupStatic(app *fiber.App, cfg *config.Config) {
	directory := cfg.Directory

	app.Static("/images", directory)

	app.Use("/", filesystem.New(filesystem.Config{
    Root: http.FS(public.PublicFS),
    Browse: false,
  }))
}
