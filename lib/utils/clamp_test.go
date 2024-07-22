package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/utils"
)

func TestClamp(t *testing.T) {

	t.Run("clamp valid value", func(t *testing.T) {
		want := 125

		got := utils.Clamp(125, 0, 255)

		assert.Equal(t, want, got, "expect %+v, got %+v", want, got)
	})

	t.Run("clamp valid value", func(t *testing.T) {
		want := 0

		got := utils.Clamp(-5, 0, 255)

		assert.Equal(t, want, got, "expect %+v, got %+v", want, got)
	})

	t.Run("clamp valid value", func(t *testing.T) {
		want := 255

		got := utils.Clamp(256, 0, 255)

		assert.Equal(t, want, got, "expect %+v, got %+v", want, got)
	})

}
