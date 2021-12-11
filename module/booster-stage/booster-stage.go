package boosterstage

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/panel"
	"dhemery.com/panelgen/shape"
)

const (
	hue = 0
	hp  = 8
)

var (
	bg = shape.HslColor(hue, 1, .97)
	fg = shape.HslColor(hue, 1, .3)
)

func Panel() *panel.Panel {
	p := panel.New("booster-stage", "BOOSTER STAGE", hp, fg, bg)

	left := p.Width()/6 + 1/3
	right := p.Width() - left
	center := p.Width() / 2

	y := 25.0
	dy := 18.5

	p.CvPort(left, y)
	p.LargeKnob(center, y, "LEVEL")
	p.Install(right, y, control.ThumbSwitch(2, 1, p.Fg, p.Bg))

	y += dy
	p.CvPort(left, y)
	p.LargeKnob(center, y, "CURVE")
	p.Install(right, y, control.ThumbSwitch(2, 1, p.Fg, p.Bg))

	y += dy
	p.CvPort(left, y)
	p.LargeKnob(center, y, "DURATION")
	p.Install(right, y, control.ThumbSwitch(3, 2, p.Fg, p.Bg))

	y = 82
	dy = 15

	p.InButtonPort(left, y, "DEFER")
	p.OutButtonPort(right, y, "ACTIVE")

	y += dy
	p.InButtonPort(left, y, "TRIG")
	p.OutButtonPort(right, y, "EOC")

	y += dy
	p.InButtonPort(left, y, "IN")
	p.OutButtonPort(right, y, "OUT")
	return p
}
