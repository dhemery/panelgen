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

type Panel struct {
	Slug       string
	Engravings []shape.Bounded
	Controls   []control.Control
	fg, bg     shape.HSL
}

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
		fg:         fg,
		bg:         bg,
	}

	p.LabelBelow(center, brandLabelY, "DHE", TitleFont)
	p.LabelAbove(center, nameLabelY, name, TitleFont)
	return p
}

func (p *Panel) LabelAbove(x, y float32, label string, font shape.Font) {
	p.Engrave(x, y, LabelAbove(label, font, p.fg))
}

func (p *Panel) LabelBelow(x, y float32, label string, font shape.Font) {
	p.Engrave(x, y, LabelBelow(label, font, p.fg))
}

func (p *Panel) Port(x, y float32, label string, labelColor shape.HSL) {
	port := p.Install(x, y, control.Port(p.fg))
	p.LabelAbove(x, port.Top()-shape.Padding, label, SmallFont)
}

func (p *Panel) CvPort(x, y float32) {
	p.Port(x, y, "CV", p.fg)
}

func (p *Panel) InPort(x, y float32, label string) {
	p.Port(x, y, label, p.fg)
}

func (p *Panel) OutPort(x, y float32, label string) {
	p.Port(x, y, label, p.fg)
}

func (p *Panel) SmallKnob(x, y float32, label string) {
	knob := p.Install(x, y, control.SmallKnob(p.fg, p.bg))
	labelY := knob.Top() - shape.Padding
	p.LabelAbove(x, labelY, label, SmallFont)
}

// Install installs the control at the specified position.
// The panel image will show the control's selected frame at that position.
// The module's svg directory will include an svg file for each frame of the control.
func (p *Panel) Install(x, y float32, c control.Control) control.Frame {
	installed := c.At(x, y)
	p.Controls = append(p.Controls, installed)
	return installed.DefaultFrame()
}

// Engrave engraves the shape into the faceplate at the specified position.
func (p *Panel) Engrave(x, y float32, s shape.Bounded) shape.Group {
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
