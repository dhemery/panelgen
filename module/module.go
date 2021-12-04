package module

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

var All []*Module

const (
	height  = 127.5
	mmPerHp = 5.08
)

type Module struct {
	Slug      string
	Width     float32
	Faceplate shape.SVG
	Overlay   shape.SVG
	Frames    map[string]shape.SVG
}

func NewModule(slug string, hp int, fg, bg shape.HSL) *Module {
	m := &Module{
		Slug:   slug,
		Width:  float32(hp) * mmPerHp,
		Frames: make(map[string]shape.SVG),
	}

	faceplateRect := shape.Rect{
		W:           m.Width,
		H:           height,
		X:           0,
		Y:           0,
		Fill:        &bg,
		Stroke:      &fg,
		StrokeWidth: 1.0,
	}
	m.AddToFaceplate(faceplateRect)
	return m
}

func (m *Module) AddControl(c control.Control, x, y float32) {
	if f := c.Faceplate; len(f.Content) > 0 {
		m.AddToFaceplate(c.Faceplate.Translate(x, y))
	}
	if o := c.Overlay; len(o.Content) > 0 {
		m.AddToOverlay(c.Overlay.Translate(x, y))
	}
	for name, svg := range c.Frames {
		m.Frames[name] = svg
	}
}

func (m *Module) AddToFaceplate(content shape.Bounded) {
	m.Faceplate.Content = append(m.Faceplate.Content, content)
	m.AddToOverlay(content)

}

func (m *Module) AddToOverlay(content shape.Bounded) {
	m.Overlay.Content = append(m.Overlay.Content, content)
}
