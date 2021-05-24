package module

import (
	"dhemery.com/panelgen/shape"
)

var All []*Module

type Module struct {
	slug      string
	faceplate shape.SVG
	overlay   shape.SVG
	controls  map[string]shape.SVG
}

func (m *Module) AddFaceplate(s []shape.Bounded) {
	m.faceplate.Add(s...)
}

func (m *Module) AddOverlay(s []shape.Bounded) {
	m.overlay.Add(s...)
}

func (m *Module) AddControl(path string, svg shape.SVG) {
	if m.controls == nil {
		m.controls = make(map[string]shape.SVG)
	}
	m.controls[path] = svg
}

func (m Module) Slug() string {
	return m.slug
}

func (m Module) Faceplate() shape.SVG {
	return m.faceplate
}

func (m Module) Image() shape.SVG {
	return m.overlay
}

func (m Module) Controls() map[string]shape.SVG {
	return m.controls
}
