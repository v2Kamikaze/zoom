package zoom

import (
	"image"

	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/intensity"
	"github.com/v2Kamikaze/zoom/lib/kernel"
	"github.com/v2Kamikaze/zoom/lib/utils"
)

func ApplySobelX(img image.Image) image.Image {
	imgX := convolution.Convolve(img, kernel.SobelX())
	norm := utils.NormalizeImage(imgX)
	return norm
}

func ApplySobelY(img image.Image) image.Image {
	imgY := convolution.Convolve(img, kernel.SobelY())
	norm := utils.NormalizeImage(imgY)
	return norm
}

func ApplySobel(img image.Image) image.Image {
	imgX := convolution.Convolve(img, kernel.SobelX())
	normX := utils.NormalizeImage(imgX)
	imgY := convolution.Convolve(img, kernel.SobelY())
	normY := utils.NormalizeImage(imgY)
	return kernel.Magnitude(normX, normY)
}

func ApplyGaussian(img image.Image, kernelSize uint, sigma float64) image.Image {
	return convolution.Convolve(img, kernel.Gaussian(kernelSize, sigma))
}

func ApplyLaplacian(img image.Image, kernelSize uint) image.Image {
	return convolution.Convolve(img, kernel.Laplacian(kernelSize))
}

func ApplyMean(img image.Image, kernelSize uint) image.Image {
	return convolution.Convolve(img, kernel.Mean(kernelSize))
}

func ApplyBin(img image.Image, threshold uint8) image.Image {
	return intensity.Binarize(img, threshold)
}

func ApplyGamma(img image.Image, gamma float64, c float64) image.Image {
	return intensity.GammaCorrection(img, gamma, c)
}

func ApplyNeg(img image.Image) image.Image {
	return intensity.Negative(img)
}

func ApplyHighBoost(img image.Image, k float64, smoothKernel [][]float64) image.Image {
	return intensity.HighBoost(img, k, smoothKernel)
}
