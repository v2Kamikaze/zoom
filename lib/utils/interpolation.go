package utils

import (
	"image"
	"image/color"
	"math"
)

func BilinearInterpolation(img image.Image, x, y float64) color.Color {
	x1, y1 := int(math.Floor(x)), int(math.Floor(y))
	x2, y2 := int(math.Ceil(x)), int(math.Ceil(y))

	c11 := img.At(x1, y1)
	c12 := img.At(x1, y2)
	c21 := img.At(x2, y1)
	c22 := img.At(x2, y2)

	r11, g11, b11, a11 := c11.RGBA()
	r12, g12, b12, a12 := c12.RGBA()
	r21, g21, b21, a21 := c21.RGBA()
	r22, g22, b22, a22 := c22.RGBA()

	fr11, fg11, fb11, fa11 := float64(r11), float64(g11), float64(b11), float64(a11)
	fr12, fg12, fb12, fa12 := float64(r12), float64(g12), float64(b12), float64(a12)
	fr21, fg21, fb21, fa21 := float64(r21), float64(g21), float64(b21), float64(a21)
	fr22, fg22, fb22, fa22 := float64(r22), float64(g22), float64(b22), float64(a22)

	dx, dy := x-float64(x1), y-float64(y1)

	fr := fr11*(1-dx)*(1-dy) + fr21*dx*(1-dy) + fr12*(1-dx)*dy + fr22*dx*dy
	fg := fg11*(1-dx)*(1-dy) + fg21*dx*(1-dy) + fg12*(1-dx)*dy + fg22*dx*dy
	fb := fb11*(1-dx)*(1-dy) + fb21*dx*(1-dy) + fb12*(1-dx)*dy + fb22*dx*dy
	fa := fa11*(1-dx)*(1-dy) + fa21*dx*(1-dy) + fa12*(1-dx)*dy + fa22*dx*dy

	return color.RGBA64{
		R: uint16(fr),
		G: uint16(fg),
		B: uint16(fb),
		A: uint16(fa),
	}
}
