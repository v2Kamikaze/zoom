package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v2Kamikaze/zoom/lib/handler"
)

func main() {
	app := fiber.New()

	effects := app.Group("/api/effects")
	effects.Post("/negative", handler.NegativeImage)
	effects.Post("/sobel-x", handler.SobelXImage)
	effects.Post("/sobel-y", handler.SobelYImage)
	effects.Post("/sobel-mag", handler.SobelMagImage)
	effects.Post("/gaussian", handler.GaussianImage)
	effects.Post("/laplacian", handler.LaplacianImage)
	effects.Post("/mean", handler.MeanImage)
	effects.Post("/bin", handler.BinImage)
	effects.Post("/gamma", handler.GammaImage)
	effects.Post("/high-boost", handler.HighBoostImage)
	effects.Post("/sharpening", handler.SharpeningImage)

	transform := app.Group("/api/transform")
	transform.Post("/scale/bilinear", handler.ScaleImageBilinear)
	transform.Post("/scale/nearest-neighbor", handler.ScaleImageNearestNeighbor)
	transform.Post("/rotate/bilinear", handler.RotateImageBilinear)
	transform.Post("/rotate/nearest-neighbor", handler.RotateImageNearestNeighbor)

	app.Listen(":8080")
}
