package imageio

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func OpenImage(filePath string) image.Image {
	var (
		file *os.File
		img  image.Image
		err  error
	)

	if file, err = os.Open(filePath); err != nil {
		panic(err)
	}

	defer file.Close()

	extension := filepath.Ext(file.Name())

	switch extension {
	case ".png":
		if img, err = png.Decode(file); err != nil {
			panic(err)
		}
	case ".jpg", ".jpeg":
		if img, err = jpeg.Decode(file); err != nil {
			panic(err)
		}
	default:
		panic("invalid image format")
	}

	return img
}
