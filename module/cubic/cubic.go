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
		port := p.Install(portControl, left, y)
		p.Engrave(cvLabel, port.X, port.Top()-shape.Padding)
		knob := p.Install(knobControl, right, y)
		knobLabelText := fmt.Sprintf(`X<tspan baseline-shift="super">%d</tspan>`, 3-row)
		knobLabel := panel.LabelAbove(knobLabelText, panel.SmallFont, fg)
		p.Engrave(knobLabel, knob.X, knob.Top()-shape.Padding)
	}

	y := float32(82)
	inGainKnob := p.Install(knobControl, left, y)
	inLabel := panel.LabelAbove("IN", panel.SmallFont, fg)
	p.Engrave(inLabel, left, inGainKnob.Top()-shape.Padding)
	outGainKnob := p.Install(knobControl, right, y)
	outLabel := panel.LabelAbove("OUT", panel.SmallFont, fg)
	p.Engrave(outLabel, right, outGainKnob.Top()-shape.Padding)

	y = y + deltaY
	inCvPort := p.Install(portControl, left, y)
	p.Engrave(cvLabel, left, inCvPort.Top()-shape.Padding)
	outCvPort := p.Install(portControl, right, y)
	p.Engrave(cvLabel, right, outCvPort.Top()-shape.Padding)

	y = y + deltaY
	inPort := p.Install(portControl, left, y)
	p.Engrave(inLabel, left, inPort.Top()-shape.Padding)

	outPort := p.Install(portControl, right, y)
	p.Engrave(outLabel, right, outPort.Top()-shape.Padding)

	return p
}
