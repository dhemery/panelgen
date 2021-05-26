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

type bounds struct {
	Top, Right, Bottom, Left float32
}

func (b bounds) Width() float32 {
	return b.Right - b.Left
}

func (b bounds) Height() float32 {
	return b.Bottom - b.Top
}

func boundsOf(shapes []Bounded) bounds {
	var b bounds
	if len(shapes) < 1 {
		return b
	}
	first := shapes[0]
	b.Top = first.Top()
	b.Right = first.Right()
	b.Bottom = first.Bottom()
	b.Left = first.Left()
	for _, s := range shapes[1:] {
		if v := s.Top(); v < b.Top {
			b.Top = v
		}
		if v := s.Right(); v > b.Right {
			b.Right = v
		}
		if v := s.Bottom(); v > b.Bottom {
			b.Bottom = v
		}
		if v := s.Left(); v < b.Left {
			b.Left = v
		}
	}
	return b
}

type SVG struct {
	Content []Bounded
}

func (s SVG) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	b := boundsOf(s.Content)
	version := xml.Attr{Name: xml.Name{Local: "version"}, Value: "1.1"}
	xmlns := xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.w3.org/2000/svg"}
	widthMM := fmt.Sprintf("%fmm", b.Width())
	width := xml.Attr{Name: xml.Name{Local: "width"}, Value: widthMM}
	heightMM := fmt.Sprintf("%fmm", b.Height())
	height := xml.Attr{Name: xml.Name{Local: "height"}, Value: heightMM}
	vb := fmt.Sprintf("%f %f %f %f", b.Left, b.Top, b.Width(), b.Height())
	viewBox := xml.Attr{Name: xml.Name{Local: "viewBox"}, Value: vb}

	start.Attr = append(start.Attr, version, xmlns, width, height, viewBox)
	start.Name = xml.Name{Local: "svg"}

	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.Encode(s.Content); err != nil {
		return err
	}
	return e.EncodeToken(start.End())
}

type G struct {
	XMLName xml.Name `xml:"g"`
	Content []Bounded
}

func (g G) Top() float32 {
	var v float32
	for _, c := range g.Content {
		if c.Top() < v {
			v = c.Top()
		}
	}
	return v
}

func (g G) Right() float32 {
	var v float32
	for _, c := range g.Content {
		if c.Right() > v {
			v = c.Right()
		}
	}
	return v
}

func (g G) Bottom() float32 {
	var v float32
	for _, c := range g.Content {
		if c.Bottom() > v {
			v = c.Bottom()
		}
	}
	return v
}

func (g G) Left() float32 {
	var v float32
	for _, c := range g.Content {
		if c.Left() < v {
			v = c.Left()
		}
	}
	return v
}

func (g G) Width() float32 {
	return g.Right() - g.Left()
}

func (g G) Height() float32 {
	return g.Bottom() - g.Top()
}
