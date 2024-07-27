package kernel

func Laplacian(size uint) [][]float64 {
	if size%2 == 0 {
		panic("o tamanho do kernel deve ser ímpar")
	}

	kernel := make([][]float64, size)
	for i := range kernel {
		kernel[i] = make([]float64, size)
	}

	center := int(size / 2)
	totalNonCenterElements := int(size*size - 1)

	for y := 0; y < int(size); y++ {
		for x := 0; x < int(size); x++ {
			if x == center && y == center {
				kernel[y][x] = float64(totalNonCenterElements)
			} else {
				kernel[y][x] = -1.0
			}
		}
	}

	return kernel
}
