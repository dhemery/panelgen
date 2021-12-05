package panel

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

const (
	Height           = 127.5
	MillimetersPerHp = 5.08
)

type installedControl struct {
	Location shape.Location
	Control  control.Control
}

func New(slug, name string, hp int, fg, bg shape.HSL) *Panel {
	faceplateRect := shape.Rect{
		W:           float32(hp) * MillimetersPerHp,
		H:           Height,
		Fill:        &bg,
		Stroke:      &fg,
		StrokeWidth: 1.0,
	}
	// TODO: Add brand and module labels
	p := &Panel{
		Slug:       "cubic",
		Engravings: []shape.Bounded{faceplateRect},
	}
	return p
}

type Panel struct {
	Slug       string
	Engravings []shape.Bounded
	Controls   []installedControl
}

func (p *Panel) Install(c control.Control, x, y float32) {
	ic := installedControl{
		Location: shape.Location{X: x, Y: y},
		Control:  c,
	}
	p.Controls = append(p.Controls, ic)
}

func (p *Panel) Engrave(s shape.Bounded, x, y float32) {
	g := shape.G{
		Content: []shape.Bounded{s},
	}
	p.Engravings = append(p.Engravings, g.Translate(x, y))
}

func (p *Panel) Faceplate() shape.SVG {
	return shape.SVG{Content: p.Engravings}
}

func (p *Panel) Image() shape.SVG {
	svg := p.Faceplate()
	for _, c := range p.Controls {
		g := shape.G{
			Content: []shape.Bounded{c.Control.SelectedFrame()},
		}
		svg.Content = append(svg.Content, g.Translate(c.Location.X, c.Location.Y))
	}
	return svg
}

func (p *Panel) Frames() map[string]shape.SVG {
	return nil
}
