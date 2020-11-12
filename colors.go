package color

import (
	"math"
)

var (
	Min  = math.SmallestNonzeroFloat64
	Max  = 1.0 - Min
	Half = 0.5
)

var (
	Black        = HSL{Min, Min, Min}
	Grey         = HSL{Min, Min, Half}
	GreyNatural  = HSL{Min, Min, Min}
	White        = HSL{Min, Min, Max}
	WhiteNatural = HSL{Min, Min, Min}

	Red         = HSL{Min, Max, Half}
	WarmRed     = HSL{1.0 / 24, Max, Half}
	Orange      = HSL{2.0 / 24, Max, Half}
	WarmYellow  = HSL{3.0 / 24, Max, Half}
	Yellow      = HSL{4.0 / 24, Max, Half}
	CoolYellow  = HSL{5.0 / 24, Max, Half}
	YellowGreen = HSL{6.0 / 24, Max, Half}
	WarmGreen   = HSL{7.0 / 24, Max, Half}
	Green       = HSL{8.0 / 24, Max, Half}
	CoolGreen   = HSL{9.0 / 24, Max, Half}
	GreenCyan   = HSL{10.0 / 24, Max, Half}
	WarmCyan    = HSL{11.0 / 24, Max, Half}
	Cyan        = HSL{12.0 / 24, Max, Half}
	CoolCyan    = HSL{13.0 / 24, Max, Half}
	BlueCyan    = HSL{14.0 / 24, Max, Half}
	CoolBlue    = HSL{15.0 / 24, Max, Half}
	Blue        = HSL{16.0 / 24, Max, Half}
	WarmBlue    = HSL{17.0 / 24, Max, Half}
	Violet      = HSL{18.0 / 24, Max, Half}
	CoolMagenta = HSL{19.0 / 24, Max, Half}
	Magenta     = HSL{20.0 / 24, Max, Half}
	WarmMagenta = HSL{21.0 / 24, Max, Half}
	RedMagenta  = HSL{22.0 / 24, Max, Half}
	CoolRed     = HSL{23.0 / 24, Max, Half}

	AllColors = []Color{
		Red,
		WarmRed,
		Orange,
		WarmYellow,
		Yellow,
		CoolYellow,
		YellowGreen,
		WarmGreen,
		Green,
		CoolGreen,
		GreenCyan,
		WarmCyan,
		Cyan,
		CoolCyan,
		BlueCyan,
		CoolBlue,
		Blue,
		WarmBlue,
		Violet,
		CoolMagenta,
		Magenta,
		WarmMagenta,
		RedMagenta,
		CoolRed,
	}
)
