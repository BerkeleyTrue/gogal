package controllers

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/internal/utils"
)

func (s *Service) Pics(c *fiber.Ctx) error {
	pics := c.Params("*")
  picsSlice := strings.Split(pics, "/")
  curUri := "/pics"
  bcSize := len(picsSlice) + 1
  breadcrumbs := make([]struct{ Name string; Uri string}, bcSize)
  breadcrumbs[0] = struct{ Name string; Uri string}{Name: "home", Uri: "/"}

  for i, bc := range strings.Split(pics, "/") {
    curUri += "/" + bc
    breadcrumbs[i + 1] = struct{Name string; Uri string}{Name: bc, Uri: curUri}
  }

	dir := s.directory + "/" + pics

  file, err := os.Open(dir)

	if err != nil {
    return err
  }

  fileInfo, err := file.Stat()

  if err != nil {
    return err
  }

  isDir := fileInfo.IsDir()

  if !isDir {
    return c.Render("pics", fiber.Map{
      "Title": pics,
      "BreadCrumbs": breadcrumbs,
      "IsDir": false,
      "Uri": "/images/" + pics,
    }, "layouts/main")
  }

	dirs := utils.GetDirectores(dir + "/", s.directory)

	return c.Render("pics", fiber.Map{
		"Title":       pics,
		"Dirs":        dirs,
		"BreadCrumbs": breadcrumbs,
		"IsDir":       true,
	}, "layouts/main")
}
