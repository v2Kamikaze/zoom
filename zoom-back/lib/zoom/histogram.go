package zoom

import (
	"image"
	"image/color"
	"log"

	"github.com/v2Kamikaze/zoom/lib/utils"
)

type HistogramChannel = rune

const (
	RChannel HistogramChannel = iota
	GChannel
	BChannel
	LChannel
)

type Histogram struct {
	R []uint `json:"r"`
	G []uint `json:"g"`
	B []uint `json:"b"`
	L []uint `json:"l"`
}

func HistogramFromImage(img image.Image) *Histogram {
	bounds := img.Bounds()
	histogramR := make([]uint, 256)
	histogramG := make([]uint, 256)
	histogramB := make([]uint, 256)
	histogramLuminance := make([]uint, 256)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			histogramR[r8]++
			histogramG[g8]++
			histogramB[b8]++

			luminance := utils.ToGrayRGB(r, g, b)
			histogramLuminance[luminance]++
		}
	}

	return &Histogram{
		R: histogramR,
		G: histogramG,
		B: histogramB,
		L: histogramLuminance,
	}
}

func (h *Histogram) EqualizeWithChannel(img image.Image, channel HistogramChannel) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	var selectedHistogram []uint

	switch channel {
	case RChannel:
		selectedHistogram = h.R
	case GChannel:
		selectedHistogram = h.G
	case BChannel:
		selectedHistogram = h.B
	case LChannel:
		selectedHistogram = h.L
	default:
		log.Panicf("canal inválido. %d", channel)
	}

	totalPixels := bounds.Dx() * bounds.Dy()

	cdf := make([]float64, 256)
	cdf[0] = float64(selectedHistogram[0]) / float64(totalPixels)
	for i := 1; i < 256; i++ {
		cdf[i] = cdf[i-1] + float64(selectedHistogram[i])/float64(totalPixels)
	}

	cdfMin := cdf[0]
	for i := 0; i < 256; i++ {
		cdf[i] = ((cdf[i] - cdfMin) / (1.0 - cdfMin)) * 255
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			var newR, newG, newB uint8

			switch channel {
			case RChannel:
				newR = uint8(cdf[r8])
				newG = g8
				newB = b8
			case GChannel:
				newR = r8
				newG = uint8(cdf[g8])
				newB = b8
			case BChannel:
				newR = r8
				newG = g8
				newB = uint8(cdf[b8])
			case LChannel:
				luminance := utils.ToGrayRGB(r, g, b)
				newGray := uint8(cdf[luminance])
				newR, newG, newB = newGray, newGray, newGray
			}

			newColor := color.RGBA{R: newR, G: newG, B: newB, A: uint8(a >> 8)}
			newImg.Set(x, y, newColor)
		}
	}

	return newImg
}
