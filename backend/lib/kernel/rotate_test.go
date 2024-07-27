package kernel_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/kernel"
)

func TestRotate(t *testing.T) {

	t.Run("kernel rotate", func(t *testing.T) {
		want := [][]float64{
			{6.123233995736757e-17, -1, 0},
			{1, 6.123233995736757e-17, 0},
			{0, 0, 1},
		}

		got := kernel.Rotate(90)

		assert.Equal(t, want, got, "expect %+v, got %+v", want, got)
	})

}
