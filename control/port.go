package control

import "dhemery.com/panelgen/shape"

const (
	PortRadius = float32(4)
)

func Port(stroke, fill shape.Color) Control {
	const (
		slug            = "port"
		nutRadius       = PortRadius
		barrelRadius    = 3.0
		holeRadius      = 1.8
		shadowThickness = .2
	)
	nut := shape.Circle{
		R:           nutRadius - shadowThickness/2,
		Stroke:      stroke,
		Fill:        fill,
		StrokeWidth: shadowThickness,
	}
	barrel := shape.Circle{
		R:           barrelRadius - shadowThickness/2,
		Stroke:      stroke,
		Fill:        shape.NoColor,
		StrokeWidth: shadowThickness,
	}
	hole := shape.Circle{
		R:    holeRadius,
		Fill: stroke,
	}
	frame := newGroupFrame(nut, barrel, hole)
	return Control{
		Frames:       map[string]Frame{slug: frame},
		DefaultFrame: frame,
	}
}
