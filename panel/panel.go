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
	p.Engrave(0, 0, faceplateRect)
	center := p.Width() / 2

	p.LabelBelow(center, brandLabelY, "DHE", TitleFont)
	p.LabelAbove(center, nameLabelY, name, TitleFont)
	return p
}

func (p *Panel) Width() float32 {
	return float32(p.Hp) * MillimetersPerHp
}

func (p *Panel) LabelAbove(x, y float32, text string, font shape.Font) {
	p.Engrave(x, y, LabelAbove(text, font, p.Fg))
}

func (p *Panel) LabelBelow(x, y float32, text string, font shape.Font) {
	p.Engrave(x, y, LabelBelow(text, font, p.Fg))
}

func (p *Panel) Port(x, y float32, name string, labelColor shape.HSL) {
	port := p.Install(x, y, control.Port(p.Fg, p.Bg))
	p.LabelAbove(x, port.Top()-shape.Padding, name, SmallFont)
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
	frame := c.DefaultFrame.At(x, y)
	p.ImageFrames = append(p.ImageFrames, frame)
	return frame
}

// Engrave engraves the shape into the faceplate at the specified position.
func (p *Panel) Engrave(x, y float32, s shape.Bounded) shape.Group {
	g := shape.NewGroup(s).Translate(x, y)
	p.Engravings = append(p.Engravings, g)
	return g
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
		H:           bounds.Height() + 2*shape.Padding,
		W:           bounds.Width() + 2*shape.Padding,
		Fill:        &fill,
		Stroke:      &p.Fg,
		StrokeWidth: shape.StrokeWidth,
		RX:          0.5,
		RY:          0.5,
	}

	boxTop := bounds.Top - shape.Padding
	boxLeft := bounds.Left - shape.Padding
	p.Engrave(boxLeft, boxTop, box)
}

func (p *Panel) boxedPort(x, y float32, name string, fill, labelColor shape.HSL) {
	barePort := control.Port(p.Fg, p.Bg)
	bareLabel := LabelAbove(name, SmallFont, labelColor)

	port := p.Install(x, y, barePort)
	positionedLabel := shape.NewGroup(bareLabel).Translate(x, port.Top()-shape.Padding)
	labelY := port.Top() - shape.Padding

	p.boxAround(fill, labelColor, port, positionedLabel)
	p.Engrave(x, labelY, bareLabel)
}

func (p *Panel) buttonPort(x, y float32, buttonOffset float32, name string, fill, labelColor shape.HSL) {
	barePort := control.Port(p.Fg, p.Bg)
	bareButton := control.Button(fill, labelColor)
	bareLabel := LabelAbove(name, SmallFont, labelColor)

	port := p.Install(x, y, barePort)
	button := p.Install(x+buttonOffset, y, bareButton)
	positionedLabel := shape.NewGroup(bareLabel).Translate(x, port.Top()-shape.Padding)
	labelY := port.Top() - shape.Padding

	p.boxAround(fill, labelColor, port, button, positionedLabel)
	p.Engrave(x, labelY, bareLabel)
}
