package control

import "dhemery.com/panelgen/shape"

const (
	KnobSizeSmall = 8.4
)

type Knob struct {
	Size  float32
	Label shape.Text
}
