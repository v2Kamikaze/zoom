package convolution

import (
	"image"
	"image/color"

	"github.com/v2Kamikaze/zoom/lib/utils"
)

func Convolve(img image.Image, kernel [][]float64) image.Image {
	bounds := img.Bounds()
	convImg := image.NewRGBA(bounds)

	kernelHeight := len(kernel)
	kernelWidth := len(kernel[0])
	kernelOffsetX := kernelWidth / 2
	kernelOffsetY := kernelHeight / 2

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var rSum, gSum, bSum float64
			var alphaSum float64
			var kernelSum float64

			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					ix := x + kx - kernelOffsetX
					iy := y + ky - kernelOffsetY

					if ix >= bounds.Min.X && ix < bounds.Max.X && iy >= bounds.Min.Y && iy < bounds.Max.Y {
						r, g, b, a := img.At(ix, iy).RGBA()
						weight := kernel[ky][kx]

						rSum += float64(r>>8) * weight
						gSum += float64(g>>8) * weight
						bSum += float64(b>>8) * weight
						alphaSum += float64(a>>8) * weight
						kernelSum += weight
					}
				}
			}

			r := uint8(utils.Clamp(rSum, 0, 255))
			g := uint8(utils.Clamp(gSum, 0, 255))
			b := uint8(utils.Clamp(bSum, 0, 255))
			var a uint8
			if kernelSum != 0 {
				a = uint8(utils.Clamp(alphaSum/kernelSum, 0, 255))
			} else {
				a = 255
			}

			convImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	return convImg
}
