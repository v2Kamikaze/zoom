package convert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v2Kamikaze/zoom/lib/convert"
)

func TestColor(t *testing.T) {

	t.Run("rgb to hsv - white", func(t *testing.T) {
		wantH, wantS, wantV := 0.0, 0.0, 1.0

		gotH, gotS, gotV := convert.RGBToHSV(255, 255, 255)

		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("rgb to hsv - red", func(t *testing.T) {
		wantH, wantS, wantV := 0.0, 1.0, 1.0

		gotH, gotS, gotV := convert.RGBToHSV(255, 0, 0)

		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("rgb to hsv - green", func(t *testing.T) {
		wantH, wantS, wantV := 120.0, 1.0, 1.0

		gotH, gotS, gotV := convert.RGBToHSV(0, 255, 0)

		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("rgb to hsv - blue", func(t *testing.T) {
		wantH, wantS, wantV := 240.0, 1.0, 1.0

		gotH, gotS, gotV := convert.RGBToHSV(0, 0, 255)

		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("rgb to hsv - black", func(t *testing.T) {
		wantH, wantS, wantV := 0.0, 0.0, 0.0

		gotH, gotS, gotV := convert.RGBToHSV(0, 0, 0)

		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("hsv to rgb - white", func(t *testing.T) {
		wantR, wantG, wantB := uint8(255), uint8(255), uint8(255)

		gotR, gotG, gotB := convert.HSVToRGB(0.0, 0.0, 1.0)

		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("hsv to rgb - red", func(t *testing.T) {
		wantR, wantG, wantB := uint8(255), uint8(0), uint8(0)

		gotR, gotG, gotB := convert.HSVToRGB(0.0, 1.0, 1.0)

		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("hsv to rgb - green", func(t *testing.T) {
		wantR, wantG, wantB := uint8(0), uint8(255), uint8(0)

		gotR, gotG, gotB := convert.HSVToRGB(120.0, 1.0, 1.0)

		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("hsv to rgb - blue", func(t *testing.T) {
		wantR, wantG, wantB := uint8(0), uint8(0), uint8(255)

		gotR, gotG, gotB := convert.HSVToRGB(240.0, 1.0, 1.0)

		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("hsv to rgb - black", func(t *testing.T) {
		wantR, wantG, wantB := uint8(0), uint8(0), uint8(0)

		gotR, gotG, gotB := convert.HSVToRGB(0.0, 0.0, 0.0)

		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("to gray average - white", func(t *testing.T) {
		wantGray := uint8(255)
		gotGray := convert.ToGrayAverage(255, 255, 255)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray average - red", func(t *testing.T) {
		wantGray := uint8(85) // (255 + 0 + 0) / 3
		gotGray := convert.ToGrayAverage(255, 0, 0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray average - green", func(t *testing.T) {
		wantGray := uint8(85) // (0 + 255 + 0) / 3
		gotGray := convert.ToGrayAverage(0, 255, 0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray average - blue", func(t *testing.T) {
		wantGray := uint8(85) // (0 + 0 + 255) / 3
		gotGray := convert.ToGrayAverage(0, 0, 255)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray weighted - white", func(t *testing.T) {
		wantGray := uint8(255)
		gotGray := convert.ToGrayWeighted(255, 255, 255)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray weighted - red", func(t *testing.T) {
		v := 0.299*255 + 0.587*0 + 0.114*0
		wantGray := uint8(v) //
		gotGray := convert.ToGrayWeighted(255, 0, 0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray weighted - green", func(t *testing.T) {
		v := 0.299*0 + 0.587*255 + 0.114*0
		wantGray := uint8(v)
		gotGray := convert.ToGrayWeighted(0, 255, 0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray weighted - blue", func(t *testing.T) {
		v := 0.299*0 + 0.587*0 + 0.114*255
		wantGray := uint8(v)
		gotGray := convert.ToGrayWeighted(0, 0, 255)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray hsv - white", func(t *testing.T) {
		wantGray := 255.0
		gotGray := convert.ToGrayHSV(0.0, 0.0, 1.0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray hsv - red", func(t *testing.T) {
		wantGray := 76.24499999999999 // Converte HSV para RGB (255, 0, 0) e depois calcula a média ponderada
		gotGray := convert.ToGrayHSV(0.0, 1.0, 1.0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray hsv - green", func(t *testing.T) {
		wantGray := 149.685 // Converte HSV para RGB (0, 255, 0) e depois calcula a média ponderada
		gotGray := convert.ToGrayHSV(120.0, 1.0, 1.0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("to gray hsv - blue", func(t *testing.T) {
		wantGray := 29.07 // Converte HSV para RGB (0, 0, 255) e depois calcula a média ponderada
		gotGray := convert.ToGrayHSV(240.0, 1.0, 1.0)
		assert.Equal(t, wantGray, gotGray)
	})

	t.Run("negative rgb - white", func(t *testing.T) {
		wantR, wantG, wantB := uint8(0), uint8(0), uint8(0)
		gotR, gotG, gotB := convert.NegativeRGB(255, 255, 255)
		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("negative rgb - red", func(t *testing.T) {
		wantR, wantG, wantB := uint8(0), uint8(255), uint8(255)
		gotR, gotG, gotB := convert.NegativeRGB(255, 0, 0)
		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("negative rgb - green", func(t *testing.T) {
		wantR, wantG, wantB := uint8(255), uint8(0), uint8(255)
		gotR, gotG, gotB := convert.NegativeRGB(0, 255, 0)
		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("negative rgb - blue", func(t *testing.T) {
		wantR, wantG, wantB := uint8(255), uint8(255), uint8(0)
		gotR, gotG, gotB := convert.NegativeRGB(0, 0, 255)
		assert.Equal(t, wantR, gotR)
		assert.Equal(t, wantG, gotG)
		assert.Equal(t, wantB, gotB)
	})

	t.Run("negative hsv - white", func(t *testing.T) {
		wantH, wantS, wantV := 0.0, 0.0, 0.0
		gotH, gotS, gotV := convert.NegativeHSV(0.0, 0.0, 1.0)
		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("negative hsv - red", func(t *testing.T) {
		wantH, wantS, wantV := 180.0, 1.0, 1.0
		gotH, gotS, gotV := convert.NegativeHSV(0.0, 1.0, 1.0)
		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("negative hsv - green", func(t *testing.T) {
		wantH, wantS, wantV := 60.0, 1.0, 1.0
		gotH, gotS, gotV := convert.NegativeHSV(120.0, 1.0, 1.0)
		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})

	t.Run("negative hsv - blue", func(t *testing.T) {
		wantH, wantS, wantV := 300.0, 1.0, 1.0
		gotH, gotS, gotV := convert.NegativeHSV(240.0, 1.0, 1.0)
		assert.Equal(t, wantH, gotH)
		assert.Equal(t, wantS, gotS)
		assert.Equal(t, wantV, gotV)
	})
}
