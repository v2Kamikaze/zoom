package effect

import (
	"image"
	"image/color"

	"github.com/v2Kamikaze/zoom/lib/convert"
)

func Negative(img image.Image) image.Image {
	bounds := img.Bounds()
	negImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)

			r, g, b, a := originalColor.RGBA()

			normalizedR := uint8(r >> 8)
			normalizedG := uint8(g >> 8)
			normalizedB := uint8(b >> 8)
			normalizedA := uint8(a >> 8)

			negR, negG, negB := convert.NegativeRGB(normalizedR, normalizedG, normalizedB)

			negColor := color.RGBA{R: negR, G: negG, B: negB, A: normalizedA}
			negImg.Set(x, y, negColor)
		}
	}

	return negImg
}
