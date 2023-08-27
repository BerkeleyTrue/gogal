package infra

import (
	"context"
	"log"
	"os"
	"os/signal"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Options(
	app.Module,
	fx.Invoke(RegisterServer),
)

func StartServer(app *fiber.App, config *config.Config) {
	cfg := config.HTTP
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := StopServer(app); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	port := cfg.Port
	// Run server.
	if err := app.Listen(":" + port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
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
