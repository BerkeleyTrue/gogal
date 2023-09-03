package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/config"
)

type Controller struct {
	directory string
}

func buildBreadcrumbs(pics string) []struct {
  Name string
  Uri  string
} {
  picsSlice := strings.Split(pics, "/")
  curUri := "/pics"
  bcSize := len(picsSlice) + 1
  breadcrumbs := make([]struct {
    Name string
    Uri  string
  }, bcSize)
  breadcrumbs[0] = struct {
    Name string
    Uri  string
  }{Name: "home", Uri: "/"}

  for i, bc := range strings.Split(pics, "/") {
    curUri += "/" + bc
    breadcrumbs[i+1] = struct {
      Name string
      Uri  string
    }{Name: bc, Uri: curUri}
  }

  return breadcrumbs
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
