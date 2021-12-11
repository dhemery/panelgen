package shape

func RectAround(fill, stroke HSL, strokeWidth float32, padding float32, elements ...Bounded) Rect {
	bounds := BoundsOf(elements...)
	return Rect{
		X:           bounds.Left - padding,
		Y:           bounds.Top - padding,
		H:           bounds.Height() + 2*padding,
		W:           bounds.Width() + 2*padding,
		Fill:        &fill,
		Stroke:      &stroke,
		StrokeWidth: strokeWidth,
		RX:          0.5,
		RY:          0.5,
	}
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
	return r.Y - r.StrokeWidth/2
}

func (r Rect) Right() float32 {
	return r.X + r.W + r.StrokeWidth/2
}

func (r Rect) Bottom() float32 {
	return r.Y + r.H + r.StrokeWidth/2
}

func (r Rect) Left() float32 {
	return r.X - r.StrokeWidth/2
}

func (r Rect) Width() float32 {
	return r.Right() - r.Left()
}

func (r Rect) Height() float32 {
	return r.Bottom() - r.Top()
}
