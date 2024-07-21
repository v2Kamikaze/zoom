package kernel

import (
	"image"
	"image/color"
	"math"

	"github.com/v2Kamikaze/zoom/lib/utils"
)

func SobelX() [][]float64 {
	return [][]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
}

func SobelY() [][]float64 {
	return [][]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
}

func Magnitude(imgX, imgY image.Image) image.Image {
	bounds := imgX.Bounds()
	magImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rX, gX, bX, _ := imgX.At(x, y).RGBA()
			rY, gY, bY, _ := imgY.At(x, y).RGBA()

			rX8, gX8, bX8 := float64(rX>>8), float64(gX>>8), float64(bX>>8)
			rY8, gY8, bY8 := float64(rY>>8), float64(gY>>8), float64(bY>>8)

			magR := math.Sqrt(rX8*rX8 + rY8*rY8)
			magG := math.Sqrt(gX8*gX8 + gY8*gY8)
			magB := math.Sqrt(bX8*bX8 + bY8*bY8)

			magnitude := math.Sqrt(magR*magR + magG*magG + magB*magB)

			normalizedMag := utils.Clamp(magnitude, 0, 255)

			magImg.SetGray(x, y, color.Gray{Y: uint8(normalizedMag)})
		}
	}

	return magImg
}
