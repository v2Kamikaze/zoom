package intensity

func ToGrayRGB(r, g, b uint32) uint8 {
	return uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
}
