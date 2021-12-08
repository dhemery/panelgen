package cubic

import (
	"fmt"

	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/panel"
	"dhemery.com/panelgen/shape"
)

func Panel() *panel.Panel {
	const (
		hue   = 180
		hp    = 5
		width = hp * shape.MillimetersPerHp
	)

	const (
		left   = width/4 + 1/3
		right  = width - left
		top    = 20
		deltaY = 15
	)

	var (
		bg = shape.HSL{H: hue, S: 1, L: .97}
		fg = shape.HSL{H: hue, S: 1, L: .3}
	)

	p := panel.New("cubic", "CUBIC", width, fg, bg)
	cvLabel := panel.LabelAbove("CV", panel.SmallFont, fg)

	portControl := control.Port(fg)
	knobControl := control.SmallKnob(fg, bg)

	for row := 0; row < 4; row++ {
		y := top + deltaY*float32(row)
		port := p.Install(left, y, portControl)
		p.Engrave(left, port.Top()-shape.Padding, cvLabel)
		knob := p.Install(right, y, knobControl)
		knobLabelText := fmt.Sprintf(`X<tspan baseline-shift="super">%d</tspan>`, 3-row)
		knobLabel := panel.LabelAbove(knobLabelText, panel.SmallFont, fg)
		p.Engrave(right, knob.Top()-shape.Padding, knobLabel)
	}

	y := float32(82)
	inGainKnob := p.Install(left, y, knobControl)
	inLabel := panel.LabelAbove("IN", panel.SmallFont, fg)
	p.Engrave(left, inGainKnob.Top()-shape.Padding, inLabel)
	outGainKnob := p.Install(right, y, knobControl)
	outLabel := panel.LabelAbove("OUT", panel.SmallFont, fg)
	p.Engrave(right, outGainKnob.Top()-shape.Padding, outLabel)

	y = y + deltaY
	inCvPort := p.Install(left, y, portControl)
	p.Engrave(left, inCvPort.Top()-shape.Padding, cvLabel)
	outCvPort := p.Install(right, y, portControl)
	p.Engrave(right, outCvPort.Top()-shape.Padding, cvLabel)

	y = y + deltaY
	inPort := p.Install(left, y, portControl)
	p.Engrave(left, inPort.Top()-shape.Padding, inLabel)

	outPort := p.Install(right, y, portControl)
	p.Engrave(right, outPort.Top()-shape.Padding, outLabel)

	return p
}
