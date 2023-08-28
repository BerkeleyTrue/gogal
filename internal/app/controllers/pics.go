package controllers

import (
	"berkeleytrue/gogal/internal/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)


func (s *Service) Pics(c *fiber.Ctx) error {
  pics := c.Params("dir")
  dir := s.directory + "/" + pics + "/"
  dirs := utils.GetDirectores(dir)

  fmt.Println(s.directory)
  fmt.Println(pics);
  fmt.Println(dir)

  return c.Render("index", fiber.Map{
    "Title":     "Hello, World!",
    "Directory": s.directory,
    "Dirs":      dirs,
  }, "layouts/main")
}
