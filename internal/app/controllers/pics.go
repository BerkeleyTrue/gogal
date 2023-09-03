package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/internal/utils"
)

func (c *Controller) Pics(ctx *fiber.Ctx) error {
	uri := ctx.Params("*")
	breadcrumbs := buildBreadcrumbs(uri)

	dir := c.directory + "/" + uri

	file, err := os.Open(dir)

	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return err
	}

	isDir := fileInfo.IsDir()
	dirs := utils.GetDirectories(dir, c.directory)

	if !isDir {
		numOfPics := len(dirs)

		thisImageIndex := -1

		for i, dir := range dirs {
			if dir.Image == "/images/"+uri {
				thisImageIndex = i
				break
			}
		}

		nextUri := ""
		// get next item uri
		if len(dirs) > thisImageIndex+1 {
			nextUri = dirs[thisImageIndex+1].Uri
		} else {
			nextUri = dirs[0].Uri
		}

		prevUri := ""
		// get prev item uri
		if thisImageIndex > 0 {
			prevUri = dirs[thisImageIndex-1].Uri
		} else {
			prevUri = dirs[len(dirs)-1].Uri
		}

		return ctx.Render("pics", fiber.Map{
			"Title":       uri,
			"BreadCrumbs": breadcrumbs,
			"IsDir":       false,
			"Uri":         "/images/" + uri,
			"NumOfPics":   numOfPics,
			"Index":       thisImageIndex + 1,
			"Next":        prevUri,
			"Prev":        nextUri,
		}, "layouts/main")
	}

	return ctx.Render("pics", fiber.Map{
		"Title":       uri,
		"Dirs":        dirs,
		"BreadCrumbs": breadcrumbs,
		"IsDir":       true,
	}, "layouts/main")
}
