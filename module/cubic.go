package module

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

func Cubic() *Module {
	const (
		hue   = 180
		hp    = 5
		width = hp * mmPerHp
	)

	const (
		left   = width/4 + 1/3
		right  = width - left
		top    = 20
		deltaY = 15
	)

	var (
		bg = shape.HSL{H: hue, S: 1, L: .97}
		fg = shape.HSL{H: hue, S: 1, L: .3}
	)

	m := NewModule("cubic", hp, fg, bg)

	y := float32(82)
	m.AddControl(control.Port(fg), left, y)
	m.AddControl(control.Port(fg), right, y)

	y = y + deltaY
	m.AddControl(control.Port(fg), left, y)
	m.AddControl(control.Port(fg), right, y)

	return m
}

func init() {
	All = append(All, Cubic())
}
