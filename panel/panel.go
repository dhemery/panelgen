package panel

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

const (
	Height      = 128.5
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

	p := &Panel{
		Slug:       "cubic",
		Engravings: []shape.Bounded{faceplateRect},
	}

	brandLabel := shape.LabelBelow("DHE", shape.TitleFont, fg)
	p.Engrave(brandLabel, center, brandLabelY)
	nameLabel := shape.LabelAbove(name, shape.TitleFont, fg)
	p.Engrave(nameLabel, center, nameLabelY)
	return p
}

type Panel struct {
	Slug       string
	Engravings []shape.Bounded
	Controls   []installedControl
}

// Install installs the control at the specified position.
// The panel image will show the control's selected frame at that position.
// The module's svg directory will include an svg file for each frame of the control.
func (p *Panel) Install(c control.Control, x, y float32) shape.G {
	ic := installedControl{
		Location: shape.Location{X: x, Y: y},
		Control:  c,
	}
	p.Controls = append(p.Controls, ic)
	return shape.NewGAt(x, y, c.SelectedFrame())
}

// Engrave engraves the shape into the faceplate at the specified position.
func (p *Panel) Engrave(s shape.Bounded, x, y float32) shape.G {
	g := shape.NewGAt(x, y, s)
	p.Engravings = append(p.Engravings, g)
	return g
}

func (p *Panel) Faceplate() shape.SVG {
	return shape.SVG{Content: p.Engravings}
}

func (p *Panel) Image() shape.SVG {
	svg := p.Faceplate()
	for _, c := range p.Controls {
		g := shape.NewGAt(c.Location.X, c.Location.Y, c.Control.SelectedFrame())
		svg.Content = append(svg.Content, g)
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
