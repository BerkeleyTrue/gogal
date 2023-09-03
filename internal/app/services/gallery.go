package services

import (
	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/domain/models"
	"berkeleytrue/gogal/internal/utils"
)

type GalleryService struct {
	directory string
}

func NewGalleryService(cfg *config.Config) *GalleryService {
	return &GalleryService{
		directory: cfg.Directory,
	}
}

func (g *GalleryService) GetDirectories(dir string) []models.DirectoryStat {
	return utils.GetDirectories(dir, g.directory)
}
