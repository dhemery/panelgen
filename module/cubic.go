package module

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

func newCubic() *Module {
	const (
		hp    = 5
		width = hp * 5.08
		left  = width/4 + 1/3
	)
	var (
		// bg = shape.HSL{H: 180, S: 1, L: .97}
		fg = shape.HSL{H: 180, S: 1, L: .3}
	)

	m := NewModule("cubic")

	m.AddControl(control.Port(fg))
	return m
}

func init() {
	All = append(All, newCubic())
}
