package shape

// TODO: Translate(Vector)
type Circle struct {
	XMLName     string  `xml:"circle"`
	CX          float32 `xml:"cx,attr,omitempty"`
	CY          float32 `xml:"cy,attr,omitempty"`
	R           float32 `xml:"r,attr,omitempty"`
	Fill        *HSL    `xml:"fill,attr"`
	Stroke      *HSL    `xml:"stroke,attr"`
	StrokeWidth float32 `xml:"stroke-width,attr,omitempty"`
}

func (c Circle) Top() float32 {
	return c.CY - c.R - c.StrokeWidth/2
}

func (c Circle) Right() float32 {
	return c.CX + c.R + c.StrokeWidth/2
}

func (c Circle) Bottom() float32 {
	return c.CY + c.R + c.StrokeWidth/2
}

func (c Circle) Left() float32 {
	return c.CX - c.R - c.StrokeWidth/2
}

func (c Circle) Width() float32 {
	return c.Right() - c.Left()
}

func (c Circle) Height() float32 {
	return c.Bottom() - c.Top()
}
