package boosterstage

import (
	"dhemery.com/panelgen/panel"
	"dhemery.com/panelgen/shape"
)

const (
	hue = 0
	hp  = 8
)

var (
	bg = shape.HSL{H: hue, S: 1, L: .97}
	fg = shape.HSL{H: hue, S: 1, L: .3}
)

func Panel() *panel.Panel {
	p := panel.New("booster-stage", "BOOSTER STAGE", hp, fg, bg)

	left := p.Width()/6 + 1/3
	right := p.Width() - left
	center := p.Width() / 2

	y := float32(25.0)
	dy := float32(18.5)

	p.CvPort(left, y)
	p.LargeKnob(center, y, "LEVEL")

	y += dy
	p.CvPort(left, y)
	p.LargeKnob(center, y, "CURVE")

	y += dy
	p.CvPort(left, y)
	p.LargeKnob(center, y, "DURATION")

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
