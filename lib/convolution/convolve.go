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
			var rSum, gSum, bSum, aSum float64

			// Verificar se esse é o melhor caminho para não perder a transparência
			var kernelSum float64

			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					ix := x + kx - kernelOffsetX
					iy := y + ky - kernelOffsetY

					if ix >= bounds.Min.X && ix < bounds.Max.X && iy >= bounds.Min.Y && iy < bounds.Max.Y {
						r, g, b, a := img.At(ix, iy).RGBA()
						rSum += float64(r>>8) * kernel[ky][kx]
						gSum += float64(g>>8) * kernel[ky][kx]
						bSum += float64(b>>8) * kernel[ky][kx]

						// Verificar se esse é o melhor caminho para não perder a transparência
						aSum += float64(a>>8) * kernel[ky][kx]
						kernelSum += kernel[ky][kx]

					}
				}
			}

			// Verificar se esse é o melhor caminho para não perder a transparência
			if kernelSum != 0 {
				aSum /= kernelSum
			}

			r := uint8(utils.Clamp(rSum, 0, 255))
			g := uint8(utils.Clamp(gSum, 0, 255))
			b := uint8(utils.Clamp(bSum, 0, 255))

			// Verificar se esse é o melhor caminho para não perder a transparência
			_ = uint8(utils.Clamp(aSum, 0, 255))

			convImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	return convImg
}
