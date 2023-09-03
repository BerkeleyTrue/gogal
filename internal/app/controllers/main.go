package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/app/services"
)

type (
	Controller struct {
		directory      string
		imageService   *services.ImageService
		galleryService *services.GalleryService
	}

	breadCrumb struct {
		Name string
		Uri  string
	}
)

func buildBreadcrumbs(uri string) []breadCrumb {
	uriComp := strings.Split(uri, "/")
	curUri := ""
	bcSize := len(uriComp) + 1

	breadcrumbs := make([]breadCrumb, 0, bcSize)

	breadcrumbs = append(breadcrumbs, breadCrumb{Name: "home", Uri: "/"})

	for _, bc := range uriComp {

		// filter out empty strings
		if bc == "" {
			continue
		}

		curUri += "/" + bc

		breadcrumbs = append(breadcrumbs, breadCrumb{Name: bc, Uri: curUri})
	}

	return breadcrumbs
}

func NewController(cfg *config.Config, imageService *services.ImageService, galleryService *services.GalleryService) *Controller {
	return &Controller{
		directory:      cfg.Directory,
		imageService:   imageService,
		galleryService: galleryService,
	}
}

func RegisterRoutes(app *fiber.App, c *Controller) {
	app.Get("/", c.Index)
	app.Get("/pics/*", c.Pics)
}
