package shape

import (
	"encoding/xml"
	"fmt"
)

type Bounded interface {
	Top() float32
	Right() float32
	Bottom() float32
	Left() float32
	Width() float32
	Height() float32
}

type container struct {
	t, r, b, l float32
	content    []Bounded
}

func (c *container) Add(shapes ...Bounded) {
	c.content = append(c.content, shapes...)
	for _, s := range shapes {
		if s.Top() < c.t {
			c.t = s.Top()
		}
		if s.Right() > c.r {
			c.r = s.Right()
		}
		if s.Bottom() > c.b {
			c.b = s.Bottom()
		}
		if s.Left() < c.l {
			c.l = s.Left()
		}
	}
}

func (c container) Top() float32 {
	return c.t
}

func (c container) Right() float32 {
	return c.r
}

func (c container) Bottom() float32 {
	return c.b
}

func (c container) Left() float32 {
	return c.l
}

func (c container) Width() float32 {
	return c.r - c.l
}

func (c container) Height() float32 {
	return c.b - c.t
}

type SVG struct {
	container
}

func (s SVG) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	version := xml.Attr{Name: xml.Name{Local: "version"}, Value: "1.1"}
	xmlns := xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.w3.org/2000/svg"}
	widthMM := fmt.Sprintf("%fmm", s.Width())
	width := xml.Attr{Name: xml.Name{Local: "width"}, Value: widthMM}
	heightMM := fmt.Sprintf("%fmm", s.Height())
	height := xml.Attr{Name: xml.Name{Local: "height"}, Value: heightMM}
	vb := fmt.Sprintf("%f %f %f %f", s.Left(), s.Top(), s.Width(), s.Height())
	viewBox := xml.Attr{Name: xml.Name{Local: "viewBox"}, Value: vb}

	start.Name = xml.Name{Local: "svg"}
	start.Attr = append(start.Attr, version, xmlns, width, height, viewBox)

	fmt.Printf("Marshaling container %+v\n\n", s.container)
	return e.EncodeElement(s.container.content, start)
}

type G struct {
	XMLName string `xml:"g"`
	container
}

func (g G) Translate(x, y float32) G {
	return g
}
