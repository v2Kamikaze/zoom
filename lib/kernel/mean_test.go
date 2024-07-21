package kernel_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/kernel"
)

func TestMean(t *testing.T) {

	t.Run("3x3 mean", func(t *testing.T) {
		got := [][]float64{
			{0.1111111111111111, 0.1111111111111111, 0.1111111111111111},
			{0.1111111111111111, 0.1111111111111111, 0.1111111111111111},
			{0.1111111111111111, 0.1111111111111111, 0.1111111111111111},
		}

		want := kernel.Mean(3)

		assert.Equal(t, want, got, "expect %+v, got %+v", want, got)
	})

}
