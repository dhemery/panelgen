package shape

type Circle struct {
	XMLName string  `xml:"circle"`
	CX      float32 `xml:"cx,attr,omitempty"`
	CY      float32 `xml:"cy,attr,omitempty"`
	R       float32 `xml:"r,attr,omitempty"`
	Fill    *HSL    `xml:"fill,attr,omitempty"`
	Stroke  *HSL    `xml:"stroke,attr,omitempty"`
}

type Rect struct {
	XMLName string  `xml:"rect"`
	X       float32 `xml:"x,attr,omitempty"`
	Y       float32 `xml:"y,attr,omitempty"`
	Width   float32 `xml:"width,attr,omitempty"`
	Height  float32 `xml:"height,attr,omitempty"`
	RX      float32 `xml:"rx,attr,omitempty"`
	RY      float32 `xml:"ry,attr,omitempty"`
	Fill    *HSL    `xml:"fill,attr,omitempty"`
	Stroke  *HSL    `xml:"stroke,attr,omitempty"`
}
