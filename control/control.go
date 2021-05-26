package control

import "dhemery.com/panelgen/shape"

type Control struct {
	Faceplate shape.G
	Overlay   shape.G
	Frames    map[string]shape.SVG
}
