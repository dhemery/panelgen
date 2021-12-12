package control

import (
	"fmt"

	"dhemery.com/panelgen/svg"
)

func thumbSwitchKnurl(length, thickness float64, stroke svg.Color) svg.Line {
	return svg.Line{
		X1:          -length / 2,
		X2:          length / 2,
		StrokeWidth: thickness,
		Stroke:      stroke,
		Cap:         "round",
	}
}

func thumbSwitchLever(width, knurlThickness float64, stroke, fill svg.Color) svg.Group {
	knurl := thumbSwitchKnurl(width, knurlThickness, stroke)
	knurls := []svg.Element{}
	for i := -2; i <= 2; i++ {
		k := knurl
		k.Y1 = knurlThickness * 2 * float64(i)
		k.Y2 = k.Y1
		knurls = append(knurls, k)
	}
	return svg.GroupOf(knurls...)
}

func ThumbSwitch(size, selection int, stroke, fill svg.Color) Control {
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
	levers := []svg.Element{}
	leverYOffset := lever.Height() * float64(size-1) / 2
	for i := 0; i < size; i++ {
		leverY := leverYOffset - float64(i)*lever.Height()
		levers = append(levers, lever.Translate(0, leverY))
	}
	b := svg.Bounds(levers...)
	housing := svg.Rect{
		X:           b.Left() - housingThickness,
		Y:           b.Top() - housingThickness,
		W:           b.Width() + 2*housingThickness,
		H:           b.Height() + 2*housingThickness,
		Stroke:      stroke,
		StrokeWidth: housingThickness,
		Fill:        fill,
		RX:          cornerRadius,
	}
	var defaultFrame svg.Element
	frames := map[string]svg.Element{}
	for i := 0; i < size; i++ {
		slug := fmt.Sprint("thumb-switch-", size, "-", i+1)
		frame := svg.GroupOf(housing, levers[i])
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
