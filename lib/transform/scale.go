package transform

import (
	"image"
	"image/color"
	"math"

	"github.com/v2Kamikaze/zoom/lib/utils"
)

func ScaleWithNearestNeighbor(img image.Image, scaleX, scaleY float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	newWidth := int(float64(width) * scaleX)
	newHeight := int(float64(height) * scaleY)

	scaled := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			originalX := int(math.Round(float64(x) / scaleX))
			originalY := int(math.Round(float64(y) / scaleY))

			if originalX >= 0 && originalX < width && originalY >= 0 && originalY < height {
				c := img.At(originalX, originalY)
				scaled.Set(x, y, c)
			} else {
				scaled.Set(x, y, color.Transparent)
			}
		}
	}

	return scaled
}

func ScaleWithBilinear(img image.Image, scaleX, scaleY float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	newWidth := int(float64(width) * scaleX)
	newHeight := int(float64(height) * scaleY)

	scaled := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			originalX := float64(x) / scaleX
			originalY := float64(y) / scaleY

			c := utils.BilinearInterpolation(img, originalX, originalY)
			scaled.Set(x, y, c)
		}
	}

	return scaled
}
