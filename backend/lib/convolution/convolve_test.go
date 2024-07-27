package convolution_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/convolution"
)

func createTestImage(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {
			if x%2 == 0 || y%2 == 0 {
				img.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			} else {
				img.Set(x, y, color.RGBA{R: 0, G: 255, B: 0, A: 255})
			}
		}
	}
	return img
}

var kernel = [][]float64{
	{0, 0, 0},
	{0, 1, 0},
	{0, 0, 0},
}

func TestConvolve(t *testing.T) {
	t.Run("convolve", func(t *testing.T) {

		img := createTestImage(10, 10)

		got := convolution.Convolve(img, kernel)

		assert.Equal(t, img, got, "expect %+v, got %+v", img, got)
	})
}
