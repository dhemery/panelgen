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

func (m *Module) AddControl(c control.Control) {
	m.Faceplate.Content = append(m.Faceplate.Content, c.Faceplate, c.Overlay)
	m.Overlay.Content = append(m.Overlay.Content, c.Overlay)
	for name, svg := range c.Frames {
		m.Frames[name] = svg
	}
}
