package controllers

import (
	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/config"
)

type Controller struct {
	directory string
}

func NewController(cfg *config.Config) *Controller {
	return &Controller{
		directory: cfg.Directory,
	}
}

func RegisterRoutes(app *fiber.App, c *Controller) {
	app.Get("/", c.Index)
	app.Get("/pics/*", c.Pics)
}
