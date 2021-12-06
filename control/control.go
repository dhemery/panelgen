package control

import "dhemery.com/panelgen/shape"

type Frame interface {
	shape.Bounded
	// At returns a copy of the frame positioned at x, y
	At(x, y float32) Frame
}

type Control struct {
	X, Y          float32
	Frames        map[string]Frame
	selectedFrame Frame
}

// At returns a copy of c positioned at x, y
func (c Control) At(x, y float32) Control {
	c.X = x
	c.Y = y
	return c
}

// SelectedFrame returns the control's selected frame at the control's position
func (c Control) SelectedFrame() Frame {
	return c.selectedFrame.At(c.X, c.Y)
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
