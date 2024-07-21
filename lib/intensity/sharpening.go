package intensity

import (
	"image"
	"image/color"

	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/utils"
)

func Sharpening(img image.Image, kernel [][]float64) image.Image {
	lapImage := convolution.Convolve(img, kernel)

	bounds := img.Bounds()
	sharpenedImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			lapColor := lapImage.At(x, y)

			origR, origG, origB, origA := originalColor.RGBA()
			lapR, lapG, lapB, lapA := lapColor.RGBA()

			newR := utils.Clamp(float64(origR>>8)-float64(lapR>>8), 0, 255)
			newG := utils.Clamp(float64(origG>>8)-float64(lapG>>8), 0, 255)
			newB := utils.Clamp(float64(origB>>8)-float64(lapB>>8), 0, 255)
			_ = utils.Clamp(float64(origA>>8)-float64(lapA>>8), 0, 255)

			sharpenedImg.Set(x, y, color.RGBA{
				R: uint8(newR),
				G: uint8(newG),
				B: uint8(newB),
				A: 255,
			})
		}
	}

	return sharpenedImg
}
