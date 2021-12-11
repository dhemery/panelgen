package panel

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

const (
	Height           = 128.5
	MillimetersPerHp = 5.08
	nameLabelY       = 9
	brandLabelY      = Height - nameLabelY
)

type Hp int

type Panel struct {
	Slug        string
	Engravings  []shape.Bounded
	ImageFrames []shape.Bounded
	Controls    []control.Control
	Fg, Bg      shape.HSL
	Hp          Hp
}

func New(slug, name string, hp Hp, fg, bg shape.HSL) *Panel {
	p := &Panel{
		Slug: slug,
		Fg:   fg,
		Bg:   bg,
		Hp:   hp,
	}

	outlineThickness := float32(0.5)
	faceplateRect := shape.Rect{
		X:           outlineThickness / 2,
		Y:           outlineThickness / 2,
		W:           p.Width() - outlineThickness,
		H:           Height - outlineThickness,
		Fill:        &bg,
		Stroke:      &fg,
		StrokeWidth: shape.StrokeWidth,
	}
	p.Engrave(faceplateRect)
	center := p.Width() / 2

	p.LabelBelow(center, brandLabelY, "DHE", TitleFont)
	p.LabelAbove(center, nameLabelY, name, TitleFont)
	return p
}

func (p *Panel) Width() float32 {
	return float32(p.Hp) * MillimetersPerHp
}

func (p *Panel) LabelAbove(x, y float32, text string, font shape.Font) {
	p.Engrave(LabelAbove(text, font, p.Fg).Translate(x, y))
}

func (p *Panel) LabelBelow(x, y float32, text string, font shape.Font) {
	p.Engrave(LabelBelow(text, font, p.Fg).Translate(x, y))
}

func (p *Panel) Port(x, y float32, name string, labelColor shape.HSL) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	labelY := port.Top() - shape.Padding
	p.LabelAbove(x, labelY, name, SmallFont)
}

func (p *Panel) CvPort(x, y float32) {
	p.Port(x, y, "CV", p.Fg)
}

func (p *Panel) InPort(x, y float32, name string) {
	p.boxedPort(x, y, name, p.Bg, p.Fg)
}

func (p *Panel) InButtonPort(x, y float32, name string) {
	buttonOffset := control.PortRadius + control.ButtonRadius + shape.Padding
	p.buttonPort(x, y, buttonOffset, name, p.Bg, p.Fg)
}

func (p *Panel) OutPort(x, y float32, name string) {
	p.boxedPort(x, y, name, p.Fg, p.Bg)
}

func (p *Panel) OutButtonPort(x, y float32, name string) {
	buttonOffset := -control.PortRadius - control.ButtonRadius - shape.Padding
	p.buttonPort(x, y, buttonOffset, name, p.Fg, p.Bg)
}

func (p *Panel) SmallKnob(x, y float32, name string) {
	knob := p.Install(x, y, control.SmallKnob(p.Fg, p.Bg))
	labelY := knob.Top() - shape.Padding
	p.LabelAbove(x, labelY, name, SmallFont)
}

func (p *Panel) LargeKnob(x, y float32, name string) {
	knob := p.Install(x, y, control.LargeKnob(p.Fg, p.Bg))
	labelY := knob.Top() - shape.Padding
	p.LabelAbove(x, labelY, name, SmallFont)
}

// Install installs the control at the specified position.
// The panel image will show the control's selected frame at that position.
// The module's svg directory will include an svg file for each frame of the control.
func (p *Panel) Install(x, y float32, c control.Control) control.Frame {
	p.Controls = append(p.Controls, c)
	frame := c.DefaultFrame.Translate(x, y)
	p.ImageFrames = append(p.ImageFrames, frame)
	return frame
}

// Engrave engraves the shape into the faceplate at the specified position.
func (p *Panel) Engrave(s shape.Bounded) {
	p.Engravings = append(p.Engravings, s)
}

func (p *Panel) FaceplateSvg() shape.Svg {
	return shape.NewSvg(p.Engravings...)
}

func (p *Panel) ImageSvg() shape.Svg {
	faceplate := shape.NewGroup(p.Engravings...)
	faceplate.Id = "faceplate"
	imageOverlay := shape.NewGroup(p.ImageFrames...)
	imageOverlay.Id = "image"
	return shape.NewSvg(faceplate, imageOverlay)
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

func (p *Panel) boxAround(fill, stroke shape.HSL, elements ...shape.Bounded) {
	bounds := shape.BoundsOf(elements...)
	box := shape.Rect{
		X:           bounds.Left - shape.Padding,
		Y:           bounds.Top - shape.Padding,
		H:           bounds.Height() + 2*shape.Padding,
		W:           bounds.Width() + 2*shape.Padding,
		Fill:        &fill,
		Stroke:      &p.Fg,
		StrokeWidth: shape.StrokeWidth,
		RX:          0.5,
		RY:          0.5,
	}
	p.Engrave(box)
}

func (p *Panel) boxedPort(x, y float32, name string, fill, labelColor shape.HSL) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	labelY := port.Top() - shape.Padding
	label := LabelAbove(name, SmallFont, labelColor).Translate(x, labelY)

	p.boxAround(fill, labelColor, port, label)
	p.Engrave(label)
}

func (p *Panel) buttonPort(x, y float32, buttonOffset float32, name string, fill, labelColor shape.HSL) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	button := p.Install(x+buttonOffset, y, control.Button(fill, labelColor))
	label := LabelAbove(name, SmallFont, labelColor).Translate(x, port.Top()-shape.Padding)

	p.boxAround(fill, labelColor, port, button, label)
	p.Engrave(label)
}
