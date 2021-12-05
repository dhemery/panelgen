package shape

func LabelAbove(text string, font Font, color HSL) Bounded {
	return label(text, font, color, labelAbove)
}

func LabelBelow(text string, font Font, color HSL) Bounded {
	return label(text, font, color, labelBelow)
}

var (
	TitleFont = Font{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   titleFontSize,
	}
	LargeFont = Font{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   largeFontSize,
	}
	SmallFont = Font{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   smallFontSize,
	}
)

func label(text string, font Font, fill HSL, alignment labelAlignment) Bounded {
	return Text{
		Font:           font,
		labelAlignment: alignment,
		Fill:           &fill,
		Content:        text,
	}
}

const (
	titleFontSize = 12 / PixelsPerMillimeter
	largeFontSize = 9 / PixelsPerMillimeter
	smallFontSize = 7 / PixelsPerMillimeter
	ascentRatio   = 2 / 3 // Correct for Proxima Nova font
)

type Font struct {
	FontFamily string  `xml:"font-family,attr"`
	FontWeight string  `xml:"font-weight,attr"`
	FontSize   float32 `xml:"font-size,attr"`
}

type labelAlignment struct {
	DominantBaseline string `xml:"dominant-baseline,attr"`
	TextAnchor       string `xml:"text-anchor,attr"`
	portionBelow     float32
	portionRight     float32
	baselineShift    float32
}

var (
	labelAbove = labelAlignment{
		DominantBaseline: "alphabetic",
		TextAnchor:       "middle",
		portionBelow:     0,
		portionRight:     0.5,
		baselineShift:    0,
	}
	labelCenter = labelAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "middle",
		portionBelow:     0.5,
		portionRight:     0.5,
		baselineShift:    0.18,
	}
	labelLeft = labelAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "end",
		portionBelow:     0.5,
		portionRight:     0,
		baselineShift:    0.18,
	}
	labelRight = labelAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "start",
		portionBelow:     0.5,
		portionRight:     1,
		baselineShift:    0.18,
	}
	labelBelow = labelAlignment{
		DominantBaseline: "hanging",
		TextAnchor:       "middle",
		portionBelow:     1,
		portionRight:     0.5,
		baselineShift:    0.07,
	}
)

type Text struct {
	XMLName string `xml:"text"`
	Font
	labelAlignment
	Fill    *HSL   `xml:"fill,attr"`
	Content string `xml:",innerxml"`
}

func (t Text) Top() float32 {
	return t.Bottom() - t.Height()
}
func (t Text) Right() float32 {
	return t.Width() * t.portionRight
}
func (t Text) Bottom() float32 {
	return t.Height() * t.portionBelow
}
func (t Text) Left() float32 {
	return t.Right() - t.Width()
}
func (t Text) Width() float32 {
	return 0.1
}
func (t Text) Height() float32 {
	// The panels use only uppercase text, so the entire height of a <text> is its ascent
	return t.FontSize * ascentRatio
}
