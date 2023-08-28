package controllers

import (
	"github.com/gofiber/fiber/v2"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/utils"
)

type Service struct {
	directory string
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		directory: cfg.Directory,
	}
}

func Register(app *fiber.App, s *Service) {
	app.Get("/", s.Index)
	app.Get("/pics/:dir", s.Pics)
}

func (s *Service) Index(c *fiber.Ctx) error {
	dirs := utils.GetDirectores(s.directory)
	return c.Render("index", fiber.Map{
		"Title":     "Hello, World!",
		"Directory": s.directory,
		"Dirs":      dirs,
	}, "layouts/main")
}
