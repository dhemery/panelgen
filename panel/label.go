package panel

import "dhemery.com/panelgen/shape"

func LabelAbove(text string, font shape.Font, color shape.HSL) shape.Bounded {
	return label(text, font, color, labelAbove)
}

func LabelBelow(text string, font shape.Font, color shape.HSL) shape.Bounded {
	return label(text, font, color, labelBelow)
}

var (
	TitleFont = shape.Font{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   titleFontSize,
	}
	LargeFont = shape.Font{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   largeFontSize,
	}
	SmallFont = shape.Font{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   smallFontSize,
	}
)

func label(text string, font shape.Font, fill shape.HSL, alignment shape.TextAlignment) shape.Bounded {
	return shape.Text{
		Font:          font,
		TextAlignment: alignment,
		Fill:          &fill,
		Content:       text,
	}
}

const (
	titleFontSize = 12 / shape.PixelsPerMillimeter
	largeFontSize = 9 / shape.PixelsPerMillimeter
	smallFontSize = 7 / shape.PixelsPerMillimeter
	ascentRatio   = 2 / 3 // Correct for Proxima Nova font
)

var (
	labelAbove = shape.TextAlignment{
		DominantBaseline: "alphabetic",
		TextAnchor:       "middle",
		PortionBelow:     0,
		PortionRight:     0.5,
		BaselineShift:    0,
	}
	labelCenter = shape.TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "middle",
		PortionBelow:     0.5,
		PortionRight:     0.5,
		BaselineShift:    0.18,
	}
	labelLeft = shape.TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "end",
		PortionBelow:     0.5,
		PortionRight:     0,
		BaselineShift:    0.18,
	}
	labelRight = shape.TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "start",
		PortionBelow:     0.5,
		PortionRight:     1,
		BaselineShift:    0.18,
	}
	labelBelow = shape.TextAlignment{
		DominantBaseline: "hanging",
		TextAnchor:       "middle",
		PortionBelow:     1,
		PortionRight:     0.5,
		BaselineShift:    0.07,
	}
)
