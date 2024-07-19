package histogram

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/v2Kamikaze/zoom/lib/intensity"
)

type HistogramChannel = uint8

const (
	R HistogramChannel = iota
	G
	B
	L
)

type Histogram struct {
	rChannel []uint
	gChannel []uint
	bChannel []uint
	lChannel []uint
}

// FromImage returns a new Histogram with all colors channel from given image.
func FromImage(img image.Image) *Histogram {
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

			luminance := uint8(0.299*float64(r8) + 0.587*float64(g8) + 0.114*float64(b8))
			histogramLuminance[luminance]++

		}
	}

	return &Histogram{
		rChannel: histogramR,
		gChannel: histogramG,
		bChannel: histogramB,
		lChannel: histogramLuminance,
	}
}

func (h *Histogram) RedChannel() []uint {
	return h.rChannel
}

func (h *Histogram) GreenChannel() []uint {
	return h.gChannel
}

func (h *Histogram) BlueChannel() []uint {
	return h.bChannel
}

func (h *Histogram) LuminanceChannel() []uint {
	return h.lChannel
}

func (h *Histogram) EqualizeWithChannel(img image.Image, channel HistogramChannel) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	var selectedHistogram []uint

	switch channel {
	case R:
		selectedHistogram = h.RedChannel()
	case G:
		selectedHistogram = h.GreenChannel()
	case B:
		selectedHistogram = h.BlueChannel()
	case L:
		selectedHistogram = h.LuminanceChannel()
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
		cdf[i] = math.Round((cdf[i] - cdfMin) * 255)
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, a := originalColor.RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			var newR, newG, newB uint8

			switch channel {
			case R:
				newR = uint8(cdf[r8])
				newG = g8
				newB = b8
			case G:
				newR = r8
				newG = uint8(cdf[g8])
				newB = b8
			case B:
				newR = r8
				newG = g8
				newB = uint8(cdf[b8])
			case L:
				luminance := intensity.ToGrayRGB(r, g, b)
				newGray := uint8(cdf[luminance])
				newR, newG, newB = newGray, newGray, newGray
			}

			newColor := color.RGBA{R: newR, G: newG, B: newB, A: uint8(a >> 8)}
			newImg.Set(x, y, newColor)
		}
	}

	return newImg
}
