package intensity

import (
	"image"
	"image/color"
	"math"
)

func GammaCorrection(img image.Image, gamma float64, c float64) image.Image {
	bounds := img.Bounds()
	gammaImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)

			r, g, b, a := originalColor.RGBA()

			normalizedR := float64(r>>8) / 255.0
			normalizedG := float64(g>>8) / 255.0
			normalizedB := float64(b>>8) / 255.0

			gammaR := c * math.Pow(normalizedR, gamma)
			gammaG := c * math.Pow(normalizedG, gamma)
			gammaB := c * math.Pow(normalizedB, gamma)

			finalR := uint8(math.Min(gammaR*255.0, 255.0))
			finalG := uint8(math.Min(gammaG*255.0, 255.0))
			finalB := uint8(math.Min(gammaB*255.0, 255.0))

			gammaColor := color.RGBA{R: finalR, G: finalG, B: finalB, A: uint8(a >> 8)}
			gammaImg.Set(x, y, gammaColor)
		}
	}

	return gammaImg
}
