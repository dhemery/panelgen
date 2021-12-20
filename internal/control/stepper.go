package control

import (
	"fmt"

	"dhemery.com/panelgen/internal/svg"
)

func Stepper(stepperSlug string, stroke, fill svg.Color, width float64, selection int, stateLabels []string) Control {
	const (
		padding     = 1
		strokeWidth = 0.25
	)
	var defaultFrame svg.Element
	frames := map[string]svg.Element{}

	for i, stateLabel := range stateLabels {
		frameSlug := fmt.Sprint(stepperSlug, "-", i+1)
		label := svg.TextCentered(stateLabel, svg.SmallFont, stroke)
		box := svg.Rect{
			X:           -width/2 - padding,
			Y:           -label.Height()/2 - padding,
			H:           label.Height() + 2*padding,
			W:           width + 2*padding,
			StrokeWidth: strokeWidth,
			Stroke:      stroke,
			Fill:        fill,
		}
		frame := svg.GroupOf(box, label)
		if selection == i+1 {
			defaultFrame = frame
		}
		frames[frameSlug] = frame
	}

	return Control{
		Frames:       frames,
		DefaultFrame: defaultFrame,
	}
}
