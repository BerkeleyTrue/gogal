package handlers

import (
	"berkeleytrue/gogal/internal/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Public(app *fiber.App) {
	app.Get("/", controllers.Index)
}
