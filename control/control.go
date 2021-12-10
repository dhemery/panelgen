package control

import "dhemery.com/panelgen/shape"

type Frame interface {
	shape.Bounded
	// At returns a copy of the frame positioned at x, y
	At(x, y float32) Frame
	// Svg returns a shape.Svg with the same bounds and content as the frame
	Svg() shape.Svg
}

type Control struct {
	Frames       map[string]Frame
	DefaultFrame Frame
}

type groupFrame struct {
	shape.Group
}

func newGroupFrame(contents ...shape.Bounded) Frame {
	return groupFrame{Group: shape.Group{Content: contents}}
}

func (f groupFrame) At(x, y float32) Frame {
	f.X = x
	f.Y = y
	return f
}

func (f groupFrame) Svg() shape.Svg {
	return shape.NewSvg([]shape.Bounded{f.Group})
}
