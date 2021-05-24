package control

import (
	"image/color"

	"dhemery.com/panelgen/shape"
)

type Port struct {
	BG    color.RGBA
	FG    color.RGBA
	Label shape.Text
}

func (p Port) Slug() string {
	return "port"
}
