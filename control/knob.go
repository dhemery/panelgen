package control

import "dhemery.com/panelgen/shape"

const (
	hugeKnobDiameter   = 19
	largeKnobDiameter  = 12.7
	mediumKnobDiameter = 10
	smallKnobDiameter  = 8.4
	tintKnobDiameter   = 7
)

func knob(slug string, diameter float32, knobColor, pointerColor shape.HSL) Control {
	radius := diameter / 2
	knob := shape.Circle{
		R:    radius,
		Fill: &knobColor,
	}

	pointerWidth := radius / 8
	pointerLength := radius - pointerWidth
	pointer := shape.Line{
		Stroke:      &pointerColor,
		StrokeWidth: pointerWidth,
		Y2:          -pointerLength,
		Cap:         "round",
	}
	frame := shape.NewGroup(knob, pointer)
	c := Control{
		Frames:    map[string]shape.Bounded{slug: frame},
		Selection: slug,
	}
	return c
}

func SmallKnob(knobColor, pointerColor shape.HSL) Control {
	return knob("small-knob", smallKnobDiameter, knobColor, pointerColor)
}
