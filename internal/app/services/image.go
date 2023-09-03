package services

import (
	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/domain/models"
	"fmt"
	"os"
)

type ImageService struct {
	directory string
}

func NewImageService(cfg *config.Config) *ImageService {
	return &ImageService{}
}

func (i *ImageService) IsPathADir(path string) bool {

	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return false
	}

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		return false
	}

	return fileInfo.IsDir()
}

func (i *ImageService) GetImages(dirs []models.DirectoryStat, pics string) (int, int, models.DirectoryStat, models.DirectoryStat) {
	numOfPics := len(dirs)

	thisImageIndex := -1

	for i, dir := range dirs {
		if dir.Image == "/images/"+pics {
			thisImageIndex = i
			break
		}
	}

	var nextDir models.DirectoryStat

	if len(dirs) > thisImageIndex+1 {
		nextDir = dirs[thisImageIndex+1]
	} else {
		nextDir = dirs[0]
	}

	var prevDir models.DirectoryStat

	if thisImageIndex > 0 {
		prevDir = dirs[thisImageIndex-1]
	} else {
		prevDir = dirs[len(dirs)-1]
	}

	return thisImageIndex, numOfPics, nextDir, prevDir
}
