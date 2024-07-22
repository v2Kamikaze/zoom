package transform

import (
	"image"
	"image/color"
	"math"

	"github.com/v2Kamikaze/zoom/lib/utils"
)

func RotateWithNearestNeighbor(img image.Image, angle float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	radians := angle * math.Pi / 180.0

	rotated := image.NewRGBA(bounds)

	cx, cy := float64(width)/2, float64(height)/2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nx := float64(x) - cx
			ny := float64(y) - cy

			originalX := nx*math.Cos(-radians) - ny*math.Sin(-radians) + cx
			originalY := nx*math.Sin(-radians) + ny*math.Cos(-radians) + cy

			nearestX := int(math.Round(originalX))
			nearestY := int(math.Round(originalY))

			if nearestX >= 0 && nearestX < width && nearestY >= 0 && nearestY < height {
				c := img.At(nearestX, nearestY)
				rotated.Set(x, y, c)
			} else {
				rotated.Set(x, y, color.Transparent)
			}
		}
	}

	return rotated
}

func RotateWithBilinear(img image.Image, angle float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	radians := angle * math.Pi / 180.0

	rotated := image.NewRGBA(bounds)

	cx, cy := float64(width)/2, float64(height)/2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nx := float64(x) - cx
			ny := float64(y) - cy

			originalX := nx*math.Cos(-radians) - ny*math.Sin(-radians) + cx
			originalY := nx*math.Sin(-radians) + ny*math.Cos(-radians) + cy

			if originalX >= 0 && originalX < float64(width) && originalY >= 0 && originalY < float64(height) {
				c := utils.BilinearInterpolation(img, originalX, originalY)
				rotated.Set(x, y, c)
			} else {
				rotated.Set(x, y, color.Transparent)
			}
		}
	}

	return rotated
}
