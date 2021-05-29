package module

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

var All []Module

type Module struct {
	Slug      string
	Faceplate shape.SVG
	Overlay   shape.SVG
	Frames    map[string]shape.SVG
}

func (m *Module) AddControl(c control.Control, x, y float32) {
	if f := c.Faceplate; len(f.Content) > 0 {
		m.Faceplate.Content = append(m.Faceplate.Content, f.Translate(x, y))
	}
	if o := c.Overlay; len(o.Content) > 0 {
		o := c.Overlay.Translate(x, y)
		m.Faceplate.Content = append(m.Faceplate.Content, o)
		m.Overlay.Content = append(m.Overlay.Content, o)
	}
	for name, svg := range c.Frames {
		m.Frames[name] = svg
	}
}
