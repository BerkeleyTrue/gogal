package controllers

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/internal/utils"
)

func (c *Controller) Pics(ctx *fiber.Ctx) error {
	isHx := ctx.Get("HX-Request") == "true"
  isRangeRequest := ctx.Get("HX-Trigger") == "range";

	uri := ctx.Params("*")
	breadcrumbs := buildBreadcrumbs(uri)

	path := c.directory + "/" + uri
	isDir := c.imageService.IsPathADir(path)
	dirs := utils.GetDirectories(path, c.directory)

	if isRangeRequest || !isDir {

    if isRangeRequest {
      thisImageIndexQuery, err := strconv.Atoi(ctx.Query("index", "0"))

      if err != nil {
        return err
      }

      thisImage, err := utils.GetAtIndex(dirs, thisImageIndexQuery - 1);

      if err != nil {
        return err
      }

      uri = uri + "/" + thisImage.Name
      path = c.directory + "/" + uri

      ctx.Set("HX-Push-Url", "/pics/" + uri)
    }

		thisImageIndex, numOfPics, nextDir, prevDir := c.imageService.GetImages(
			dirs,
			uri,
		)

    breadcrumbs = buildBreadcrumbs(uri)

		dirname := strings.Replace(
			filepath.Dir(path),
			c.directory,
			"",
			1,
		)

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
			"IsHx":        isHx,
			"Dirname":     "/pics" + dirname,
		}, "layouts/main")
	}

	return ctx.Render("index", fiber.Map{
		"Title":       uri,
		"Dirs":        dirs,
		"BreadCrumbs": breadcrumbs,
		"IsDir":       true,
		"IsHx":        isHx,
	}, "layouts/main")
}
