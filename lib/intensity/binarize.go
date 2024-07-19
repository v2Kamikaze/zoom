package intensity

import (
	"image"
	"image/color"
)

func Binarize(img image.Image, threshold uint8) image.Image {
	bounds := img.Bounds()
	binImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()

			gray := ToGrayRGB(r, g, b)

			var binColor color.RGBA
			if gray > threshold {
				binColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			} else {
				binColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
			}

			binImg.Set(x, y, binColor)
		}
	}

	return binImg
}
