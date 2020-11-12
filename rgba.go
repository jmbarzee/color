package color

import "math"

// RGBA is an RGB with an Alpha portion
type RGBA struct {
	RGB
	A float64
}

// ToUInt32WRGB converts an RGBA to a uint32.
// Each portion of the color is represented by 8 bits
func (c RGBA) ToUInt32WRGB() uint32 {
	val := uint32(c.R*math.MaxUint8) << 0
	val |= uint32(c.G*math.MaxUint8) << 8
	val |= uint32(c.B*math.MaxUint8) << 16
	val |= uint32(c.A*math.MaxUint8) << 24
	return val
}

// FromUInt32WRGB converts and uint32 to an RGBA.
// Each portion of the color is represented by 8 bits
func FromUInt32WRGB(wgrb uint32) RGBA {
	mask := uint32(0x000000ff)

	uint8r := mask & (wgrb >> 0)
	r := float64(uint8r) / math.MaxUint8
	uint8g := mask & (wgrb >> 8)
	g := float64(uint8g) / math.MaxUint8
	uint8b := mask & (wgrb >> 16)
	b := float64(uint8b) / math.MaxUint8
	uint8a := mask & (wgrb >> 24)
	a := float64(uint8a) / math.MaxUint8
	return RGBA{
		RGB: RGB{
			R: r,
			G: g,
			B: b,
		},
		A: a,
	}
}
