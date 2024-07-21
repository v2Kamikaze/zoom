package transform

import (
	"image"
	"math"

	"github.com/v2Kamikaze/zoom/lib/kernel"
	"github.com/v2Kamikaze/zoom/lib/utils"
)

func RotateWithNearestNeighbor(img image.Image, angle float64) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	centerX := float64(width) / 2
	centerY := float64(height) / 2

	rotationMatrix := kernel.Rotate(angle)
	newImg := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tx := float64(x) - centerX
			ty := float64(y) - centerY

			srcX := rotationMatrix[0][0]*tx + rotationMatrix[0][1]*ty + centerX
			srcY := rotationMatrix[1][0]*tx + rotationMatrix[1][1]*ty + centerY

			if srcX >= 0 && srcX < float64(width) && srcY >= 0 && srcY < float64(height) {
				origX := int(math.Round(srcX))
				origY := int(math.Round(srcY))
				newImg.Set(x, y, img.At(origX, origY))
			}
		}
	}

	return newImg
}

func RotateWithBilinear(img image.Image, angle float64) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	centerX := float64(width) / 2
	centerY := float64(height) / 2

	rotationMatrix := kernel.Rotate(angle)
	newImg := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tx := float64(x) - centerX
			ty := float64(y) - centerY

			srcX := rotationMatrix[0][0]*tx + rotationMatrix[0][1]*ty + centerX
			srcY := rotationMatrix[1][0]*tx + rotationMatrix[1][1]*ty + centerY

			if srcX >= 0 && srcX < float64(width-1) && srcY >= 0 && srcY < float64(height-1) {
				x1 := int(math.Floor(srcX))
				y1 := int(math.Floor(srcY))
				x2 := x1 + 1
				y2 := y1 + 1

				q11 := img.At(x1, y1)
				q12 := img.At(x1, y2)
				q21 := img.At(x2, y1)
				q22 := img.At(x2, y2)

				r := utils.BilinearInterpolate(q11, q12, q21, q22, srcX, srcY, float64(x1), float64(y1), float64(x2), float64(y2))
				newImg.Set(x, y, r)
			}
		}
	}

	return newImg
}
