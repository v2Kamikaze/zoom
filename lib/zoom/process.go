package zoom

import (
	"image"

	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/intensity"
	"github.com/v2Kamikaze/zoom/lib/kernel"
)

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
