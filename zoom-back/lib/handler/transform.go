package handler

import (
	"image"

	"github.com/gofiber/fiber/v2"
	"github.com/v2Kamikaze/zoom/lib/zoom"
)

// ScaleImageBilinear aplica o escalamento bilinear à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	x: Fator de escala no eixo X (padrão 1)
//	y: Fator de escala no eixo Y (padrão 1)
func ScaleImageBilinear(c *fiber.Ctx) error {
	scaleX := c.QueryFloat("x", 1)
	scaleY := c.QueryFloat("y", 1)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyScaleWithBilinear(img, scaleX, scaleY)
	})
}

// ScaleImageNearestNeighbor aplica o escalamento com vizinho mais próximo à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	x: Fator de escala no eixo X (padrão 1)
//	y: Fator de escala no eixo Y (padrão 1)
func ScaleImageNearestNeighbor(c *fiber.Ctx) error {
	scaleX := c.QueryFloat("x", 1)
	scaleY := c.QueryFloat("y", 1)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyScaleWithNearestNeighbor(img, scaleX, scaleY)
	})
}

// RotateImageBilinear aplica a rotação bilinear à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	a: Ângulo de rotação em graus (padrão 0)
func RotateImageBilinear(c *fiber.Ctx) error {
	angle := c.QueryFloat("a", 0)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyRotateWithBilinear(img, angle)
	})
}

// RotateImageNearestNeighbor aplica a rotação com vizinho mais próximo à imagem recebida e retorna a imagem processada.
// Parâmetros:
//
//	a: Ângulo de rotação em graus (padrão 0)
func RotateImageNearestNeighbor(c *fiber.Ctx) error {
	angle := c.QueryFloat("a", 0)

	return ProcessImage(c, func(img image.Image) image.Image {
		return zoom.ApplyRotateWithNearestNeighbor(img, angle)
	})
}
