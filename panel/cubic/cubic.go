package cubic

import (
	"fmt"

	"dhemery.com/panelgen/panel"
	"dhemery.com/panelgen/svg"
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
	bg = svg.HslColor(hue, 1, .97)
	fg = svg.HslColor(hue, 1, .3)
)

func Panel() *panel.Panel {
	p := panel.New("CUBIC", hp, fg, bg)
	left := p.Width/4 + 4/3
	right := p.Width - left

	for row := 0; row < 4; row++ {
		y := top + deltaY*float64(row)
		p.CvPort(left, y)
		p.SmallKnob(right, y, coefficientKnobLabel(3-row))
	}

	y := 82.0
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
