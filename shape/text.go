package shape

type Font struct {
	FontFamily  string  `xml:"font-family,attr"`
	FontWeight  string  `xml:"font-weight,attr"`
	FontSize    float64 `xml:"font-size,attr"`
	AscentRatio float64 `xml:"-"`
}

type TextAlignment struct {
	DominantBaseline string  `xml:"dominant-baseline,attr"`
	TextAnchor       string  `xml:"text-anchor,attr"`
	PortionBelow     float64 `xml:"-"`
	PortionRight     float64 `xml:"-"`
	BaselineShift    float64 `xml:"-"`
}

func TextAbove(x, y float64, content string, font Font, color Color) Text {
	return newText(x, y, content, alignTextAbove, font, color)
}

func TextBelow(x, y float64, content string, font Font, color Color) Text {
	return newText(x, y, content, alignTextBelow, font, color)
}

func TextCentered(x, y float64, content string, font Font, color Color) Text {
	return newText(x, y, content, alignTextCentered, font, color)
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

func newText(x, y float64, content string, alignment TextAlignment, font Font, color Color) Text {
	return Text{
		X:             x,
		Y:             y + font.FontSize*font.AscentRatio*alignment.BaselineShift,
		Font:          font,
		TextAlignment: alignment,
		Fill:          color,
		Content:       content,
	}
}

type Text struct {
	XMLName string `xml:"text"`
	Font
	TextAlignment
	X       float64 `xml:"x,attr,omitempty"`
	Y       float64 `xml:"y,attr,omitempty"`
	Fill    Color   `xml:"fill,attr"`
	Content string  `xml:",innerxml"`
}

func (t Text) Top() float64 {
	return t.Bottom() - t.Height()
}
func (t Text) Right() float64 {
	return t.X + t.Width()*t.PortionRight
}
func (t Text) Bottom() float64 {
	return t.Y + t.Height()*t.PortionBelow
}
func (t Text) Left() float64 {
	return t.Right() - t.Width()
}
func (t Text) Width() float64 {
	return 0.1
}
func (t Text) Height() float64 {
	// The panels use only uppercase text, so the entire height of a <text> is its ascent
	return t.FontSize * t.AscentRatio
}
