package control

import (
	"dhemery.com/panelgen/svg"
)

func Port(stroke, fill svg.Color) Control {
	const (
		slug            = "port"
		nutRadius       = PortRadius
		barrelRadius    = 3.0
		holeRadius      = 1.8
		shadowThickness = .2
	)
	nut := svg.Circle{
		R:           nutRadius - shadowThickness/2,
		Stroke:      stroke,
		Fill:        fill,
		StrokeWidth: shadowThickness,
	}
	barrel := svg.Circle{
		R:           barrelRadius - shadowThickness/2,
		Stroke:      stroke,
		Fill:        svg.NoColor,
		StrokeWidth: shadowThickness,
	}
	hole := svg.Circle{
		R:    holeRadius,
		Fill: stroke,
	}
	frame := svg.GroupOf(nut, barrel, hole)
	return Control{
		Frames:       map[string]svg.Element{slug: frame},
		DefaultFrame: frame,
	}
}

const (
	PortRadius = 4
)
