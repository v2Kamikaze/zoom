package filter

import "math"

func Gaussian(size uint, sigma float64) [][]float64 {
	if size%2 == 0 {
		panic("O tamanho do kernel deve ser ímpar")
	}

	kernel := make([][]float64, size)
	for i := range kernel {
		kernel[i] = make([]float64, size)
	}

	sum := 0.0
	mid := int(size / 2)
	for y := -mid; y <= mid; y++ {
		for x := -mid; x <= mid; x++ {
			value := (1 / (2 * math.Pi * sigma * sigma)) * math.Exp(-(float64(x*x+y*y))/(2*sigma*sigma))
			kernel[y+mid][x+mid] = value
			sum += value
		}
	}

	for y := range kernel {
		for x := range kernel[y] {
			kernel[y][x] /= sum
		}
	}

	return kernel
}
