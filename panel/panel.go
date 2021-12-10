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
	outlineThickness := float32(0.5)
	faceplateRect := shape.Rect{
		X:           outlineThickness / 2,
		Y:           outlineThickness / 2,
		W:           width - outlineThickness,
		H:           Height - outlineThickness,
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

func (p *Panel) LabelAbove(x, y float32, text string, font shape.Font) {
	p.Engrave(x, y, LabelAbove(text, font, p.fg))
}

func (p *Panel) LabelBelow(x, y float32, text string, font shape.Font) {
	p.Engrave(x, y, LabelBelow(text, font, p.fg))
}

func (p *Panel) Port(x, y float32, name string, labelColor shape.HSL) {
	port := p.Install(x, y, control.Port(p.fg, p.bg))
	p.LabelAbove(x, port.Top()-shape.Padding, name, SmallFont)
}

func (p *Panel) CvPort(x, y float32) {
	p.Port(x, y, "CV", p.fg)
}

func (p *Panel) InPort(x, y float32, name string) {
	p.boxedPort(x, y, name, p.bg, p.fg)
}

func (p *Panel) OutPort(x, y float32, name string) {
	p.boxedPort(x, y, name, p.fg, p.bg)
}

func (p *Panel) SmallKnob(x, y float32, name string) {
	knob := p.Install(x, y, control.SmallKnob(p.fg, p.bg))
	labelY := knob.Top() - shape.Padding
	p.LabelAbove(x, labelY, name, SmallFont)
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
	return shape.NewSvg(p.Engravings)
}

func (p *Panel) ImageSvg() shape.Svg {
	content := p.Engravings
	for _, c := range p.Controls {
		content = append(content, c.DefaultFrame())
	}
	return shape.NewSvg(content)
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

func (p *Panel) boxedPort(x, y float32, name string, fill, labelColor shape.HSL) {
	barePort := control.Port(p.fg, p.bg)
	bareLabel := LabelAbove(name, SmallFont, labelColor)

	port := p.Install(x, y, barePort)
	labelY := port.Top() - shape.Padding

	box := shape.Rect{
		H:           port.Height() + bareLabel.Height() + 3*shape.Padding,
		W:           port.Width() + 2*shape.Padding,
		Fill:        &fill,
		Stroke:      &p.fg,
		StrokeWidth: shape.StrokeWidth,
		RX:          0.5,
		RY:          0.5,
	}

	boxTop := labelY - bareLabel.Height() - shape.Padding
	boxLeft := port.Left() - shape.Padding
	p.Engrave(boxLeft, boxTop, box)
	p.Engrave(x, labelY, bareLabel)
}
