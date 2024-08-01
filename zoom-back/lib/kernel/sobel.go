package kernel

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
