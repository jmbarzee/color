package color

// Color is any representation of a color
type Color interface {
	// HSL returns a color represented by Hue, Saturation, and Lightness
	HSL() HSL
	// RGB returns a color represented by Red, Green, Blue
	RGB() RGB
}
