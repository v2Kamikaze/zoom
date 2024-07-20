package main

import (
	"fmt"

	"github.com/v2Kamikaze/zoom/lib/filter"
	"github.com/v2Kamikaze/zoom/lib/histogram"
	"github.com/v2Kamikaze/zoom/lib/imageio"
	"github.com/v2Kamikaze/zoom/lib/zoom"
)

type imageInfo struct {
	name string
	ext  string
}

func main() {

	infos := []imageInfo{
		{name: "zero", ext: "png"},
		{name: "go", ext: "jpg"},
	}

	for _, info := range infos {

		img := imageio.OpenImage(fmt.Sprintf("./assets/%s.%s", info.name, info.ext))

		negImg := zoom.ApplyNeg(img)
		gammImg := zoom.ApplyGamma(img, 2.2, 1.0)
		binImg := zoom.ApplyBin(img, 128)
		meanImg := zoom.ApplyMean(img, 5)
		gaussImg := zoom.ApplyGaussian(img, 5, 1.5)
		laplaImg := zoom.ApplyLaplacian(img, 5)
		highBoostImg := zoom.ApplyHighBoost(img, 1.5, filter.Gaussian(5, 1.5))

		hist := histogram.FromImage(img)
		imgHistR := hist.EqualizeWithChannel(img, histogram.R)
		imgHistG := hist.EqualizeWithChannel(img, histogram.G)
		imgHistB := hist.EqualizeWithChannel(img, histogram.B)
		imgHistL := hist.EqualizeWithChannel(img, histogram.L)

		imageio.SaveImage(fmt.Sprintf("./assets/%s-neg.%s", info.name, info.ext), negImg)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-gamma.%s", info.name, info.ext), gammImg)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-bin.%s", info.name, info.ext), binImg)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-mean.%s", info.name, info.ext), meanImg)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-gauss.%s", info.name, info.ext), gaussImg)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-laplace.%s", info.name, info.ext), laplaImg)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-high-boost.%s", info.name, info.ext), highBoostImg)

		imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-r.%s", info.name, info.ext), imgHistR)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-g.%s", info.name, info.ext), imgHistG)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-b.%s", info.name, info.ext), imgHistB)
		imageio.SaveImage(fmt.Sprintf("./assets/%s-hist-l.%s", info.name, info.ext), imgHistL)
	}

}
