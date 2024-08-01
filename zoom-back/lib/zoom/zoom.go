package zoom

import (
	"image"

	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/effect"
	"github.com/v2Kamikaze/zoom/lib/kernel"
	"github.com/v2Kamikaze/zoom/lib/transform"
)

func ApplyInvFourier(img image.Image) image.Image {
	return effect.CalculateInverseFourier(img)
}

func ApplyFourier(img image.Image) image.Image {
	return effect.CalculateCenterFourier(img)
}

func ApplyScaleWithBilinear(img image.Image, scaleX float64, scaleY float64) image.Image {
	return transform.ScaleWithBilinear(img, scaleX, scaleY)
}

func ApplyScaleWithNearestNeighbor(img image.Image, scaleX float64, scaleY float64) image.Image {
	return transform.ScaleWithNearestNeighbor(img, scaleX, scaleY)
}

func ApplyRotateWithNearestNeighbor(img image.Image, angle float64) image.Image {
	return transform.RotateWithNearestNeighbor(img, angle)
}

func ApplyRotateWithBilinear(img image.Image, angle float64) image.Image {
	return transform.RotateWithBilinear(img, angle)
}

func ApplySobelX(img image.Image) image.Image {
	imgX := convolution.Convolve(img, kernel.SobelX())
	norm := effect.Normalize(imgX)
	return norm
}

func ApplySobelY(img image.Image) image.Image {
	imgY := convolution.Convolve(img, kernel.SobelY())
	norm := effect.Normalize(imgY)
	return norm
}

func ApplySobel(img image.Image) image.Image {
	imgX := convolution.Convolve(img, kernel.SobelX())
	normX := effect.Normalize(imgX)
	imgY := convolution.Convolve(img, kernel.SobelY())
	normY := effect.Normalize(imgY)
	return effect.Magnitude(normX, normY)
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
	return effect.Binarize(img, threshold)
}

func ApplyGamma(img image.Image, gamma float64, c float64) image.Image {
	return effect.GammaCorrection(img, gamma, c)
}

func ApplyNeg(img image.Image) image.Image {
	return effect.Negative(img)
}

func ApplyHighBoost(img image.Image, k float64, smoothKernel [][]float64) image.Image {
	return effect.HighBoost(img, k, smoothKernel)
}

func ApplySharpening(img image.Image, kernelSize uint) image.Image {
	return effect.Sharpening(img, kernel.Laplacian(kernelSize))
}
