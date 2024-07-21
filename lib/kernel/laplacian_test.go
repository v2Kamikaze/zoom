package kernel_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/kernel"
)

func TestLaplacian(t *testing.T) {

	t.Run("3x3 laplacian", func(t *testing.T) {
		got := [][]float64{
			{-1, -1, -1},
			{-1, 8, -1},
			{-1, -1, -1},
		}

		want := kernel.Laplacian(3)

		assert.Equal(t, want, got, "expect %+v, got %+v", want, got)
	})

}
