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
	return t.Width() * t.PortionRight
}
func (t Text) Bottom() float32 {
	return t.Height() * t.PortionBelow
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
