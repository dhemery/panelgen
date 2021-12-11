package shape

import (
	"encoding/xml"
	"fmt"
)

type Group struct {
	XMLName                  xml.Name `xml:"g"`
	Id                       string   `xml:"id,attr,omitempty"`
	Translation              string   `xml:"transform,attr,omitempty"`
	Elements                 []Bounded
	top, right, bottom, left float32
}

func NewGroup(elements ...Bounded) Group {
	b := BoundsOf(elements...)
	return Group{
		Elements: elements,
		top:      b.Top,
		right:    b.Right,
		bottom:   b.Bottom,
		left:     b.Left,
	}
}

func (g Group) Translate(x, y float32) Group {
	g.Translation = fmt.Sprintf("translate(%f %f)", x, y)
	g.top += y
	g.bottom += y
	g.left += x
	g.right += x
	return g
}

func (g Group) Top() float32 {
	return g.top
}

func (g Group) Right() float32 {
	return g.right
}

func (g Group) Bottom() float32 {
	return g.bottom
}

func (g Group) Left() float32 {
	return g.left
}

func (g Group) Width() float32 {
	return g.Right() - g.Left()
}

func (g Group) Height() float32 {
	return g.Bottom() - g.Top()
}
