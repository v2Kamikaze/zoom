package convolution

import (
	"image"
	"image/color"
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

			for ky := 0; ky < kernelHeight; ky++ {
				for kx := 0; kx < kernelWidth; kx++ {
					ix := x + kx - kernelOffsetX
					iy := y + ky - kernelOffsetY

					if ix >= bounds.Min.X && ix < bounds.Max.X && iy >= bounds.Min.Y && iy < bounds.Max.Y {
						r, g, b, _ := img.At(ix, iy).RGBA()
						rSum += float64(r>>8) * kernel[ky][kx]
						gSum += float64(g>>8) * kernel[ky][kx]
						bSum += float64(b>>8) * kernel[ky][kx]
					}
				}
			}

			r := uint8(clamp(rSum, 0, 255))
			g := uint8(clamp(gSum, 0, 255))
			b := uint8(clamp(bSum, 0, 255))

			convImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	return convImg
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
