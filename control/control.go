package control

import "dhemery.com/panelgen/shape"

type Frame interface {
	shape.Bounded
	// At returns a copy of the frame translated by dx, dy
	Translate(dx, dy float64) Frame
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
	return groupFrame{Group: shape.NewGroup(contents...)}
}

func (f groupFrame) Translate(x, y float64) Frame {
	g := shape.NewGroup(f.Elements...).Translate(x, y)
	return groupFrame{Group: g}
}

func (f groupFrame) Svg() shape.Svg {
	return shape.NewSvg(f.Group)
}
