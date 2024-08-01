package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/v2Kamikaze/zoom/lib/handler"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

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
	effects.Post("/fourier", handler.FourierImage)

	transform := app.Group("/api/transform")
	transform.Post("/scale/bilinear", handler.ScaleImageBilinear)
	transform.Post("/scale/nearest-neighbor", handler.ScaleImageNearestNeighbor)
	transform.Post("/rotate/bilinear", handler.RotateImageBilinear)
	transform.Post("/rotate/nearest-neighbor", handler.RotateImageNearestNeighbor)

	histogram := app.Group("/api/histogram")
	histogram.Post("/rgbl", handler.HistogramRGBL)
	histogram.Post("/equalize", handler.HistogramEqualize)

	app.Listen(":8080")
}
