package effect

import (
	"image"
	"image/color"
	"math"
	"math/cmplx"
)

// IFFT2D realiza a Transformada Rápida de Fourier Inversa 2D.
func IFFT2D(input []complex128, width, height int) []complex128 {
	for y := 0; y < height; y++ {
		start := y * width
		IFFT1D(input[start : start+width])
	}

	for x := 0; x < width; x++ {
		column := make([]complex128, height)
		for y := 0; y < height; y++ {
			column[y] = input[y*width+x]
		}
		IFFT1D(column)
		for y := 0; y < height; y++ {
			input[y*width+x] = column[y]
		}
	}

	// Dividir cada valor pelo número total de elementos para normalizar
	for i := range input {
		input[i] /= complex(float64(width*height), 0)
	}

	return input
}

// IFFT1D realiza a Transformada Rápida de Fourier Inversa 1D.
func IFFT1D(input []complex128) {
	n := len(input)
	if n <= 1 {
		return
	}

	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = input[i*2]
		odd[i] = input[i*2+1]
	}

	IFFT1D(even)
	IFFT1D(odd)

	for i := 0; i < n/2; i++ {
		t := cmplx.Exp(complex(0, 2*math.Pi*float64(i)/float64(n))) * odd[i]
		input[i] = even[i] + t
		input[i+n/2] = even[i] - t
	}
}

// UncenterFFT desloca o espectro de volta ao seu estado original.
func UncenterFFT(input []complex128, width, height int) []complex128 {
	uncentered := make([]complex128, len(input))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newX := (x + width/2) % width
			newY := (y + height/2) % height
			uncentered[y*width+x] = input[newY*width+newX]
		}
	}

	return uncentered
}

// CalculateInverseFourier reconstrói a imagem original a partir do espectro.
func CalculateInverseFourier(spectrumImg image.Image) image.Image {
	bounds := spectrumImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	data := make([]complex128, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := color.GrayModel.Convert(spectrumImg.At(x, y)).(color.Gray)
			data[y*width+x] = complex(float64(gray.Y), 0)
		}
	}

	uncentered := UncenterFFT(data, width, height)
	reconstructed := IFFT2D(uncentered, width, height)

	resultImg := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			val := real(reconstructed[y*width+x])
			if val < 0 {
				val = 0
			}
			if val > 255 {
				val = 255
			}
			resultImg.SetGray(x, y, color.Gray{Y: uint8(val)})
		}
	}

	return resultImg
}
