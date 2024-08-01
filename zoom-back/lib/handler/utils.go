package handler

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/gofiber/fiber/v2"
)

var validExt = []string{"jpeg", "jpg", "png"}

type ImageProcessorFunc func(image.Image) image.Image

func ProcessImage(c *fiber.Ctx, f ImageProcessorFunc) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Não foi possível processar a imagem")
	}

	fileReader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Não foi possível abrir o arquivo")
	}
	defer fileReader.Close()

	img, ext, err := image.Decode(fileReader)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao decodificar a imagem")
	}

	if !validImageExt(ext) {
		return c.Status(fiber.StatusBadRequest).SendString("Por enquanto, apenas formatos jpeg e png são aceitos")
	}

	procImage := f(img)

	var buffer bytes.Buffer

	if err = encodeImage(&buffer, procImage, ext); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao codificar a imagem")
	}

	c.Type(ext)
	return c.Send(buffer.Bytes())
}

func encodeImage(buffer io.Writer, img image.Image, ext string) error {
	switch ext {
	case "jpg", "jpeg":
		return jpeg.Encode(buffer, img, nil)
	case "png":
		return png.Encode(buffer, img)
	default:
		return errors.New("formato de imagem inválido")
	}
}

func validImageExt(ext string) bool {

	for _, v := range validExt {
		if v == ext {
			return true
		}
	}

	return false
}
