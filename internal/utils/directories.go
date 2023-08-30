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
		fmt.Println(err)
		return "", false
	}

	for _, dir := range subdirectories {
		// ignore hidden files
		if strings.HasPrefix(dir.Name(), ".") {
			continue
		}

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

func GetDirectories(imagePath, baseDirectory string) []models.DirectoryStat {
	// read imagePath and get all directories
	f, err := os.Open(imagePath)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer f.Close()

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

	var dirStats DirStats

	for _, dirEntry := range dirEntries {
    name := dirEntry.Name()
    if strings.HasPrefix(name, ".") {
      continue
    }

		path := imagePath + "/" + name

		isDir := dirEntry.IsDir()

		image := strings.Replace(path, baseDirectory, "/images", 1)

		if isDir {
			foundImage, isFound := getImage(path)

			if isFound {
				image = strings.Replace(foundImage, baseDirectory, "/images", 1)
			} else {
        image = ""
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
