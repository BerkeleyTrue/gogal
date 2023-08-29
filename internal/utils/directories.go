package utils

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"

	"berkeleytrue/gogal/internal/domain/models"
)

type DirStats []models.DirectoryStat

func (d DirStats) ByImage(i, j int) bool {
	return d[i].Image < d[j].Image
}

// depth first search for image to represent directory
func getImage(path string) (string, bool) {
	var (
		firstImage = ""
		firstDir   = ""
	)

	subdirectories, err := os.ReadDir(path)

	if err != nil {
		return "", false
	}

	for _, dir := range subdirectories {
		if dir.IsDir() {

			if firstDir == "" {
				firstDir = path + "/" + dir.Name()
			}

		} else {

			if firstImage == "" {
				firstImage = path + "/" + dir.Name()
			}
		}

		if firstImage != "" && firstDir != "" {
			break
		}
	}

	if firstImage != "" {
		return firstImage, true
	}

	if firstDir == "" {
		return "", false
	}

	return getImage(firstDir)
}

func GetDirectores(imagePath, baseDirectory string) ([]models.DirectoryStat) {
	// read imagePath and get all directories
	f, err := os.Open(imagePath)

	if err != nil {
	  fmt.Println(err)
		return nil
	}

	// if f is a file, return error
	if err != nil {
	  fmt.Println(err)
		return nil
	}

	dirEntries, err := f.ReadDir(0)

	if err != nil {
	  fmt.Println(err)
		return nil
	}

	type dirIntr struct {
		path     string
		dirEntry fs.DirEntry
	}

	var directories []dirIntr

	// map dirEntries to a struct of fileinfo and path
	for _, fileinfo := range dirEntries {
		directories = append(
			directories,
			dirIntr{
				path:     imagePath + fileinfo.Name(),
				dirEntry: fileinfo,
			},
		)
	}

	var dirStats DirStats

	for _, dirIntr := range directories {
		path := dirIntr.path
		dirEntry := dirIntr.dirEntry

		isDir := dirEntry.IsDir()

		var image string = path

		if isDir {
			foundImage, isFound := getImage(path)
			if isFound {
				image =  strings.Replace(foundImage, imagePath, "/images", 1)
			}
		}

    subpath := strings.Replace(path, baseDirectory, "", 1)

		uri := "/pics" + subpath

		dirStats = append(dirStats,
			models.DirectoryStat{
				IsDirectory: isDir,
				Name:        dirEntry.Name(),
				Uri:         uri,
				Image:       image,
			},
		)
	}

	sort.Slice(dirStats, dirStats.ByImage)

	return dirStats
}
