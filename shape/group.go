package shape

import (
	"encoding/xml"
	"fmt"
)

type Group struct {
	XMLName xml.Name `xml:"g"`
	Vector
	Content []Bounded
}

func NewGroup(content ...Bounded) Group {
	return Group{Content: content}
}

func NewGroupAt(x, y float32, content ...Bounded) Group {
	return Group{Content: content, Vector: Vector{x, y}}
}

func (g Group) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

// TODO: Translate(Vector)
func (g Group) Translate(x, y float32) Group {
	g.X += x
	g.Y += y
	return g
}

func (g Group) Top() float32 {
	top := g.Y
	for _, c := range g.Content {
		if ct := c.Top() + g.Y; ct < top {
			top = ct
		}
	}
	return top
}

func (g Group) Right() float32 {
	right := g.X
	for _, c := range g.Content {
		if cr := c.Right() + g.Y; cr > right {
			right = cr
		}
	}
	return right
}

func (g Group) Bottom() float32 {
	bottom := g.Y
	for _, c := range g.Content {
		if cb := c.Bottom() + g.Y; cb > bottom {
			bottom = cb
		}
	}
	return bottom
}

func (g Group) Left() float32 {
	left := g.X
	for _, c := range g.Content {
		if cl := c.Left() + g.X; cl < left {
			left = cl
		}
	}
	return left
}

func (g Group) Width() float32 {
	return g.Right() - g.Left()
}

func (g Group) Height() float32 {
	return g.Bottom() - g.Top()
}
