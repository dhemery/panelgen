package control

import "dhemery.com/panelgen/shape"

type mod interface {
	AddFaceplate(s []shape.Bounded)
	AddOverlay(s []shape.Bounded)
	AddControl(path string, svg shape.SVG)
}

type Port struct {
	Color shape.HSL
}

func (p Port) AddTo(m mod, x, y float32) {
	const (
		nutRadius       = 4
		barrelRadius    = 3
		holeRadius      = 1.8
		shadowThickness = .2
	)
	nut := shape.Circle{
		R:           nutRadius - shadowThickness/2,
		Stroke:      &p.Color,
		StrokeWidth: shadowThickness,
	}
	barrel := shape.Circle{
		R:           barrelRadius - shadowThickness/2,
		Stroke:      &p.Color,
		StrokeWidth: shadowThickness,
	}
	hole := shape.Circle{
		R:    holeRadius,
		Fill: &p.Color,
	}
	shapes := []shape.Bounded{nut, barrel, hole}

	var svg shape.SVG
	svg.Add(shapes...)

	m.AddOverlay(shapes)
	m.AddControl("port", svg)
}
