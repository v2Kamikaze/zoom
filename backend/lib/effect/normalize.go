package effect

import (
	"image"
	"image/color"

	"github.com/v2Kamikaze/zoom/lib/utils"
)

func Normalize(img image.Image) image.Image {
	bounds := img.Bounds()
	normImg := image.NewRGBA(bounds)

	var minPixel, maxPixel float64 = 255, 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			gray := utils.ToGrayRGB(r, g, b)
			if float64(gray) < minPixel {
				minPixel = float64(gray)
			}
			if float64(gray) > maxPixel {
				maxPixel = float64(gray)
			}
		}
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := utils.ToGrayRGB(r, g, b)

			normalized := ((float64(gray) - minPixel) / (maxPixel - minPixel)) * 255
			normImg.Set(x, y, color.RGBA{uint8(normalized), uint8(normalized), uint8(normalized), 255})
		}
	}

	return normImg
}
