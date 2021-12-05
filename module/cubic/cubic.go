package cubic

import (
	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/panel"
	"dhemery.com/panelgen/shape"
)

func Panel() *panel.Panel {
	const (
		hue   = 180
		hp    = 5
		width = hp * shape.MillimetersPerHp
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

	p := panel.New("cubic", "CUBIC", width, fg, bg)

	y := float32(82)
	p.Install(control.Port(fg), left, y)
	p.Install(control.Port(fg), right, y)

	y = y + deltaY
	p.Install(control.Port(fg), left, y)
	p.Install(control.Port(fg), right, y)

	return p
}
