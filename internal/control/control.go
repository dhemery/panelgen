package control

import "dhemery.com/panelgen/internal/svg"

type Control struct {
	Frames       map[string]svg.Element
	DefaultFrame svg.Element
}
