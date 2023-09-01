package controllers

import (
	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/internal/utils"
)

func (c *Controller) Index(ctx *fiber.Ctx) error {
	dirs := utils.GetDirectories(c.directory, c.directory)

	return ctx.Render("index", fiber.Map{
		"Title":     "Home",
		"Dirs":      dirs,
	}, "layouts/main")
}
