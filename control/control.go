package control

import "dhemery.com/panelgen/svg"

type Control struct {
	Frames       map[string]svg.Element
	DefaultFrame svg.Element
}
