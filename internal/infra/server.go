package infra

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app"
)

var Module = fx.Options(
	app.Module,
	fx.Invoke(RegisterServer),
)

func StartServer(app *fiber.App, config *config.Config) error {
	cfg := config.HTTP
	port := cfg.Port
	go app.Listen(":" + port)

  return nil
}

func StopServer(app *fiber.App) error {
	return app.Shutdown()
}

func RegisterServer(lc fx.Lifecycle, app *fiber.App, config *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			StartServer(app, config)
			return nil
		},
		OnStop: func(_ context.Context) error {
			return StopServer(app)
		},
	})
}
