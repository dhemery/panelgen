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

	panel := panel.New("cubic", "CUBIC", width, fg, bg)
	cvLabel := shape.LabelAbove("CV", shape.SmallFont, fg)

	portControl := control.Port(fg)
	knobControl := control.SmallKnob(fg, bg)

	for row := 0; row < 4; row++ {
		y := top + deltaY*float32(row)
		port := panel.Install(portControl, left, y)
		panel.Engrave(cvLabel, port.X, port.Top()-shape.Padding)
		knob := panel.Install(knobControl, right, y)
		knobLabelText := fmt.Sprintf(`X<tspan baseline-shift="super">%d</tspan>`, 3-row)
		knobLabel := shape.LabelAbove(knobLabelText, shape.SmallFont, fg)
		panel.Engrave(knobLabel, knob.X, knob.Top()-shape.Padding)
	}

	y := float32(82)
	inGainKnob := panel.Install(knobControl, left, y)
	inLabel := shape.LabelAbove("IN", shape.SmallFont, fg)
	panel.Engrave(inLabel, left, inGainKnob.Top()-shape.Padding)
	outGainKnob := panel.Install(knobControl, right, y)
	outLabel := shape.LabelAbove("OUT", shape.SmallFont, fg)
	panel.Engrave(outLabel, right, outGainKnob.Top()-shape.Padding)

	y = y + deltaY
	inCvPort := panel.Install(portControl, left, y)
	panel.Engrave(cvLabel, left, inCvPort.Top()-shape.Padding)
	outCvPort := panel.Install(portControl, right, y)
	panel.Engrave(cvLabel, right, outCvPort.Top()-shape.Padding)

	y = y + deltaY
	inPort := panel.Install(portControl, left, y)
	panel.Engrave(inLabel, left, inPort.Top()-shape.Padding)

	outPort := panel.Install(portControl, right, y)
	panel.Engrave(outLabel, right, outPort.Top()-shape.Padding)

	return panel
}
