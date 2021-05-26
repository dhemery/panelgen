package shape

type S struct {
}

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
	return c.CY - c.R - c.StrokeWidth
}

func (c Circle) Right() float32 {
	return c.CX + c.R + c.StrokeWidth
}

func (c Circle) Bottom() float32 {
	return c.CY + c.R + c.StrokeWidth
}

func (c Circle) Left() float32 {
	return c.CX - c.R - c.StrokeWidth
}

func (c Circle) Width() float32 {
	return c.Right() - c.Left()
}

func (c Circle) Height() float32 {
	return c.Bottom() - c.Top()
}

type Rect struct {
	XMLName     string  `xml:"rect"`
	X           float32 `xml:"x,attr,omitempty"`
	Y           float32 `xml:"y,attr,omitempty"`
	W           float32 `xml:"width,attr,omitempty"`
	H           float32 `xml:"height,attr,omitempty"`
	RX          float32 `xml:"rx,attr,omitempty"`
	RY          float32 `xml:"ry,attr,omitempty"`
	Fill        *HSL    `xml:"fill,attr,omitempty"`
	Stroke      *HSL    `xml:"stroke,attr,omitempty"`
	StrokeWidth float32 `xml:"stroke-width,attr,omitempty"`
}

func (r Rect) Top() float32 {
	return r.Y - r.StrokeWidth
}

func (r Rect) Right() float32 {
	return r.X + r.W + r.StrokeWidth
}

func (r Rect) Bottom() float32 {
	return r.Y + r.H + r.StrokeWidth
}

func (r Rect) Left() float32 {
	return r.X - r.StrokeWidth
}

func (r Rect) Width() float32 {
	return r.Right() - r.Left()
}

func (r Rect) Height() float32 {
	return r.Bottom() - r.Top()
}
