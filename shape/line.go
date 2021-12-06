package shape

import "math"

// TODO: Translate(Vector)
type Line struct {
	XMLName     string  `xml:"line"`
	X1          float32 `xml:"x1,attr,omitempty"`
	Y1          float32 `xml:"y1,attr,omitempty"`
	X2          float32 `xml:"x2,attr,omitempty"`
	Y2          float32 `xml:"y2,attr,omitempty"`
	Stroke      *HSL    `xml:"stroke,attr,omitempty"`
	StrokeWidth float32 `xml:"stroke-width,attr,omitempty"`
	Cap         string  `xml:"stroke-linecap,attr,omitempty"`
}

func (l Line) Top() float32 {
	return float32(math.Min(float64(l.Y1), float64(l.Y2)))
}

func (l Line) Right() float32 {
	return float32(math.Max(float64(l.X1), float64(l.X2)))
}

func (l Line) Bottom() float32 {
	return float32(math.Max(float64(l.Y1), float64(l.Y2)))
}

func (l Line) Left() float32 {
	return float32(math.Min(float64(l.X1), float64(l.X2)))
}

func (l Line) Width() float32 {
	return l.Right() - l.Left()
}

func (l Line) Height() float32 {
	return l.Bottom() - l.Top()
}
