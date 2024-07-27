package utils

import (
	"image"
	"image/draw"
)

func NextPower2(n int) int {
	if n <= 0 {
		return 1
	}
	if (n & (n - 1)) == 0 {
		return n
	}
	power := 1
	for power < n {
		power <<= 1
	}
	return power
}

func PadImagePow2(img image.Image) image.Image {
	bounds := img.Bounds()
	width := NextPower2(bounds.Dx())
	height := NextPower2(bounds.Dy())

	// Cria uma nova imagem com dimensões ajustadas para potência de 2
	paddedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Preenche a imagem com a cor de fundo (preto)
	draw.Draw(paddedImg, paddedImg.Bounds(), image.Black, image.Point{}, draw.Src)

	// Calcula o offset para centralizar a imagem original
	offsetX := (width - bounds.Dx()) / 2
	offsetY := (height - bounds.Dy()) / 2

	// Desenha a imagem original na nova imagem com padding
	draw.Draw(paddedImg, image.Rect(offsetX, offsetY, offsetX+bounds.Dx(), offsetY+bounds.Dy()), img, bounds.Min, draw.Over)

	return paddedImg
}
