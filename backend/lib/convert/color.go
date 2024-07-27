package convert

import (
	"math"
)

// RGBToHSV converte uma cor RGB para HSV.
func RGBToHSV(r, g, b uint8) (h, s, v float64) {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0

	max := math.Max(math.Max(rf, gf), bf)
	min := math.Min(math.Min(rf, gf), bf)
	delta := max - min

	// Hue calculation
	if delta == 0 {
		h = 0
	} else if max == rf {
		h = 60 * math.Mod((gf-bf)/delta, 6)
	} else if max == gf {
		h = 60 * ((bf-rf)/delta + 2)
	} else {
		h = 60 * ((rf-gf)/delta + 4)
	}
	if h < 0 {
		h += 360
	}

	// Saturation calculation
	if max == 0 {
		s = 0
	} else {
		s = delta / max
	}

	// Value calculation
	v = max

	return
}

// HSVToRGB converte uma cor HSV para RGB.
func HSVToRGB(h, s, v float64) (r, g, b uint8) {
	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
	m := v - c

	var r1, g1, b1 float64

	switch {
	case 0 <= h && h < 60:
		r1, g1, b1 = c, x, 0
	case 60 <= h && h < 120:
		r1, g1, b1 = x, c, 0
	case 120 <= h && h < 180:
		r1, g1, b1 = 0, c, x
	case 180 <= h && h < 240:
		r1, g1, b1 = 0, x, c
	case 240 <= h && h < 300:
		r1, g1, b1 = x, 0, c
	case 300 <= h && h < 360:
		r1, g1, b1 = c, 0, x
	}

	r = uint8((r1 + m) * 255)
	g = uint8((g1 + m) * 255)
	b = uint8((b1 + m) * 255)

	return
}

func ToGrayAverage(r, g, b uint8) (gray uint8) {
	gray = uint8((float64(r) + float64(g) + float64(b)) / 3)
	return
}

// ToGrayWeighted converte uma cor RGB para escala de cinza usando a média ponderada.
func ToGrayWeighted(r, g, b uint8) (gray uint8) {
	gray = uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
	return
}

// ToGrayHSV converte uma cor HSV para escala de cinza usando a média ponderada.
func ToGrayHSV(h, s, v float64) (gray float64) {
	// Converte HSV para RGB
	r, g, b := HSVToRGB(h, s, v)

	// Aplica a média ponderada
	gray = 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	return
}

// NegativeRGB calcula o negativo de uma cor RGB.
func NegativeRGB(r, g, b uint8) (nr, ng, nb uint8) {
	nr = 255 - r
	ng = 255 - g
	nb = 255 - b
	return
}

// NegativeHSV calcula o negativo de uma cor HSV.
func NegativeHSV(h, s, v float64) (nh, ns, nv float64) {
	r, g, b := HSVToRGB(h, s, v)
	nr, ng, nb := NegativeRGB(r, g, b)
	nh, ns, nv = RGBToHSV(nr, nb, ng)
	return
}
