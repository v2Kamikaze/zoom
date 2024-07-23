package effect

import (
	"image"
	"image/color"
	"math"
	"math/cmplx"
)

// FFT2D realiza a Transformada Rápida de Fourier 2D.
func FFT2D(input []complex128, width, height int) []complex128 {
	// FFT em cada linha
	for y := 0; y < height; y++ {
		start := y * width
		FFT1D(input[start : start+width])
	}

	// FFT em cada coluna
	for x := 0; x < width; x++ {
		column := make([]complex128, height)
		for y := 0; y < height; y++ {
			column[y] = input[y*width+x]
		}
		FFT1D(column)
		for y := 0; y < height; y++ {
			input[y*width+x] = column[y]
		}
	}

	return input
}

// FFT1D realiza a Transformada Rápida de Fourier 1D.
func FFT1D(input []complex128) {
	n := len(input)
	if n <= 1 {
		return
	}

	// Divida a entrada em duas partes
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = input[i*2]
		odd[i] = input[i*2+1]
	}

	// Realize FFT recursivamente
	FFT1D(even)
	FFT1D(odd)

	// Combine os resultados
	for i := 0; i < n/2; i++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(i)/float64(n))) * odd[i]
		input[i] = even[i] + t
		input[i+n/2] = even[i] - t
	}
}

// CalculateFourier calcula a Transformada Discreta de Fourier (DFT) de uma imagem,
// desloca o espectro e aplica uma escala logarítmica para a visualização.
func CalculateFourier(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Convert image to grayscale
	grayscale := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			grayscale.Set(x, y, gray)
		}
	}

	// Create a 2D slice to hold the grayscale pixel values
	data := make([]complex128, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			data[y*width+x] = complex(float64(grayscale.GrayAt(x, y).Y), 0)
		}
	}

	// Perform 2D FFT
	result := FFT2D(data, width, height)

	// Calculate magnitude and apply log scale
	magnitude := make([]float64, width*height)
	for i, c := range result {
		magnitude[i] = 20 * math.Log10(cmplx.Abs(c)+1)
	}

	// Create a new grayscale image for the result
	resultImg := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mag := magnitude[y*width+x]
			if mag > 255 {
				mag = 255
			}
			resultImg.SetGray(x, y, color.Gray{Y: uint8(mag)})
		}
	}

	return resultImg
}

func CenterFFT(input []complex128, width, height int) []complex128 {
	centered := make([]complex128, len(input))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Calcula o índice da nova posição centralizada
			newX := (x + width/2) % width
			newY := (y + height/2) % height
			centered[newY*width+newX] = input[y*width+x]
		}
	}

	return centered
}

// CalculateFourier calcula a Transformada Discreta de Fourier (DFT) de uma imagem,
// desloca o espectro para o centro e aplica uma escala logarítmica para a visualização.
func CalculateCenterFourier(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Convert image to grayscale
	grayscale := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			grayscale.Set(x, y, gray)
		}
	}

	// Create a 2D slice to hold the grayscale pixel values
	data := make([]complex128, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			data[y*width+x] = complex(float64(grayscale.GrayAt(x, y).Y), 0)
		}
	}

	// Perform 2D FFT
	result := FFT2D(data, width, height)

	// Center the FFT result
	centeredResult := CenterFFT(result, width, height)

	// Calculate magnitude and apply log scale
	magnitude := make([]float64, width*height)
	for i, c := range centeredResult {
		magnitude[i] = 20 * math.Log10(cmplx.Abs(c)+1)
	}

	// Create a new grayscale image for the result
	resultImg := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mag := magnitude[y*width+x]
			if mag > 255 {
				mag = 255
			}
			resultImg.SetGray(x, y, color.Gray{Y: uint8(mag)})
		}
	}

	return resultImg
}
