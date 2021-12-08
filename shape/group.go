package shape

import (
	"encoding/xml"
	"fmt"
)

type Group struct {
	XMLName xml.Name `xml:"g"`
	X       float32
	Y       float32
	Content []Bounded
}

func NewGroup(content ...Bounded) Group {
	return Group{Content: content}
}

func NewGroupAt(x, y float32, content ...Bounded) Group {
	return Group{Content: content, X: x, Y: y}
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

func (g Group) At(x, y float32) Group {
	g.X += x
	g.Y += y
	return g
}

func (g Group) Top() float32 {
	top := float32(20000)
	for _, c := range g.Content {
		if ct := c.Top(); ct < top {
			top = ct
		}
	}
	return top + g.Y
}

func (g Group) Right() float32 {
	right := float32(-20000)
	for _, c := range g.Content {
		if cr := c.Right(); cr > right {
			right = cr
		}
	}
	return right + g.X
}

func (g Group) Bottom() float32 {
	bottom := float32(-20000)
	for _, c := range g.Content {
		if cb := c.Bottom(); cb > bottom {
			bottom = cb
		}
	}
	return bottom + g.Y
}

func (g Group) Left() float32 {
	left := float32(20000)
	for _, c := range g.Content {
		if cl := c.Left(); cl < left {
			left = cl
		}
	}
	return left + g.X
}

func (g Group) Width() float32 {
	return g.Right() - g.Left()
}

func (g Group) Height() float32 {
	return g.Bottom() - g.Top()
}
