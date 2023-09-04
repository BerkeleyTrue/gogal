package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Index(ctx *fiber.Ctx) error {
  isHx := ctx.Get("HX-Request") == "true"
	breadcrumbs := buildBreadcrumbs("")
	dirs := c.galleryService.GetDirectories(c.directory)

	return ctx.Render("index", fiber.Map{
		"Title":       "Home",
		"Dirs":        dirs,
		"BreadCrumbs": breadcrumbs,
		"IsDir":       true,
		"IsHx":        isHx,
	}, "layouts/main")
}
