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
	shapes := []shape.Bounded{nut, barrel, hole}
	c := Control{
		Overlay: shape.G{Content: shapes},
		Frames: map[string]shape.SVG{
			"port": {Content: shapes},
		},
	}
	return c
}
