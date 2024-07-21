package transform

import (
	"image"
	"math"

	"github.com/v2Kamikaze/zoom/lib/utils"
)

func ScaleWithNearestNeighbor(img image.Image, scaleX, scaleY float64) image.Image {
	bounds := img.Bounds()
	newWidth := int(float64(bounds.Dx()) * scaleX)
	newHeight := int(float64(bounds.Dy()) * scaleY)
	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := int(float64(x) / scaleX)
			srcY := int(float64(y) / scaleY)

			newImg.Set(x, y, img.At(srcX, srcY))
		}
	}

	return newImg
}

func ScaleWithBilinear(img image.Image, scaleX, scaleY float64) image.Image {
	bounds := img.Bounds()
	newWidth := int(float64(bounds.Dx()) * scaleX)
	newHeight := int(float64(bounds.Dy()) * scaleY)
	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := float64(x) / scaleX
			srcY := float64(y) / scaleY

			x1 := int(math.Floor(srcX))
			y1 := int(math.Floor(srcY))
			x2 := x1 + 1
			y2 := y1 + 1

			if x2 >= bounds.Dx() {
				x2 = bounds.Dx() - 1
			}
			if y2 >= bounds.Dy() {
				y2 = bounds.Dy() - 1
			}

			q11 := img.At(x1, y1)
			q12 := img.At(x1, y2)
			q21 := img.At(x2, y1)
			q22 := img.At(x2, y2)

			r := utils.BilinearInterpolate(q11, q12, q21, q22, srcX, srcY, float64(x1), float64(y1), float64(x2), float64(y2))
			newImg.Set(x, y, r)
		}
	}

	return newImg
}
