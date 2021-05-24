package module

import (
	"dhemery.com/panelgen/shape"
)

func newCubic() *Module {
	const (
		hp    = 5
		width = hp * 5.08
		left  = width/4 + 1/3
	)

	var faceplate shape.SVG
	var overlay shape.SVG
	m := &Module{
		slug:      "cubic",
		faceplate: &faceplate,
		overlay:   &overlay,
	}

	overlay.Content = append(overlay.Content, shape.Circle{
		CX: 22.3984,
		Fill:   &shape.HSL{H: 180, S: 1, L: .97},
		Stroke: &shape.HSL{H: 180, S: 1, L: .3},
	})
	return m
}

func init() {
	All = append(All, newCubic())
}
