package intensity

import (
	"image"
	"image/color"

	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/utils"
)

func HighBoost(img image.Image, k float64, smoothKernel [][]float64) image.Image {
	bounds := img.Bounds()
	highBoostImg := image.NewRGBA(bounds)

	smoothedImg := convolution.Convolve(img, smoothKernel)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			origR, origG, origB, _ := img.At(x, y).RGBA()
			smoothR, smoothG, smoothB, _ := smoothedImg.At(x, y).RGBA()

			hbR := utils.Clamp(float64(origR>>8)+k*(float64(origR>>8)-float64(smoothR>>8)), 0, 255)
			hbG := utils.Clamp(float64(origG>>8)+k*(float64(origG>>8)-float64(smoothG>>8)), 0, 255)
			hbB := utils.Clamp(float64(origB>>8)+k*(float64(origB>>8)-float64(smoothB>>8)), 0, 255)

			highBoostImg.Set(x, y, color.RGBA{R: uint8(hbR), G: uint8(hbG), B: uint8(hbB), A: 255})
		}
	}

	return highBoostImg
}
