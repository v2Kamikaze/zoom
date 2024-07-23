package effect

import (
	"image"
	"image/color"
	"math"
	"math/cmplx"
)

// IFFT2D realiza a Transformada Inversa de Fourier 2D.
func IFFT2D(input []complex128, width, height int) []complex128 {
	// IFFT em cada linha
	for y := 0; y < height; y++ {
		start := y * width
		IFFT1D(input[start : start+width])
	}

	// IFFT em cada coluna
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

	return input
}

// IFFT1D realiza a Transformada Inversa de Fourier 1D.
func IFFT1D(input []complex128) {
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

	// Realize IFFT recursivamente
	IFFT1D(even)
	IFFT1D(odd)

	// Combine os resultados
	for i := 0; i < n/2; i++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(i)/float64(n))) * odd[i]
		input[i] = even[i] + t
		input[i+n/2] = even[i] - t
	}

	// Normaliza os resultados
	for i := 0; i < n; i++ {
		input[i] /= complex(float64(n), 0)
	}
}

// CalculateInverseFourier calcula a Transformada Inversa de Fourier (IFT) de uma imagem
// e gera uma imagem com a transformação inversa aplicada.
func CalculateInverseFourier(img image.Image) image.Image {
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

	// Perform 2D IFFT
	result := IFFT2D(data, width, height)

	// Create a new grayscale image for the result
	resultImg := image.NewGray(bounds)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Convert complex result to real grayscale value
			realVal := cmplx.Abs(result[y*width+x])
			if realVal > 255 {
				realVal = 255
			}
			resultImg.SetGray(x, y, color.Gray{Y: uint8(realVal)})
		}
	}

	return resultImg
}
