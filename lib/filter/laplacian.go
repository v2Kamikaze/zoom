package filter

func Laplacian(size uint) [][]float64 {
	if size%2 == 0 {
		panic("O tamanho do kernel deve ser ímpar")
	}

	return [][]float64{
		{0, -1, 0},
		{-1, 4, -1},
		{0, -1, 0},
	}
}
