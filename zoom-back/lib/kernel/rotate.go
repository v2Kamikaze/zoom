package kernel

import "math"

func Rotate(angle float64) [][]float64 {
	rad := angle * math.Pi / 180

	return [][]float64{
		{math.Cos(rad), -math.Sin(rad), 0},
		{math.Sin(rad), math.Cos(rad), 0},
		{0, 0, 1},
	}
}
