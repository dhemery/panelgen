package control

import "dhemery.com/panelgen/shape"

type mod interface {
	AddFaceplate(s []shape.Bounded)
	AddOverlay(s []shape.Bounded)
	AddControl(path string, svg shape.SVG)
}

type Port struct {
	MetalColor  shape.HSL
	ShadowColor shape.HSL
}

func (p Port) AddTo(m mod, x, y float32) {
	const (
		nutRadius    = 4
		barrelRadius = 3
		holeRadius   = 1.8
		strokeWidth  = .5
	)
	nut := shape.Circle{
		R:           nutRadius + strokeWidth/2,
		Stroke:      &p.ShadowColor,
		StrokeWidth: strokeWidth,
		Fill:        &p.MetalColor,
	}
	barrel := shape.Circle{
		R:           barrelRadius + strokeWidth/2,
		Stroke:      &p.ShadowColor,
		StrokeWidth: strokeWidth,
	}
	hole := shape.Circle{
		R:           holeRadius + strokeWidth/2,
		Stroke:      &p.ShadowColor,
		StrokeWidth: strokeWidth,
		Fill:        &p.ShadowColor,
	}
	shapes := []shape.Bounded{nut, barrel, hole}

	var g shape.G
	g.Add(shapes...)
	var svg shape.SVG
	svg.Add(g)

	m.AddOverlay(shapes)
	m.AddControl("port", svg)
}
