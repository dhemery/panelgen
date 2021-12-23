package panel

import (
	"fmt"

	"dhemery.com/panelgen/internal/control"
	"dhemery.com/panelgen/internal/svg"
)

func init() {
	for _, s := range []scannibal{4, 8, 16} {
		registerBuilder(s.slug(), s.build)
	}
}

const (
	scannibalStepWidthHp         = 2.25
	scannibalStepWidth           = scannibalStepWidthHp * mmPerHp
	scannibalGlobalControlsWidth = padding + control.PortDiameter + padding
	scannibalLabelsWidth         = 10.0
)

type scannibal int

func (s scannibal) slug() string {
	return fmt.Sprintf("scannibal-%d", s)
}
func (s scannibal) hp() Hp {
	const baseHp Hp = 11
	return baseHp + Hp(float64(s)*scannibalStepWidthHp)
}
func (s scannibal) width() float64 {
	return s.hp().toMM()
}
func (s scannibal) margin() float64 {
	excessWidth := s.width() - padding - s.contentWidth()
	return excessWidth / 4.0
}

func (s scannibal) stepBlockWidth() float64 {
	return float64(s) * scannibalStepWidth
}
func (s scannibal) contentWidth() float64 {
	return scannibalGlobalControlsWidth + scannibalLabelsWidth + s.stepBlockWidth() + scannibalGlobalControlsWidth
}

