package handler

import (
	"image"

	"github.com/gofiber/fiber/v2"
	"github.com/v2Kamikaze/zoom/lib/kernel"
	"github.com/v2Kamikaze/zoom/lib/zoom"
)

// SobelXImage aplica o filtro Sobel no eixo X à imagem recebida e retorna a imagem processada.
func SobelXImage(c *fiber.Ctx) error {
	return ProcessImage(c, zoom.ApplySobelX)
}

// SobelYImage aplica o filtro Sobel no eixo Y à imagem recebida e retorna a imagem processada.
func SobelYImage(c *fiber.Ctx) error {
	return ProcessImage(c, zoom.ApplySobelY)
}

// SobelMagImage aplica o filtro Sobel nos eixos X e Y e calcula a magnitude, retornando a imagem processada.
func SobelMagImage(c *fiber.Ctx) error {
	return ProcessImage(c, zoom.ApplySobel)
}

// GaussianImage aplica o filtro Gaussiano à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	ks: Tamanho do kernel (padrão 3)
//	s: Sigma do kernel (padrão 1.0)
func GaussianImage(c *fiber.Ctx) error {
	kernelSize := c.QueryInt("ks", 3)
	sigma := c.QueryFloat("s", 1.0)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyGaussian(img, uint(kernelSize), sigma)
	})
}

// LaplacianImage aplica o filtro Laplaciano à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	ks: Tamanho do kernel (padrão 3)
func LaplacianImage(c *fiber.Ctx) error {
	kernelSize := c.QueryInt("ks", 3)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyLaplacian(img, uint(kernelSize))
	})
}

// MeanImage aplica o filtro de média à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	ks: Tamanho do kernel (padrão 3)
func MeanImage(c *fiber.Ctx) error {
	kernelSize := c.QueryInt("ks", 3)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyMean(img, uint(kernelSize))
	})
}

// BinImage aplica uma binarização à imagem recebida com o limiar fornecido e retorna a imagem processada.
// Parâmetros:
//
//	t: Limiar de binarização (padrão 128)
func BinImage(c *fiber.Ctx) error {
	threshold := c.QueryInt("t", 128)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyBin(img, uint8(threshold))
	})
}

// GammaImage aplica correção gama à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	g: Valor do gamma (padrão 2.0)
//	c: Constante multiplicativa (padrão 1.0)
func GammaImage(c *fiber.Ctx) error {
	gamma := c.QueryFloat("g", 2.0)
	cor := c.QueryFloat("c", 1.0)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyGamma(img, gamma, cor)
	})
}

// NegativeImage aplica o efeito negativo à imagem recebida e retorna a imagem processada.
func NegativeImage(c *fiber.Ctx) error {
	return ProcessImage(c, zoom.ApplyNeg)
}

// HighBoostImage aplica o filtro High Boost à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	ks: Tamanho do kernel (padrão 3)
//	s: Sigma do kernel Gaussiano (padrão 1.0)
//	k: Fator de boost (padrão 1.5)
func HighBoostImage(c *fiber.Ctx) error {
	kernelSize := c.QueryInt("ks", 3)
	sigma := c.QueryFloat("s", 1.0)
	k := c.QueryFloat("k", 1.5)

	return ProcessImage(c, func(img image.Image) image.Image {
		smoothKernel := kernel.Gaussian(uint(kernelSize), sigma)
		return zoom.ApplyHighBoost(img, k, smoothKernel)
	})
}

// SharpeningImage aplica o filtro de nitidez à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	ks: Tamanho do kernel (padrão 3)
func SharpeningImage(c *fiber.Ctx) error {
	kernelSize := c.QueryInt("ks", 3)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplySharpening(img, uint(kernelSize))
	})
}
