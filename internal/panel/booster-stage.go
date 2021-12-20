package panel

import (
	"dhemery.com/panelgen/internal/svg"
)

func init() {
	registerBuilder("booster-stage", BoosterStage)
}

func BoosterStage() *Panel {
	const (
		hue = 0
		hp  = 8
	)

	var (
		bg = svg.HslColor(hue, 1.0, .97)
		fg = svg.HslColor(hue, 1.0, .3)
	)

	p := NewPanel("BOOSTER STAGE", hp, fg, bg)
	left := p.Width/6.0 + 1.0/3.0
	right := p.Width - left
	center := p.Width / 2.0

	y := 25.0
	dy := 18.5

	p.HLine(left, right, y)
	p.CvPort(left, y)
	p.LargeKnob(center, y, "LEVEL")
	p.LevelRangeSwitch(right, y, 1)

	y += dy
	p.HLine(left, right, y)
	p.CvPort(left, y)
	p.LargeKnob(center, y, "CURVE")
	p.ShapeSwitch(right, y, 1)

	y += dy
	p.HLine(left, right, y)
	p.CvPort(left, y)
	p.LargeKnob(center, y, "DURATION")
	p.DurationRangeSwitch(right, y, 2)

	y = 82
	dy = 15

	p.InButtonPort(left, y, "DEFER")
	p.OutButtonPort(right, y, "ACTIVE")

	y += dy
	p.InButtonPort(left, y, "TRIG")
	p.OutButtonPort(right, y, "EOC")

	y += dy
	p.InPort(left, y, "IN")
	p.OutPort(right, y, "OUT")
	return p
}
