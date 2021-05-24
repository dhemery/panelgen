package module

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

var All []*Module

type Module struct {
	slug      string
	faceplate *shape.SVG
	overlay   *shape.SVG
	controls  []control.Frame
}

func (m Module) Slug() string {
	return m.slug
}

func (m Module) Faceplate() *shape.SVG {
	return m.faceplate
}

func (m Module) Image() *shape.SVG {
	return m.overlay
}

func (m Module) Controls() []control.Frame {
	return m.controls
}
