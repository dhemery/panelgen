package panel

import "dhemery.com/panelgen/internal/svg"

func init() {
	registerBuilder("blossom", Blossom)
}

func Blossom() *Panel {
	const (
		hue = 270
		hp  = 10
	)

	var (
		fg = svg.HslColor(hue, 1.0, 0.5)
		bg = svg.HslColor(hue, 0.66, 0.97)
	)

	p := NewPanel("BLOSSOM", hp, fg, bg, "blossom")

	left := p.Width / 7.0
	right := p.Width - left
	left_center := (right-left)/3.0 + left
	right_center := p.Width - left_center

	y := 25.0
	p.HLine(left, right_center, y)
	p.CvPort(left, y)
	p.Attenuverter(left_center, y)
	p.LargeKnob(right_center, y, "SPEED")

	dy := 18.5
	y += dy
	p.HLine(left, right, y)
	p.CvPort(left, y)
	p.Attenuverter(left_center, y)
	p.LargeKnob(right_center, y, "RATIO")
	p.ThumbSwitch(right, y, 2, "LOCK", "FREE")

	y += dy
	p.HLine(left, right_center, y)
	p.CvPort(left, y)
	p.Attenuverter(left_center, y)
	p.LargeKnob(right_center, y, "DEPTH")

	y += dy
	p.HLine(left, right_center, y)
	p.CvPort(left, y)
	p.Attenuverter(left_center, y)
	p.LargeKnob(right_center, y, "PHASE")

	port_offset := 1.25
	y = 97.0
	p.HLine(left, right, y)
	p.CvPort(left, y)
	p.SmallKnob(left_center, y, "GAIN")
	p.LevelRangeSwitch(right_center, y, 1)
	p.OutPort(right, y+port_offset, "X OUT")

	dy = 15.0
	y += dy
	p.HLine(left, right, y)
	p.CvPort(left, y)
	p.SmallKnob(left_center, y, "GAIN")
	p.LevelRangeSwitch(right_center, y, 1)
	p.OutPort(right, y+port_offset, "Y OUT")

	return p
}
