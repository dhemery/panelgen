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
		bg                        = svg.HslColor(hue, 1, .97)
		fg                        = svg.HslColor(hue, 1, .3)
		shapeSwitchLabels         = []string{"J", "S"}
		levelRangeSwitchLabels    = []string{"BI", "UNI"}
		durationRangeSwitchLabels = []string{"1", "10", "100"}
	)

	p := NewPanel("BOOSTER STAGE", hp, fg, bg)
	left := p.Width/6 + 1/3
	right := p.Width - left
	center := p.Width / 2

	y := 25.0
	dy := 18.5

	p.CvPort(left, y)
	p.LargeKnob(center, y, "LEVEL")
	p.ThumbSwitch(right, y, 1, levelRangeSwitchLabels)

	y += dy
	p.CvPort(left, y)
	p.LargeKnob(center, y, "CURVE")
	p.ThumbSwitch(right, y, 1, shapeSwitchLabels)

	y += dy
	p.CvPort(left, y)
	p.LargeKnob(center, y, "DURATION")
	p.ThumbSwitch(right, y, 2, durationRangeSwitchLabels)

	y = 82
	dy = 15

	p.InButtonPort(left, y, "DEFER")
	p.OutButtonPort(right, y, "ACTIVE")

	y += dy
	p.InButtonPort(left, y, "TRIG")
	p.OutButtonPort(right, y, "EOC")

	y += dy
	p.InButtonPort(left, y, "IN")
	p.OutButtonPort(right, y, "OUT")
	return p
}