func (s scannibal) build() *Panel {
	const (
		hue = 30
	)
	var (
		hp = s.hp()
		fg = svg.HslColor(hue, 10, .1)
		bg = svg.HslColor(hue, .1, .93)
	)
	p := NewPanel(fmt.Sprintf("SCANNIBAL %d", s), hp, fg, bg, "scannibal")

	const (
		top    = 23.0
		bottom = 117.0
	)
	var (
		portBoxAscent        = padding + svg.SmallFont.FontSize + control.PortRadius + 0.22
		portBoxDescent       = padding + control.PortRadius
		globalControlsTop    = top + portBoxAscent
		globalControlsBottom = bottom - portBoxDescent
		globalControlsDy     = (globalControlsBottom - globalControlsTop) / 4.0
	)
	var (
		margin           = s.margin()
		globalInputsLeft = margin
		globalInputsX    = globalInputsLeft + control.PortRadius + padding
		sequenceLengthY  = globalControlsTop + 0.0*globalControlsDy
		aY               = globalControlsTop + 1.0*globalControlsDy
		bY               = globalControlsTop + 2.0*globalControlsDy
		cY               = globalControlsTop + 3.0*globalControlsDy
		phaseInY         = globalControlsTop + 4.0*globalControlsDy
	)
	p.SmallKnob(globalInputsX, sequenceLengthY, "STEPS")
	p.InPort(globalInputsX, aY, "A")
	p.InPort(globalInputsX, bY, "B")
	p.InPort(globalInputsX, cY, "C")
	p.InPort(globalInputsX, phaseInY, "ϕ")

	const (
		intraSectionGlue = 0.15
		interSectionGlue = 2.0
		stepperWidth     = control.SmallKnobDiameter
	)
	var (
		anchorModes            = []string{"SMPL", "TRACK"}
		startAnchorModeStepper = control.Stepper("anchor-mode", fg, bg, svg.SmallFont, stepperWidth, 1, anchorModes)

		stepperHeight     = startAnchorModeStepper.Height()
		stepperHalfHeight = stepperHeight / 2.0
		knobRadius        = control.SmallKnobDiameter / 2.0

		startAnchorModeY    = top + stepperHalfHeight
		startAnchorSourceY  = startAnchorModeY + stepperHalfHeight + intraSectionGlue + stepperHalfHeight
		startAnchorLevelY   = startAnchorSourceY + stepperHalfHeight + intraSectionGlue + knobRadius
		startAnchorLevelCvY = startAnchorLevelY + knobRadius + intraSectionGlue + control.PortRadius

		endAnchorModeY    = startAnchorLevelCvY + control.PortRadius + interSectionGlue + stepperHalfHeight
		endAnchorSourceY  = endAnchorModeY + stepperHalfHeight + intraSectionGlue + stepperHalfHeight
		endAnchorLevelY   = endAnchorSourceY + stepperHalfHeight + intraSectionGlue + knobRadius
		endAnchorLevelCvY = endAnchorLevelY + knobRadius + intraSectionGlue + control.PortRadius

		shapeY       = endAnchorLevelCvY + control.PortRadius + interSectionGlue + stepperHalfHeight
		curvatureY   = shapeY + stepperHalfHeight + intraSectionGlue + knobRadius
		curvatureCvY = curvatureY + knobRadius + intraSectionGlue + control.PortRadius

		durationY   = curvatureCvY + control.PortRadius + interSectionGlue + knobRadius
		durationCvY = durationY + knobRadius + intraSectionGlue + control.PortRadius

		startAnchorLabelY = (startAnchorSourceY + startAnchorLevelCvY) / 2.0
		endAnchorLabelY   = (endAnchorSourceY + endAnchorLevelCvY) / 2.0
		shapeLabelY       = (shapeY + curvatureCvY) / 2.0
		durationLabelY    = (durationY + durationCvY) / 2.0
	)

	var (
		labelsLeft = globalInputsLeft + scannibalLabelsWidth + margin
		labelsX    = labelsLeft + scannibalLabelsWidth
	)
	p.Engrave(labelsX, startAnchorLabelY, svg.TextLeft("ϕ＝0", svg.LargeFont, fg))
	p.Engrave(labelsX, endAnchorLabelY, svg.TextLeft("ϕ＝1", svg.LargeFont, fg))
	p.Engrave(labelsX, shapeLabelY, svg.TextLeft("SHAPE", svg.LargeFont, fg))
	p.Engrave(labelsX, durationLabelY, svg.TextLeft("[ ϕ ]", svg.LargeFont, fg))

	var (
		endAnchorModeStepper     = control.Stepper("anchor-mode", fg, bg, svg.SmallFont, stepperWidth, 2, anchorModes)
		anchorSources            = []string{"LEVEL", "A", "B", "C", "OUT"}
		startAnchorSourceStepper = control.Stepper("anchor-source", fg, bg, svg.SmallFont, stepperWidth, 5, anchorSources)
		endAnchorSourceStepper   = control.Stepper("anchor-source", fg, bg, svg.SmallFont, stepperWidth, 1, anchorModes)
		shapeNames               = []string{"J", "S"}
		shapeStepper             = control.Stepper("shape", fg, bg, svg.SmallFont, stepperWidth, 1, shapeNames)
		knob                     = control.SmallKnob(fg, bg)
		port                     = control.Port(fg, bg)

		stepBlockLeft  = labelsX + padding
		progressLightY = top - control.LightDiameter*1.5
		stepLabelY     = progressLightY - control.LightDiameter*1.5
	)
	p.VLine(stepBlockLeft, top, bottom)
	for step := 0; step < int(s); step++ {
		left := stepBlockLeft + float64(step)*scannibalStepWidth
		right := left + scannibalStepWidth
		x := left + scannibalStepWidth/2.0
		p.VLine(right, top, bottom)

		p.Light(x, progressLightY)
		p.Engrave(x, stepLabelY, svg.TextAbove(fmt.Sprint(step+1), svg.LargeFont, fg))

		p.Install(x, startAnchorModeY, startAnchorModeStepper)
		p.Install(x, startAnchorSourceY, startAnchorSourceStepper)
		p.Install(x, startAnchorLevelY, knob)
		p.Install(x, startAnchorLevelCvY, port)

		p.Install(x, endAnchorModeY, endAnchorModeStepper)
		p.Install(x, endAnchorSourceY, endAnchorSourceStepper)
		p.Install(x, endAnchorLevelY, knob)
		p.Install(x, endAnchorLevelCvY, port)

		p.Install(x, shapeY, shapeStepper)
		p.Install(x, curvatureY, knob)
		p.Install(x, curvatureCvY, port)

		p.Install(x, durationY, knob)
		p.Install(x, durationCvY, port)
	}

	var (
		globalOutputsX    = stepBlockLeft + s.stepBlockWidth() + margin + control.PortRadius + padding
		levelRangeY       = globalControlsTop + 0.0*globalControlsDy
		stepNumberOutputY = globalControlsTop + 2.0*globalControlsDy
		stepPhaseOutputY  = globalControlsTop + 3.0*globalControlsDy
		outY              = globalControlsTop + 4.0*globalControlsDy
	)
	p.LevelRangeSwitch(globalOutputsX, levelRangeY, 2)
	p.OutPort(globalOutputsX, stepNumberOutputY, "STEP #")
	p.OutPort(globalOutputsX, stepPhaseOutputY, "STEP ϕ")
	p.OutPort(globalOutputsX, outY, "OUT")

	return p
}
