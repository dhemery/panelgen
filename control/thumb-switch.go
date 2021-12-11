package control

import "dhemery.com/panelgen/shape"

func thumbSwitchKnurl(length, thickness float32, stroke shape.Color) shape.Line {
	return shape.Line{
		X1:          -length / 2.0,
		X2:          length / 2.0,
		StrokeWidth: thickness,
		Stroke:      stroke,
		Cap:         "round",
	}
}

func thumbSwitchLever(width, knurlThickness float32, stroke, fill shape.Color) shape.Group {
	knurl := thumbSwitchKnurl(width, knurlThickness, stroke)
	knurls := []shape.Bounded{}
	for i := -2; i <= 2; i++ {
		k := knurl
		k.Y1 = float32(i) * knurlThickness * 2.0
		k.Y2 = k.Y1
		knurls = append(knurls, k)
	}
	return shape.NewGroup(knurls...)
}

func ThumbSwitch2(stroke, fill shape.Color, selection int) Control {
	const (
		size             = 2
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
	leverYOffset := lever.Height() / 2
	for i := 0; i < size; i++ {
		leverY := leverYOffset - float32(i)*lever.Height()
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
	states := []Frame{}
	for i := 0; i < size; i++ {
		states = append(states, newGroupFrame(housing, levers[i]))
	}

	return Control{
		Frames: map[string]Frame{
			"thumb-switch-2-1": states[0],
			"thumb-switch-2-2": states[1],
		},
		DefaultFrame: states[selection-1],
	}
}
