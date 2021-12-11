package shape

type Font struct {
	FontFamily  string  `xml:"font-family,attr"`
	FontWeight  string  `xml:"font-weight,attr"`
	FontSize    float32 `xml:"font-size,attr"`
	AscentRatio float32 `xml:"-"`
}

type TextAlignment struct {
	DominantBaseline string  `xml:"dominant-baseline,attr"`
	TextAnchor       string  `xml:"text-anchor,attr"`
	PortionBelow     float32 `xml:"-"`
	PortionRight     float32 `xml:"-"`
	BaselineShift    float32 `xml:"-"`
}

func TextAbove(x, y float32, content string, font Font, color HSL) Text {
	return Text{
		X:             x,
		Y:             y,
		Font:          font,
		TextAlignment: alignTextAbove,
		Fill:          &color,
		Content:       content,
	}
}

func TextBelow(x, y float32, content string, font Font, color HSL) Text {
	return Text{
		X:             x,
		Y:             y,
		Font:          font,
		TextAlignment: alignTextBelow,
		Fill:          &color,
		Content:       content,
	}
}

var (
	TitleFont = Font{
		FontFamily:  "Proxima Nova",
		FontWeight:  "bold",
		FontSize:    titleFontSize,
		AscentRatio: proximaNovaAscentRatio,
	}
	LargeFont = Font{
		FontFamily:  "Proxima Nova",
		FontWeight:  "bold",
		FontSize:    largeFontSize,
		AscentRatio: proximaNovaAscentRatio,
	}
	SmallFont = Font{
		FontFamily:  "Proxima Nova",
		FontWeight:  "bold",
		FontSize:    smallFontSize,
		AscentRatio: proximaNovaAscentRatio,
	}
)

func newText(x, y float32, content string, font Font, fill HSL, alignment TextAlignment) Text {
	return Text{
		X:             x,
		Y:             y,
		Font:          font,
		TextAlignment: alignment,
		Fill:          &fill,
		Content:       content,
	}
}

const (
	mmPerInch = 25.4
	pxPerInch = 75.0
	pxPerMm   = pxPerInch / mmPerInch
)

const (
	titleFontSize          = 12.0 / pxPerMm
	largeFontSize          = 9.0 / pxPerMm
	smallFontSize          = 7.0 / pxPerMm
	proximaNovaAscentRatio = 2.0 / 3.0
)

var (
	alignTextAbove = TextAlignment{
		DominantBaseline: "alphabetic",
		TextAnchor:       "middle",
		PortionBelow:     0,
		PortionRight:     0.5,
		BaselineShift:    0,
	}
	alignTextCentered = TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "middle",
		PortionBelow:     0.5,
		PortionRight:     0.5,
		BaselineShift:    0.18,
	}
	alignTextLeft = TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "end",
		PortionBelow:     0.5,
		PortionRight:     0,
		BaselineShift:    0.18,
	}
	alignTextRight = TextAlignment{
		DominantBaseline: "middle",
		TextAnchor:       "start",
		PortionBelow:     0.5,
		PortionRight:     1,
		BaselineShift:    0.18,
	}
	alignTextBelow = TextAlignment{
		DominantBaseline: "hanging",
		TextAnchor:       "middle",
		PortionBelow:     1,
		PortionRight:     0.5,
		BaselineShift:    0.07,
	}
)

type Text struct {
	XMLName string `xml:"text"`
	Font
	TextAlignment
	X       float32 `xml:"x,attr,omitempty"`
	Y       float32 `xml:"y,attr,omitempty"`
	Fill    *HSL    `xml:"fill,attr"`
	Content string  `xml:",innerxml"`
}

func (t Text) Top() float32 {
	return t.Bottom() - t.Height()
}
func (t Text) Right() float32 {
	return t.X + t.Width()*t.PortionRight
}
func (t Text) Bottom() float32 {
	return t.Y + t.Height()*t.PortionBelow
}
func (t Text) Left() float32 {
	return t.Right() - t.Width()
}
func (t Text) Width() float32 {
	return 0.1
}
func (t Text) Height() float32 {
	// The panels use only uppercase text, so the entire height of a <text> is its ascent
	return t.FontSize * t.AscentRatio
}
