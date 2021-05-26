package module

import (
	"fmt"

	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

var All []*Module

type Module struct {
	Slug      string
	Faceplate shape.SVG
	Overlay   shape.SVG
	Frames    map[string]shape.SVG
}

func NewModule(slug string) *Module {
	return &Module{
		Slug:   slug,
		Frames: make(map[string]shape.SVG),
	}
}
func (m *Module) AddControl(c *control.Control) {
	m.Faceplate.Add(c.Faceplate)
	m.Overlay.Add(c.Overlay)
	for name, svg := range c.Frames {
		fmt.Println("Module", m.Slug, "adding", name)
		m.Frames[name] = svg
	}
}
