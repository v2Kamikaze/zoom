package kernel_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/kernel"
)

func TestGaussian(t *testing.T) {

	t.Run("3x3 gaussian", func(t *testing.T) {
		got := [][]float64{
			{0.0947416582101747, 0.1183180127031206, 0.0947416582101747},
			{0.1183180127031206, 0.14776131634681883, 0.1183180127031206},
			{0.0947416582101747, 0.1183180127031206, 0.0947416582101747},
		}

		want := kernel.Gaussian(3, 1.5)

		assert.Equal(t, want, got, "expect %+v, got %+v", want, got)
	})

}
