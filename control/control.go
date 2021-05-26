package control

import "dhemery.com/panelgen/shape"

type Control struct {
	Faceplate shape.G
	Overlay   shape.G
	Frames    map[string]shape.SVG
}

func NewControl() *Control {
	return &Control{
		Frames: make(map[string]shape.SVG),
	}
}

// A Frame is an SVG with a slug.
type Frame struct {
	shape.SVG
	Name string `xml:"-"`
}

func (f Frame) Slug() string {
	return f.Name
}
