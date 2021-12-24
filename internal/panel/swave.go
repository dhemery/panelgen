package panel

import (
	"dhemery.com/panelgen/internal/control"
	"dhemery.com/panelgen/internal/svg"
)

func init() {
	registerBuilder("swave", swave)
}

func swave() *Panel {
	const (
		hue = 16
		hp  = 4
	)
	var (
		fg = svg.HslColor(hue, 1, .5)
		bg = svg.HslColor(hue, 1, .97)
	)

	p := NewPanel("SWAVE", hp, fg, bg, "swave")

	const (
		width     = mmPerHp * hp
		x         = width / 2.0
		shaperTop = 25.0
		shaperDy  = 18.5
		portsTop  = 90.0
		portsDy   = 15.0
	)

	y := shaperTop
	p.LevelRangeSwitch(x, y, 1)
	y += shaperDy
	p.LargeKnob(x, y, "CURVE")
	y += shaperDy
	padding := (width - control.PortDiameter - control.TinyKnobDiameter) / 3.0
	cvX := padding + control.PortRadius
	avX := width - padding - control.TinyKnobRadius
	p.CvPort(cvX, y)
	p.Attenuverter(avX, y)

	y = portsTop
	p.InPort(x, y, "IN")
	y += portsDy
	p.OutPort(x, y, "OUT")

	return p
}
