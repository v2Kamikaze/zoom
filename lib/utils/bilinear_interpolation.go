package utils

import "image/color"

func BilinearInterpolate(q11, q12, q21, q22 color.Color, x, y, x1, y1, x2, y2 float64) color.Color {
	r1, g1, b1, a1 := q11.RGBA()
	r2, g2, b2, a2 := q12.RGBA()
	r3, g3, b3, a3 := q21.RGBA()
	r4, g4, b4, a4 := q22.RGBA()

	xf := x - x1
	yf := y - y1

	r := uint8(((1-xf)*(1-yf)*float64(r1) + (1-xf)*yf*float64(r2) + xf*(1-yf)*float64(r3) + xf*yf*float64(r4)) / 255)
	g := uint8(((1-xf)*(1-yf)*float64(g1) + (1-xf)*yf*float64(g2) + xf*(1-yf)*float64(g3) + xf*yf*float64(g4)) / 255)
	b := uint8(((1-xf)*(1-yf)*float64(b1) + (1-xf)*yf*float64(b2) + xf*(1-yf)*float64(b3) + xf*yf*float64(b4)) / 255)
	_ = uint8(((1-xf)*(1-yf)*float64(a1) + (1-xf)*yf*float64(a2) + xf*(1-yf)*float64(a3) + xf*yf*float64(a4)) / 255)

	return color.RGBA{R: r, G: g, B: b, A: 255}
}
