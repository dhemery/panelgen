package svg

import (
	"encoding/xml"
	"fmt"
)

type Group struct {
	XMLName                  xml.Name `xml:"g"`
	Id                       string   `xml:"id,attr,omitempty"`
	Translation              string   `xml:"transform,attr,omitempty"`
	Elements                 []Bounded
	top, right, bottom, left float64
}

func GroupOf(elements ...Bounded) Group {
	b := Bounds(elements...)
	return Group{
		Elements: elements,
		top:      b.Top(),
		right:    b.Right(),
		bottom:   b.Bottom(),
		left:     b.Left(),
	}
}

func (g Group) Translate(x, y float64) Element {
	g.Translation = fmt.Sprintf("translate(%f %f)", x, y)
	g.top += y
	g.bottom += y
	g.left += x
	g.right += x
	return g
}

func (g Group) Top() float64 {
	return g.top
}

func (g Group) Right() float64 {
	return g.right
}

func (g Group) Bottom() float64 {
	return g.bottom
}

func (g Group) Left() float64 {
	return g.left
}

func (g Group) Width() float64 {
	return g.Right() - g.Left()
}

func (g Group) Height() float64 {
	return g.Bottom() - g.Top()
}
