package filter

import (
	"image"
	"image/color"

	"github.com/v2Kamikaze/zoom/lib/convolution"
)

func HighBoost(img image.Image, kernelSize uint, boostFactor float64) image.Image {

	// Gera o kernel Laplaciano
	laplacianKernel := Laplacian(kernelSize)

	// Convolui a imagem com o kernel Laplaciano para obter a imagem de alta frequência
	laplacianImg := convolution.Convolve(img, laplacianKernel)

	// Cria a imagem resultante
	bounds := img.Bounds()
	highBoostImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			origColor := img.At(x, y)
			origR, origG, origB, origA := origColor.RGBA()
			convColor := laplacianImg.At(x, y)
			convR, convG, convB, _ := convColor.RGBA()

			// Calcula a imagem resultante com o fator de aumento
			r := uint8(clamp(float64(origR>>8)+boostFactor*float64(convR>>8), 0, 255))
			g := uint8(clamp(float64(origG>>8)+boostFactor*float64(convG>>8), 0, 255))
			b := uint8(clamp(float64(origB>>8)+boostFactor*float64(convB>>8), 0, 255))

			highBoostImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: uint8(origA >> 8)})
		}
	}

	return highBoostImg
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
