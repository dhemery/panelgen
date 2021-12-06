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

func New(slug, name string, width float32, fg, bg shape.HSL) *Panel {
	faceplateRect := shape.Rect{
		X:           shape.StrokeWidth / 2,
		Y:           shape.StrokeWidth / 2,
		W:           width - shape.StrokeWidth,
		H:           Height - shape.StrokeWidth,
		Fill:        &bg,
		Stroke:      &fg,
		StrokeWidth: shape.StrokeWidth,
	}
	center := width / 2

	p := &Panel{
		Slug:       "cubic",
		Engravings: []shape.Bounded{faceplateRect},
	}

	brandLabel := LabelBelow("DHE", TitleFont, fg)
	p.Engrave(brandLabel, center, brandLabelY)
	nameLabel := LabelAbove(name, TitleFont, fg)
	p.Engrave(nameLabel, center, nameLabelY)
	return p
}

type Panel struct {
	Slug       string
	Engravings []shape.Bounded
	Controls   []control.Control
}

// Install installs the control at the specified position.
// The panel image will show the control's selected frame at that position.
// The module's svg directory will include an svg file for each frame of the control.
func (p *Panel) Install(c control.Control, x, y float32) control.Frame {
	installed := c.At(x, y)
	p.Controls = append(p.Controls, installed)
	return installed.DefaultFrame()
}

// Engrave engraves the shape into the faceplate at the specified position.
func (p *Panel) Engrave(s shape.Bounded, x, y float32) shape.Group {
	g := shape.NewGroupAt(x, y, s)
	p.Engravings = append(p.Engravings, g)
	return g
}

func (p *Panel) FaceplateSvg() shape.Svg {
	return shape.Svg{Content: p.Engravings}
}

func (p *Panel) ImageSvg() shape.Svg {
	svg := p.FaceplateSvg()
	for _, c := range p.Controls {
		svg.Content = append(svg.Content, c.DefaultFrame())
	}
	return svg
}

func (p *Panel) FrameSvgs() map[string]shape.Svg {
	frames := map[string]shape.Svg{}
	for _, control := range p.Controls {
		for slug, frame := range control.Frames {
			frames[slug] = frame.Svg()
		}
	}
	return frames
}
