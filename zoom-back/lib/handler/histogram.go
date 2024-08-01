package handler

import (
	"image"

	"github.com/gofiber/fiber/v2"
	"github.com/v2Kamikaze/zoom/lib/zoom"
)

// HistogramRGBL gera os histogramas dos canais R, G, B e luminância da imagem fornecida e retorna como JSON.
// Parâmetros: Nenhum.
// Retorno: JSON contendo os histogramas para cada canal.
func HistogramRGBL(c *fiber.Ctx) error {
	var hist *zoom.Histogram

	ProcessImage(c, func(img image.Image) image.Image {
		hist = zoom.HistogramFromImage(img)
		return img
	})

	return c.JSON(hist)
}

// HistogramEqualize aplica a equalização de histograma em um canal especificado (R, G, B ou luminância) da imagem.
// Parâmetros:
// - ch: Canal de cor a ser equalizado ('r', 'g', 'b', 'l'). Padrão: 'l'.
// Retorno: Imagem com o canal equalizado.
func HistogramEqualize(c *fiber.Ctx) error {
	ch := c.Query("ch", "l")

	return ProcessImage(c, func(img image.Image) image.Image {
		hist := zoom.HistogramFromImage(img)
		var channel zoom.HistogramChannel

		switch ch {
		case "r":
			channel = zoom.RChannel
		case "g":
			channel = zoom.GChannel
		case "b":
			channel = zoom.BChannel
		case "l":
			channel = zoom.LChannel
		}

		return hist.EqualizeWithChannel(img, channel)
	})
}
