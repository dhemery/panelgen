package control

import (
	"dhemery.com/panelgen/svg"
)

func TinyKnob(knobColor, pointerColor svg.Color) Control {
	return knob("knob-tiny", tinyKnobDiameter, knobColor, pointerColor)
}

func SmallKnob(knobColor, pointerColor svg.Color) Control {
	return knob("knob-small", smallKnobDiameter, knobColor, pointerColor)
}

func MediumKnob(knobColor, pointerColor svg.Color) Control {
	return knob("knob-medium", mediumKnobDiameter, knobColor, pointerColor)
}

func LargeKnob(knobColor, pointerColor svg.Color) Control {
	return knob("knob-large", largeKnobDiameter, knobColor, pointerColor)
}

const (
	hugeKnobDiameter   = 19
	largeKnobDiameter  = 12.7
	mediumKnobDiameter = 10
	smallKnobDiameter  = 8.4
	tinyKnobDiameter   = 7
)

func knob(slug string, diameter float64, knobColor, pointerColor svg.Color) Control {
	radius := diameter / 2
	knob := svg.Circle{
		R:    radius,
		Fill: knobColor,
	}

	pointerWidth := radius / 8
	pointerLength := radius - pointerWidth
	pointer := svg.Line{
		Stroke:      pointerColor,
		StrokeWidth: pointerWidth,
		Y2:          -pointerLength,
		Cap:         "round",
	}
	frame := svg.GroupOf(knob, pointer)
	return Control{
		Frames:       map[string]svg.Element{slug: frame},
		DefaultFrame: frame,
	}
}
