package control

import "dhemery.com/panelgen/shape"

type Frame struct {
	shape.SVG
	Name string `xml:"-"`
}

func (f Frame) Slug() string {
	return f.Name
}
