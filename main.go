package main

import (
	"github.com/v2Kamikaze/zoom/lib/convolution"
	"github.com/v2Kamikaze/zoom/lib/filter"
	"github.com/v2Kamikaze/zoom/lib/histogram"
	"github.com/v2Kamikaze/zoom/lib/imageio"
	"github.com/v2Kamikaze/zoom/lib/intensity"
)

var meanKernel = filter.Mean(5)
var gaussianKernel = filter.Gaussian(5, 1.6)
var laplacianKernel = filter.Laplacian(5)

func main() {
	img := imageio.OpenImage("./assets/zero.png")

	negImg := intensity.Negative(img)
	gammImg := intensity.GammaCorrection(img, 2.2, 1.0)
	binImg := intensity.Binarize(img, 128)
	meanImg := convolution.Convolve(img, meanKernel)
	gaussImg := convolution.Convolve(img, gaussianKernel)
	laplaImg := convolution.Convolve(img, laplacianKernel)

	hist := histogram.FromImage(img)
	imgHistR := hist.EqualizeWithChannel(img, histogram.R)
	imgHistG := hist.EqualizeWithChannel(img, histogram.G)
	imgHistB := hist.EqualizeWithChannel(img, histogram.B)
	imgHistL := hist.EqualizeWithChannel(img, histogram.L)

	imageio.SaveImage("./assets/zero-neg.png", negImg)
	imageio.SaveImage("./assets/zero-gamma.png", gammImg)
	imageio.SaveImage("./assets/zero-bin.png", binImg)
	imageio.SaveImage("./assets/zero-mean.png", meanImg)
	imageio.SaveImage("./assets/zero-gauss.png", gaussImg)
	imageio.SaveImage("./assets/zero-laplace.png", laplaImg)
	imageio.SaveImage("./assets/zero-hist-r.png", imgHistR)
	imageio.SaveImage("./assets/zero-hist-g.png", imgHistG)
	imageio.SaveImage("./assets/zero-hist-b.png", imgHistB)
	imageio.SaveImage("./assets/zero-hist-l.png", imgHistL)

	img = imageio.OpenImage("./assets/go.jpg")

	negImg = intensity.Negative(img)
	gammImg = intensity.GammaCorrection(img, 2.2, 1.0)
	binImg = intensity.Binarize(img, 128)
	meanImg = convolution.Convolve(img, meanKernel)
	gaussImg = convolution.Convolve(img, gaussianKernel)
	laplaImg = convolution.Convolve(img, laplacianKernel)

	hist = histogram.FromImage(img)
	imgHistR = hist.EqualizeWithChannel(img, histogram.R)
	imgHistG = hist.EqualizeWithChannel(img, histogram.G)
	imgHistB = hist.EqualizeWithChannel(img, histogram.B)
	imgHistL = hist.EqualizeWithChannel(img, histogram.L)

	imageio.SaveImage("./assets/go-neg.jpg", negImg)
	imageio.SaveImage("./assets/go-gamma.jpg", gammImg)
	imageio.SaveImage("./assets/go-bin.jpg", binImg)
	imageio.SaveImage("./assets/go-mean.jpg", meanImg)
	imageio.SaveImage("./assets/go-gauss.jpg", gaussImg)
	imageio.SaveImage("./assets/go-laplace.jpg", laplaImg)
	imageio.SaveImage("./assets/go-hist-r.jpg", imgHistR)
	imageio.SaveImage("./assets/go-hist-g.jpg", imgHistG)
	imageio.SaveImage("./assets/go-hist-b.jpg", imgHistB)
	imageio.SaveImage("./assets/go-hist-l.jpg", imgHistL)

}
