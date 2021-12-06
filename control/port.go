package control

import "dhemery.com/panelgen/shape"

func Port(color shape.HSL) Control {
	const (
		nutRadius       = 4
		barrelRadius    = 3
		holeRadius      = 1.8
		shadowThickness = .2
	)
	nut := shape.Circle{
		R:           nutRadius - shadowThickness/2,
		Stroke:      &color,
		StrokeWidth: shadowThickness,
	}
	barrel := shape.Circle{
		R:           barrelRadius - shadowThickness/2,
		Stroke:      &color,
		StrokeWidth: shadowThickness,
	}
	hole := shape.Circle{
		R:    holeRadius,
		Fill: &color,
	}
	frame := shape.NewGroup(nut, barrel, hole)
	c := Control{
		Selection: "port",
		Frames:    map[string]shape.Bounded{"port": frame},
	}
	return c
}
