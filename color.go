package color

type Color interface {
	// HSL returns a color represented by Hue, Saturation, and Lightness
	HSL()
	// RGB returns a color represented by Red, Green, Blue
	RGB()
}
