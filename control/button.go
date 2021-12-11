package control

import "dhemery.com/panelgen/shape"

const (
	ButtonDiameter = 6
	ButtonRadius   = ButtonDiameter / 2
)

func buttonState(buttonColor, ringColor shape.HSL) Frame {
	const (
		thickness = ButtonDiameter / 6
		radius    = ButtonRadius - thickness
	)
	c := shape.Circle{
		R:           radius,
		Fill:        &buttonColor,
		Stroke:      &ringColor,
		StrokeWidth: thickness,
	}
	return newGroupFrame(c)
}

func button(slug string, pressedColor, releasedColor shape.HSL) Control {
	released := buttonState(releasedColor, releasedColor)
	pressed := buttonState(pressedColor, releasedColor)
	return Control{
		Frames: map[string]Frame{
			slug + "-released": released,
			slug + "-pressed":  pressed,
		},
		DefaultFrame: released,
	}
}

func Button(pressedColor, releasedColor shape.HSL) Control {
	return button("button", pressedColor, releasedColor)
}

func OutputButton(pressedColor, releasedColor shape.HSL) Control {
	return button("output-button", pressedColor, releasedColor)
}
