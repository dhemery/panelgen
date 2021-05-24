package control

import "dhemery.com/panelgen/shape"

// A Frame is an SVG with a slug.
type Frame struct {
	shape.SVG
	Name string `xml:"-"`
}

func (f Frame) Slug() string {
	return f.Name
}
