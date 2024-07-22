package intensity

import (
	"image"
	"image/color"

	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/utils"
)

func Sharpening(img image.Image, kernel [][]float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	convolvedImg := convolution.Convolve(img, kernel)

	sharpened := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			convolvedColor := color.RGBAModel.Convert(convolvedImg.At(x, y)).(color.RGBA)

			r := utils.Clamp(int(originalColor.R)+int(convolvedColor.R), 0, 255)
			g := utils.Clamp(int(originalColor.G)+int(convolvedColor.G), 0, 255)
			b := utils.Clamp(int(originalColor.B)+int(convolvedColor.B), 0, 255)
			a := originalColor.A

			sharpened.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: a})
		}
	}

	return sharpened
}
