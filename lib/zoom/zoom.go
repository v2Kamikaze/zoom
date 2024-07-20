package zoom

import (
	"image"

	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/filter"
	"github.com/v2Kamikaze/zoom/lib/intensity"
)

func ApplyGaussian(img image.Image, kernelSize uint, sigma float64) image.Image {
	return convolution.Convolve(img, filter.Gaussian(kernelSize, sigma))
}

func ApplyLaplacian(img image.Image, kernelSize uint) image.Image {
	return convolution.Convolve(img, filter.Laplacian(kernelSize))
}

func ApplyMean(img image.Image, kernelSize uint) image.Image {
	return convolution.Convolve(img, filter.Mean(kernelSize))
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
	return filter.HighBoost(img, k, smoothKernel)
}
