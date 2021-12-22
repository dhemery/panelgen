package panel

import (
	"dhemery.com/panelgen/internal/svg"
)

func init() {
	registerBuilder("fuzzy-logic-h", buildFuzzyLogic("FUZZY LOGIC H"))
	registerBuilder("fuzzy-logic-z", buildFuzzyLogic("FUZZY LOGIC Z"))
}

func buildFuzzyLogic(name string) buildFunc {
	return func() *Panel {
		return fuzzyLogic(name)
	}
}

func fuzzyLogic(name string) *Panel {
	const (
		hue = 220.0
		hp  = Hp(9)
	)
	var (
		fg = svg.HslColor(hue, 1, .3)
		bg = svg.HslColor(hue, 1, .97)
	)
	p := NewPanel(name, hp, fg, bg, "fuzzy-logic")

	const (
		top               = 4.0 * mmPerHp
		dy                = 3.0 * mmPerHp
		abPortX           = 1.5 * mmPerHp
		abNegateButtonX   = 3.0 * mmPerHp
		abNegatePortX     = 3.25 * mmPerHp
		acInputY          = top
		bdInputY          = top + 1.0*dy
		levelRangeSwitchY = top + 0.5*dy
	)
	var (
		cdPortX            = p.Width - abPortX
		cdNegateButtonX    = p.Width - abNegateButtonX
		cdNegatePortX      = p.Width - abNegatePortX
		andY               = top + 2.0*dy
		orY                = top + 3.0*dy
		xorY               = top + 4.0*dy
		implicationY       = top + 5.0*dy
		contraImplicationY = top + 6.0*dy
		levelRangeSwitchX  = p.Width / 2.0
	)
	p.LevelRangeSwitch(levelRangeSwitchX, levelRangeSwitchY, 1)
	p.Port(abPortX, acInputY, "A", fg)
	p.Button(abNegateButtonX, acInputY, "¬")
	p.Button(cdNegateButtonX, acInputY, "¬")
	p.Port(cdPortX, acInputY, "C", fg)
	p.Port(abPortX, bdInputY, "B", fg)
	p.Button(abNegateButtonX, bdInputY, "¬")
	p.Button(cdNegateButtonX, bdInputY, "¬")
	p.Port(cdPortX, bdInputY, "D", fg)

	p.OutPort(abPortX, andY, "AND")
	p.OutPort(abNegatePortX, andY, "¬")
	p.OutPort(cdNegatePortX, andY, "¬")
	p.OutPort(cdPortX, andY, "AND")

	p.OutPort(abPortX, orY, "OR")
	p.OutPort(abNegatePortX, orY, "¬")
	p.OutPort(cdNegatePortX, orY, "¬")
	p.OutPort(cdPortX, orY, "OR")

	p.OutPort(abPortX, xorY, "XOR")
	p.OutPort(abNegatePortX, xorY, "¬")
	p.OutPort(cdNegatePortX, xorY, "¬")
	p.OutPort(cdPortX, xorY, "XOR")

	p.OutPort(abPortX, implicationY, "A &#x22B2; B")
	p.OutPort(abNegatePortX, implicationY, "¬")
	p.OutPort(cdNegatePortX, implicationY, "¬")
	p.OutPort(cdPortX, implicationY, "C &#x22B2; D")

	p.OutPort(abPortX, contraImplicationY, "A &#x22B3; B")
	p.OutPort(abNegatePortX, contraImplicationY, "¬")
	p.OutPort(cdNegatePortX, contraImplicationY, "¬")
	p.OutPort(cdPortX, contraImplicationY, "C &#x22B3; D")

	return p
}
