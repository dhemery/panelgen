package panel

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/svg"
)

const (
	boxRadius   = 0.5
	padding     = 1
	strokeWidth = 0.35
)

type Hp int

func (hp Hp) toMM() float64 {
	const mmPerHp = 5.08
	return float64(hp) * mmPerHp
}

type Panel struct {
	Slug       string
	Engravings []svg.Bounded
	Frames     []svg.Bounded
	Controls   []control.Control
	Fg, Bg     svg.Color
	Hp         Hp
}

func New(slug, name string, hp Hp, fg, bg svg.Color) *Panel {
	const (
		nameLabelY       = 9
		outlineThickness = 0.5
		panelHeight      = 128.5
		brandLabelY      = panelHeight - nameLabelY
	)

	p := &Panel{
		Slug: slug,
		Fg:   fg,
		Bg:   bg,
		Hp:   hp,
	}

	faceplateRect := svg.Rect{
		X:           outlineThickness / 2,
		Y:           outlineThickness / 2,
		W:           p.Width() - outlineThickness,
		H:           panelHeight - outlineThickness,
		Fill:        bg,
		Stroke:      fg,
		StrokeWidth: outlineThickness,
	}
	p.Engrave(faceplateRect)
	center := p.Width() / 2

	p.Engrave(svg.TextBelow("DHE", svg.TitleFont, p.Fg).Translate(center, brandLabelY))
	p.Engrave(svg.TextAbove(name, svg.TitleFont, p.Fg).Translate(center, nameLabelY))
	return p
}

func (p *Panel) Width() float64 {
	return p.Hp.toMM()
}

func (p *Panel) Port(x, y float64, name string, labelColor svg.Color) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	labelY := port.Top() - padding
	p.Engrave(svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY))
}

func (p *Panel) CvPort(x, y float64) {
	p.Port(x, y, "CV", p.Fg)
}

func (p *Panel) InPort(x, y float64, name string) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	labelY := port.Top() - padding
	label := svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY)
	p.Engrave(boxAround(p.Bg, p.Fg, port, label))
	p.Engrave(label)
}

func (p *Panel) InButtonPort(x, y float64, name string) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	buttonX := x + control.PortRadius + control.ButtonRadius + padding
	button := p.Install(buttonX, y, control.Button(p.Bg, p.Fg))
	labelY := port.Top() - padding
	label := svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY)
	p.Engrave(boxAround(p.Bg, p.Fg, port, button, label))
	p.Engrave(label)
}

func (p *Panel) OutPort(x, y float64, name string) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	labelY := port.Top() - padding
	label := svg.TextAbove(name, svg.SmallFont, p.Bg).Translate(x, labelY)
	p.Engrave(boxAround(p.Fg, p.Fg, port, label))
	p.Engrave(label)
}

func (p *Panel) OutButtonPort(x, y float64, name string) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	buttonX := x - control.PortRadius - control.ButtonRadius - padding
	button := p.Install(buttonX, y, control.OutputButton(p.Fg, p.Bg))
	labelY := port.Top() - padding
	label := svg.TextAbove(name, svg.SmallFont, p.Bg).Translate(x, labelY)
	p.Engrave(boxAround(p.Fg, p.Fg, port, button, label))
	p.Engrave(label)
}

func (p *Panel) SmallKnob(x, y float64, name string) {
	knob := p.Install(x, y, control.SmallKnob(p.Fg, p.Bg))
	labelY := knob.Top() - padding
	p.Engrave(svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY))
}

func (p *Panel) LargeKnob(x, y float64, name string) {
	knob := p.Install(x, y, control.LargeKnob(p.Fg, p.Bg))
	labelY := knob.Top() - padding
	p.Engrave(svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY))
}

// Install installs the control at the specified position.
// The panel image will show the control's selected frame at that position.
// The module's svg directory will include an svg file for each frame of the control.
func (p *Panel) Install(x, y float64, c control.Control) svg.Element {
	p.Controls = append(p.Controls, c)
	frame := c.DefaultFrame.Translate(x, y)
	p.Frames = append(p.Frames, frame)
	return frame
}

// Engrave engraves the shape into the faceplate at the specified position.
func (p *Panel) Engrave(s svg.Bounded) {
	p.Engravings = append(p.Engravings, s)
}

func (p *Panel) FaceplateSvg() svg.Svg {
	return svg.NewSvg(p.Engravings...)
}

func (p *Panel) ImageSvg() svg.Svg {
	faceplateGroup := svg.GroupOf(p.Engravings...)
	faceplateGroup.Id = "faceplate"
	imageGroup := svg.GroupOf(p.Frames...)
	imageGroup.Id = "image"
	return svg.NewSvg(faceplateGroup, imageGroup)
}

func (p *Panel) FrameSvgs() map[string]svg.Svg {
	frames := map[string]svg.Svg{}
	for _, control := range p.Controls {
		for slug, frame := range control.Frames {
			frames[slug] = svg.NewSvg(frame)
		}
	}
	return frames
}

func boxAround(fill, stroke svg.Color, elements ...svg.Bounded) svg.Rect {
	bounds := svg.Bounds(elements...)
	return svg.Rect{
		X:           bounds.Left() - padding,
		Y:           bounds.Top() - padding,
		H:           bounds.Height() + 2*padding,
		W:           bounds.Width() + 2*padding,
		Fill:        fill,
		Stroke:      stroke,
		StrokeWidth: strokeWidth,
		RX:          boxRadius,
		RY:          boxRadius,
	}
}
