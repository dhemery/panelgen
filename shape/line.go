package shape

import "math"

type Line struct {
	XMLName     string  `xml:"line"`
	X1          float64 `xml:"x1,attr,omitempty"`
	Y1          float64 `xml:"y1,attr,omitempty"`
	X2          float64 `xml:"x2,attr,omitempty"`
	Y2          float64 `xml:"y2,attr,omitempty"`
	Stroke      Color   `xml:"stroke,attr"`
	StrokeWidth float64 `xml:"stroke-width,attr,omitempty"`
	Cap         string  `xml:"stroke-linecap,attr,omitempty"`
}

func (l Line) Top() float64 {
	return math.Min(l.Y1, l.Y2) - l.StrokeWidth
}

func (l Line) Right() float64 {
	return math.Max(l.X1, l.X2) + l.StrokeWidth
}

func (l Line) Bottom() float64 {
	return math.Max(l.Y1, l.Y2) + l.StrokeWidth
}

func (l Line) Left() float64 {
	return math.Min(l.X1, l.X2) - l.StrokeWidth
}

func (l Line) Width() float64 {
	return l.Right() - l.Left()
}

func (l Line) Height() float64 {
	return l.Bottom() - l.Top()
}
