package main

import (
	"fmt"

	"github.com/v2Kamikaze/zoom/lib/background"
	"github.com/v2Kamikaze/zoom/lib/imageio"
	"github.com/v2Kamikaze/zoom/lib/kernel"
	"github.com/v2Kamikaze/zoom/lib/zoom"
)

func main() {

	infos := []struct {
		name string
		ext  string
	}{
		{name: "zero", ext: "png"},
		{name: "go", ext: "jpg"},
	}

	runner := background.NewBackgroundRunner()

	runner.Add(func() {

	})

	for _, info := range infos {
		img := imageio.OpenImage(fmt.Sprintf("./assets/%s.%s", info.name, info.ext))
		hist := zoom.FromImage(img)

		runner.Add(func() {
			sobelXImg := zoom.ApplySobelX(img)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-sobel-x.%s", info.name, info.ext), sobelXImg)
		})

		runner.Add(func() {
			sobelYImg := zoom.ApplySobelY(img)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-sobel-y.%s", info.name, info.ext), sobelYImg)
		})

		runner.Add(func() {
			sobelImg := zoom.ApplySobel(img)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-sobel-full.%s", info.name, info.ext), sobelImg)
		})

		runner.Add(func() {
			negImg := zoom.ApplyNeg(img)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-neg.%s", info.name, info.ext), negImg)
		})

		runner.Add(func() {
			gammImg := zoom.ApplyGamma(img, 2.2, 1.0)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-gamma.%s", info.name, info.ext), gammImg)
		})

		runner.Add(func() {
			binImg := zoom.ApplyBin(img, 128)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-bin.%s", info.name, info.ext), binImg)
		})

		runner.Add(func() {
			meanImg := zoom.ApplyMean(img, 5)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-mean.%s", info.name, info.ext), meanImg)
		})

		runner.Add(func() {
			gaussImg := zoom.ApplyGaussian(img, 5, 1.5)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-gauss.%s", info.name, info.ext), gaussImg)
		})

		runner.Add(func() {
			laplaImg := zoom.ApplyLaplacian(img, 5)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-laplace.%s", info.name, info.ext), laplaImg)
		})

		runner.Add(func() {
			highBoostImg := zoom.ApplyHighBoost(img, 1.5, kernel.Gaussian(5, 1.5))
			imageio.SaveImage(fmt.Sprintf("./assets/%s-high-boost.%s", info.name, info.ext), highBoostImg)
		})

		runner.Add(func() {
			imgHistR := hist.EqualizeWithChannel(img, zoom.RChannel)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-r.%s", info.name, info.ext), imgHistR)
		})

		runner.Add(func() {
			imgHistG := hist.EqualizeWithChannel(img, zoom.GChannel)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-g.%s", info.name, info.ext), imgHistG)
		})

		runner.Add(func() {
			imgHistB := hist.EqualizeWithChannel(img, zoom.BChannel)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-b.%s", info.name, info.ext), imgHistB)
		})

		runner.Add(func() {
			imgHistL := hist.EqualizeWithChannel(img, zoom.LChannel)
			imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-l.%s", info.name, info.ext), imgHistL)
		})

	}

	runner.RunAndWait()
	runner.Clear()
}
