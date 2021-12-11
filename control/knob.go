package control

import "dhemery.com/panelgen/shape"

func TinyKnob(knobColor, pointerColor shape.Color) Control {
	return knob("knob-tiny", tinyKnobDiameter, knobColor, pointerColor)
}

func SmallKnob(knobColor, pointerColor shape.Color) Control {
	return knob("knob-small", smallKnobDiameter, knobColor, pointerColor)
}

func MediumKnob(knobColor, pointerColor shape.Color) Control {
	return knob("knob-medium", mediumKnobDiameter, knobColor, pointerColor)
}

func LargeKnob(knobColor, pointerColor shape.Color) Control {
	return knob("knob-large", largeKnobDiameter, knobColor, pointerColor)
}

const (
	hugeKnobDiameter   = 19
	largeKnobDiameter  = 12.7
	mediumKnobDiameter = 10
	smallKnobDiameter  = 8.4
	tinyKnobDiameter   = 7
)

func knob(slug string, diameter float32, knobColor, pointerColor shape.Color) Control {
	radius := diameter / 2
	knob := shape.Circle{
		R:    radius,
		Fill: knobColor,
	}

	pointerWidth := radius / 8
	pointerLength := radius - pointerWidth
	pointer := shape.Line{
		Stroke:      pointerColor,
		StrokeWidth: pointerWidth,
		Y2:          -pointerLength,
		Cap:         "round",
	}
	frame := newGroupFrame(knob, pointer)
	return Control{
		Frames:       map[string]Frame{slug: frame},
		DefaultFrame: frame,
	}
}
