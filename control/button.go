package control

import (
	"dhemery.com/panelgen/svg"
)

const (
	ButtonDiameter = 6
	ButtonRadius   = ButtonDiameter / 2
)

func buttonState(buttonColor, ringColor svg.Color) svg.Circle {
	const (
		thickness = ButtonDiameter / 6
		radius    = ButtonRadius - thickness
	)
	return svg.Circle{
		R:           radius,
		Fill:        buttonColor,
		Stroke:      ringColor,
		StrokeWidth: thickness,
	}
}

func button(slug string, pressedColor, releasedColor svg.Color) Control {
	released := buttonState(releasedColor, releasedColor)
	pressed := buttonState(pressedColor, releasedColor)
	return Control{
		Frames: map[string]svg.Element{
			slug + "-released": released,
			slug + "-pressed":  pressed,
		},
		DefaultFrame: released,
	}
}

func Button(pressedColor, releasedColor svg.Color) Control {
	return button("button", pressedColor, releasedColor)
}

func OutputButton(pressedColor, releasedColor svg.Color) Control {
	return button("output-button", pressedColor, releasedColor)
}
