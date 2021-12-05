package panel

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

const (
	Height      = 127.5
	nameLabelY  = 9
	brandLabelY = Height - nameLabelY
)

type installedControl struct {
	Location shape.Location
	Control  control.Control
}

func New(slug, name string, width float32, fg, bg shape.HSL) *Panel {
	faceplateRect := shape.Rect{
		W:           width,
		H:           Height,
		Fill:        &bg,
		Stroke:      &fg,
		StrokeWidth: shape.StrokeWidth,
	}
	center := width / 2
	brandLabel := shape.Label("DHE", shape.TitleFont, fg, shape.LabelBelow)
	brandG := shape.G{Content: []shape.Bounded{brandLabel}}.Translate(center, brandLabelY)
	nameLabel := shape.Label(name, shape.TitleFont, fg, shape.LabelAbove)
	nameG := shape.G{Content: []shape.Bounded{nameLabel}}.Translate(center, nameLabelY)

	p := &Panel{
		Slug: "cubic",
		Engravings: []shape.Bounded{
			faceplateRect, brandG, nameG},
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
	frames := map[string]shape.SVG{}
	for _, installation := range p.Controls {
		for slug, content := range installation.Control.Frames {
			svg := shape.SVG{Content: []shape.Bounded{content}}
			frames[slug] = svg
		}
	}
	return frames
}
