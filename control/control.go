package control

import "dhemery.com/panelgen/shape"

type Frame struct{}

type Control struct {
	Frames    map[string]shape.Bounded
	Selection string
}

func (c Control) SelectedFrame() shape.Bounded {
	return c.Frames[c.Selection]
}
