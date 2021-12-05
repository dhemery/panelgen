package shape

import (
	"encoding/xml"
	"fmt"
)

type Location struct {
	X float32
	Y float32
}

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
	if len(shapes) < 1 {
		return bounds{}
	}
	first := shapes[0]
	b := bounds{
		Top:    first.Top(),
		Right:  first.Right(),
		Bottom: first.Bottom(),
		Left:   first.Left(),
	}
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
	Location
	Content []Bounded
}

func NewG(content ...Bounded) G {
	return G{Content: content}
}

func NewGAt(x, y float32, content ...Bounded) G {
	return G{Content: content, Location: Location{x, y}}
}

func (g G) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if g.X != 0 || g.Y != 0 {
		tx := xml.Attr{Name: xml.Name{Local: "transform"}, Value: fmt.Sprintf("translate(%f %f)", g.X, g.Y)}
		start.Attr = append(start.Attr, tx)
	}
	start.Name = xml.Name{Local: "g"}

	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.Encode(g.Content); err != nil {
		return err
	}
	return e.EncodeToken(start.End())
}

func (g G) Translate(x, y float32) G {
	g.X = x
	g.Y = y
	return g
}

func (g G) Top() float32 {
	top := g.Y
	for _, c := range g.Content {
		if ct := c.Top() + g.Y; ct < top {
			top = ct
		}
	}
	return top
}

func (g G) Right() float32 {
	right := g.X
	for _, c := range g.Content {
		if cr := c.Right() + g.Y; cr > right {
			right = cr
		}
	}
	return right
}

func (g G) Bottom() float32 {
	bottom := g.Y
	for _, c := range g.Content {
		if cb := c.Bottom() + g.Y; cb > bottom {
			bottom = cb
		}
	}
	return bottom
}

func (g G) Left() float32 {
	left := g.X
	for _, c := range g.Content {
		if cl := c.Left() + g.X; cl < left {
			left = cl
		}
	}
	return left
}

func (g G) Width() float32 {
	return g.Right() - g.Left()
}

func (g G) Height() float32 {
	return g.Bottom() - g.Top()
}
