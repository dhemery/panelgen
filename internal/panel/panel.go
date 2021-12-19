package panel

import (
	"dhemery.com/panelgen/internal/control"
	"dhemery.com/panelgen/internal/svg"
)

type Hp int

func (hp Hp) toMM() float64 {
	const mmPerHp = 5.08
	return float64(hp) * mmPerHp
}

type Panel struct {
	Engravings []svg.Element
	Frames     []svg.Element
	Controls   []control.Control
	Fg, Bg     svg.Color
	Width      float64
}

const (
	padding     = 1.0
	strokeWidth = 0.35
)

func NewPanel(name string, hp Hp, fg, bg svg.Color) *Panel {
	const (
		nameLabelY       = 9.0
		outlineThickness = 0.5
		panelHeight      = 128.5
		brandLabelY      = panelHeight - nameLabelY
	)

	p := &Panel{
		Fg:    fg,
		Bg:    bg,
		Width: hp.toMM(),
	}

	faceplateRect := svg.Rect{
		X:           outlineThickness / 2,
		Y:           outlineThickness / 2,
		W:           p.Width - outlineThickness,
		H:           panelHeight - outlineThickness,
		Fill:        bg,
		Stroke:      fg,
		StrokeWidth: outlineThickness,
	}
	p.Engrave(faceplateRect)
	center := p.Width / 2

	p.Engrave(svg.TextAbove(name, svg.TitleFont, p.Fg).Translate(center, nameLabelY))
	p.Engrave(svg.TextBelow("DHE", svg.TitleFont, p.Fg).Translate(center, brandLabelY))
	return p
}

func (p *Panel) Port(x, y float64, name string, labelColor svg.Color) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	labelY := port.Top() - padding
	p.Engrave(svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY))
}

func (p *Panel) CvPort(x, y float64) {
	p.Port(x, y, "CV", p.Fg)
}

func (p *Panel) Line(x1, y1, x2, y2 float64) {
	line := svg.Line{
		X1:          x1,
		Y1:          y1,
		X2:          x2,
		Y2:          y2,
		Stroke:      p.Fg,
		StrokeWidth: strokeWidth,
	}
	p.Engrave(line)
}

func (p *Panel) HLine(x1, x2, y float64) {
	p.Line(x1, y, x2, y)
}

func (p *Panel) VLine(x, y1, y2 float64) {
	p.Line(x, y1, x, y2)
}

func (p *Panel) InPort(x, y float64, name string) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	labelY := port.Top() - padding
	label := svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY)
	p.Engrave(boxAround(p.Fg, p.Bg, port, label))
	p.Engrave(label)
}

func (p *Panel) InButtonPort(x, y float64, name string) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	buttonX := x + control.PortRadius + control.ButtonRadius + padding
	button := p.Install(buttonX, y, control.Button(p.Bg, p.Fg))
	labelY := port.Top() - padding
	label := svg.TextAbove(name, svg.SmallFont, p.Fg).Translate(x, labelY)
	p.Engrave(boxAround(p.Fg, p.Bg, port, button, label))
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
	p.Engrave(svg.TextAbove(name, svg.LargeFont, p.Fg).Translate(x, labelY))
}

func (p *Panel) ThumbSwitch(x, y float64, selection int, labels []string) {
	size := len(labels)
	frame := p.Install(x, y, control.ThumbSwitch(size, selection, p.Fg, p.Bg))
	labelY := frame.Bottom() + padding
	p.Engrave(svg.TextBelow(labels[0], svg.SmallFont, p.Fg).Translate(x, labelY))
	if size == 3 {
		labelX := frame.Right() + padding
		p.Engrave(svg.TextRight(labels[1], svg.SmallFont, p.Fg).Translate(labelX, y))
	}
	labelY = frame.Top() - padding
	p.Engrave(svg.TextAbove(labels[size-1], svg.SmallFont, p.Fg).Translate(x, labelY))
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
func (p *Panel) Engrave(s svg.Element) {
	p.Engravings = append(p.Engravings, s)
}

func (p *Panel) ImageSvg() svg.Svg {
	faceplateGroup := svg.GroupOf(p.Engravings...)
	faceplateGroup.Id = "faceplate"
	controlsGroup := svg.GroupOf(p.Frames...)
	controlsGroup.Id = "controls"
	return svg.NewSvg(faceplateGroup, controlsGroup)
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

func boxAround(stroke, fill svg.Color, elements ...svg.Element) svg.Rect {
	const (
		cornerRadius = 1.0
		strokeWidth  = 0.35
	)
	bounds := svg.Bounds(elements...)
	return svg.Rect{
		X:           bounds.Left() - padding,
		Y:           bounds.Top() - padding,
		H:           bounds.Height() + 2.0*padding,
		W:           bounds.Width() + 2.0*padding,
		Fill:        fill,
		Stroke:      stroke,
		StrokeWidth: strokeWidth,
		RX:          cornerRadius,
		RY:          cornerRadius,
	}
}
