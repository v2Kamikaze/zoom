package imageio

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func SaveImage(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".jpg", ".jpeg":
		if err = jpeg.Encode(outputFile, img, nil); err != nil {
			panic(err)
		}
	case ".png":
		if err = png.Encode(outputFile, img); err != nil {
			panic(err)
		}
	default:
		panic("invalid image format")
	}
}
