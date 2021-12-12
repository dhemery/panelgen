package control

import (
	"fmt"

	"dhemery.com/panelgen/shape"
)

func Stepper(stepperSlug string, stroke, fill shape.Color, width float64, selection int, stateLabels ...string) Control {
	const (
		padding     = 1
		strokeWidth = 0.25
	)
	var defaultFrame Frame
	frames := map[string]Frame{}

	for i, stateLabel := range stateLabels {
		frameSlug := fmt.Sprint(stepperSlug, "-", i+1)
		label := shape.TextCentered(0, 0, stateLabel, shape.SmallFont, stroke)
		box := shape.Rect{
			X:           -width/2 - padding,
			Y:           -label.Height()/2 - padding,
			H:           label.Height() + 2*padding,
			W:           width + 2*padding,
			StrokeWidth: strokeWidth,
			Stroke:      stroke,
			Fill:        fill,
		}
		frame := newGroupFrame(box, label)
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
