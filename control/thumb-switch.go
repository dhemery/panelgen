package control

import (
	"fmt"

	"dhemery.com/panelgen/shape"
)

func thumbSwitchKnurl(length, thickness float64, stroke shape.Color) shape.Line {
	return shape.Line{
		X1:          -length / 2,
		X2:          length / 2,
		StrokeWidth: thickness,
		Stroke:      stroke,
		Cap:         "round",
	}
}

func thumbSwitchLever(width, knurlThickness float64, stroke, fill shape.Color) shape.Group {
	knurl := thumbSwitchKnurl(width, knurlThickness, stroke)
	knurls := []shape.Bounded{}
	for i := -2; i <= 2; i++ {
		k := knurl
		k.Y1 = knurlThickness * 2 * float64(i)
		k.Y2 = k.Y1
		knurls = append(knurls, k)
	}
	return shape.NewGroup(knurls...)
}

func ThumbSwitch(size, selection int, stroke, fill shape.Color) Control {
	const (
		width            = 3.0
		housingThickness = width / 8.0
		housingWidth     = width - housingThickness
		cornerRadius     = housingThickness / 2.0
		knurlThickness   = 0.25
		knurlLength      = housingWidth - knurlThickness
		padding          = housingThickness
	)
	lever := thumbSwitchLever(knurlLength, knurlThickness, stroke, fill)
	levers := []shape.Bounded{}
	leverYOffset := lever.Height() * float64(size-1) / 2
	for i := 0; i < size; i++ {
		leverY := leverYOffset - float64(i)*lever.Height()
		levers = append(levers, lever.Translate(0, leverY))
	}
	b := shape.Bounds(levers...)
	housing := shape.Rect{
		X:           b.Left() - housingThickness,
		Y:           b.Top() - housingThickness,
		W:           b.Width() + 2*housingThickness,
		H:           b.Height() + 2*housingThickness,
		Stroke:      stroke,
		StrokeWidth: housingThickness,
		Fill:        fill,
		RX:          cornerRadius,
	}
	var defaultFrame Frame
	frames := map[string]Frame{}
	for i := 0; i < size; i++ {
		slug := fmt.Sprint("thumb-switch-", size, "-", i+1)
		frame := newGroupFrame(housing, levers[i])
		frames[slug] = frame
		if selection == i+1 {
			defaultFrame = frame
		}
	}

	return Control{
		Frames:       frames,
		DefaultFrame: defaultFrame,
	}
}
