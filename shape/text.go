package shape

var (
	LabelAbove = TextAlignment{
		DominantBaseline: "alphabetic",
		TextAnchor:       "middle",
		portionBelow:     0,
		portionRight:     0.5,
		baselineShift:    0,
	}
	LabelCenter = TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "middle",
		portionBelow:     0.5,
		portionRight:     0.5,
		baselineShift:    0.18,
	}
	LabelLeft = TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "end",
		portionBelow:     0.5,
		portionRight:     0,
		baselineShift:    0.18,
	}
	LabelRight = TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "start",
		portionBelow:     0.5,
		portionRight:     1,
		baselineShift:    0.18,
	}
	LabelBelow = TextAlignment{
		DominantBaseline: "hanging",
		TextAnchor:       "middle",
		portionBelow:     1,
		portionRight:     0.5,
		baselineShift:    0.07,
	}
)

var (
	TitleFont = TextFont{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   12 / PixelsPerMillimeter,
	}
	LargeFont = TextFont{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   9 / PixelsPerMillimeter,
	}
	SmallFont = TextFont{
		FontFamily: "Proxima Nova",
		FontWeight: "bold",
		FontSize:   7 / PixelsPerMillimeter,
	}
)

func Label(text string, font TextFont, fill HSL, alignment TextAlignment) Bounded {
	return Text{
		TextFont:      font,
		TextAlignment: alignment,
		Fill:          &fill,
		Content:       text,
	}
}

type TextAlignment struct {
	DominantBaseline string `xml:"dominant-baseline,attr"`
	TextAnchor       string `xml:"text-anchor,attr"`
	portionBelow     float32
	portionRight     float32
	baselineShift    float32
}

const (
	ascentRatio = 2 / 3 // Correct for Proxima Nova font
)

type TextFont struct {
	FontFamily string  `xml:"font-family,attr"`
	FontWeight string  `xml:"font-weight,attr"`
	FontSize   float32 `xml:"font-size,attr"`
}

type Text struct {
	XMLName string `xml:"text"`
	TextFont
	TextAlignment
	Fill    *HSL   `xml:"fill,attr"`
	Content string `xml:",chardata"`
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
	return t.FontSize * ascentRatio
}
