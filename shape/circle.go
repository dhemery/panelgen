package shape

type Circle struct {
	XMLName     string  `xml:"circle"`
	CX          float64 `xml:"cx,attr,omitempty"`
	CY          float64 `xml:"cy,attr,omitempty"`
	R           float64 `xml:"r,attr,omitempty"`
	Fill        Color   `xml:"fill,attr"`
	Stroke      Color   `xml:"stroke,attr"`
	StrokeWidth float64 `xml:"stroke-width,attr,omitempty"`
}

func (c Circle) Top() float64 {
	return c.CY - c.R - c.StrokeWidth/2
}

func (c Circle) Right() float64 {
	return c.CX + c.R + c.StrokeWidth/2
}

func (c Circle) Bottom() float64 {
	return c.CY + c.R + c.StrokeWidth/2
}

func (c Circle) Left() float64 {
	return c.CX - c.R - c.StrokeWidth/2
}

func (c Circle) Width() float64 {
	return c.Right() - c.Left()
}

func (c Circle) Height() float64 {
	return c.Bottom() - c.Top()
}
