package utils

import (
	"log"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("getImage", func() {
	root, err := filepath.Abs("../../")

	if err != nil {
		log.Fatal(err)
		return
	}

	Context("when the subdirectory is empty", func() {
		It("should return an empty string", func() {
			image, isFound := getImage(root + "/assets/pics/sub-empty")
			Expect(image).To(Equal(""))
			Expect(isFound).To(Equal(false))
		})
	})

	Context("when the subdirectory has images", func() {
		It("should return first image", func() {
			image, isFound := getImage(root + "/assets")

			Expect(image).ToNot(Equal(""))
			Expect(isFound).To(Equal(true))
			Expect(strings.HasSuffix(image, ".jpg")).To(Equal(true))

			imagePath := strings.Replace(image, root, "", 1)

			Expect(imagePath).To(Equal("/assets/pics/abs/babs/cat1.jpg"))
		})
	})

	Context("when the path is the last directory and is empty", func() {
		It("should return an empty string", func() {
			image, isFound := getImage(root + "/assets/pics/sub-empty/empty")

			Expect(image).To(Equal(""))
			Expect(isFound).To(Equal(false))
		})
	})

	Context("when the path is the last directory and has images", func() {
		It("should return first image", func() {
			image, isFound := getImage(root + "/assets/pics/foo/bar")

			Expect(image).ToNot(Equal(""))
			Expect(isFound).To(Equal(true))
			Expect(strings.HasSuffix(image, ".jpg")).To(Equal(true))

			imagePath := strings.Replace(image, root, "", 1)

			Expect(imagePath).To(Equal("/assets/pics/foo/bar/cat3.jpg"))
		})
	})
})

var _ = Describe("GetDirectories", func() {
	root, err := filepath.Abs("../../assets/pics")

	if err != nil {
		log.Fatal(err)
		return
	}

	Context("when the path is empty and last directory", func() {
		It("should return an empty array", func() {
			directories := GetDirectories(root+"/sub-empty/empty", root)

			Expect(len(directories)).To(Equal(0))
		})
	})

	Context("when the path is empty and not last directory", func() {
		It("should return directorystats with no images", func() {
			directories := GetDirectories(root+"/sub-empty", root)
			Expect(len(directories)).To(Equal(1))
			Expect(directories[0].IsDirectory).To(Equal(true))
			Expect(directories[0].Name).To(Equal("empty"))
			Expect(directories[0].Uri).To(Equal("/pics/sub-empty/empty"))
			Expect(directories[0].Image).To(Equal(""))
		})
	})

	Context("when the sub directory has images", func() {
		It("should return stats", func() {
			directories := GetDirectories(root+"/abs", root)

			Expect(len(directories)).To(Equal(1))

			Expect(directories[0].Uri).To(Equal("/pics/abs/babs"))
			Expect(directories[0].Image).To(Equal("/images/abs/babs/cat1.jpg"))
		})
	})

	Context("when the directory has no sub directories but it has images", func() {
	  It("should return stats", func() {
      directories := GetDirectories(root+"/foo/bar", root)

      Expect(len(directories)).To(Equal(2))

      Expect(directories[0].Uri).To(Equal("/pics/foo/bar/cat3.jpg"))
      Expect(directories[0].Image).To(Equal("/images/foo/bar/cat3.jpg"))

      Expect(directories[1].Uri).To(Equal("/pics/foo/bar/cat4.jpg"))
      Expect(directories[1].Image).To(Equal("/images/foo/bar/cat4.jpg"))
    })
  })
})
