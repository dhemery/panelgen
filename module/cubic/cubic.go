package cubic

import (
	"fmt"

	"dhemery.com/panelgen/panel"
	"dhemery.com/panelgen/shape"
)

const (
	hue = 180
	hp  = 5
)

const (
	top    = 20
	deltaY = 15
)

var (
	bg = shape.HSL{H: hue, S: 1, L: .97}
	fg = shape.HSL{H: hue, S: 1, L: .3}
)

func Panel() *panel.Panel {
	p := panel.New("cubic", "CUBIC", hp, fg, bg)
	left := p.Width()/4 + 4/3
	right := p.Width() - left

	for row := 0; row < 4; row++ {
		y := top + deltaY*float32(row)
		p.CvPort(left, y)
		p.SmallKnob(right, y, coefficientKnobLabel(3-row))
	}

	y := float32(82)
	p.SmallKnob(left, y, "IN")
	p.SmallKnob(right, y, "OUT")

	y = y + deltaY
	p.CvPort(left, y)  // IN gain
	p.CvPort(right, y) // OUT gain

	y = y + deltaY
	p.InPort(left, y, "IN")
	p.OutPort(right, y, "OUT")

	return p
}

func coefficientKnobLabel(exponent int) string {
	const labelFormat = `X<tspan baseline-shift="super">%d</tspan>`
	return fmt.Sprintf(labelFormat, exponent)
}
