package filter

func Mean(size uint) [][]float64 {
	if size%2 == 0 {
		panic("O tamanho do kernel deve ser ímpar")
	}

	kernel := make([][]float64, size)
	value := 1.0 / float64(size*size)

	for i := range kernel {
		kernel[i] = make([]float64, size)
		for j := range kernel[i] {
			kernel[i][j] = value
		}
	}

	return kernel
}
