package controllers

import (
	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/internal/utils"
)

func (c *Controller) Pics(ctx *fiber.Ctx) error {
	uri := ctx.Params("*")
	breadcrumbs := buildBreadcrumbs(uri)

	path := c.directory + "/" + uri
	isDir := c.imageService.IsPathADir(path)
	dirs := utils.GetDirectories(path, c.directory)

	if !isDir {
		thisImageIndex, numOfPics, nextDir, prevDir := c.imageService.GetImages(dirs, uri)

		return ctx.Render("index", fiber.Map{
			"Title":       uri,
			"BreadCrumbs": breadcrumbs,
			"IsDir":       false,
			"Uri":         "/images/" + uri,
			"NumOfPics":   numOfPics,
			"Index":       thisImageIndex + 1,
			"Next":        nextDir.Uri,
			"Prev":        prevDir.Uri,
			"NextImage":   nextDir.Image,
			"PrevImage":   prevDir.Image,
		}, "layouts/main")
	}

	return ctx.Render("index", fiber.Map{
		"Title":       uri,
		"Dirs":        dirs,
		"BreadCrumbs": breadcrumbs,
		"IsDir":       true,
	}, "layouts/main")
}
