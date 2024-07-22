package kernel_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/kernel"
)

func TestSobel(t *testing.T) {

	t.Run("sobel x and y", func(t *testing.T) {

		wantX := [][]float64{
			{-1, 0, 1},
			{-2, 0, 2},
			{-1, 0, 1},
		}

		wantY := [][]float64{
			{-1, -2, -1},
			{0, 0, 0},
			{1, 2, 1},
		}

		gotX := kernel.SobelX()
		gotY := kernel.SobelY()

		assert.Equal(t, wantX, gotX, "expect %+v, got %+v", wantX, gotX)
		assert.Equal(t, wantY, gotY, "expect %+v, got %+v", wantY, gotY)
	})

}
