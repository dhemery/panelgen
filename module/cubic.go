package module

import (
	"fmt"
	"image/color"

	"dhemery.com/panelgen/control"
	"dhemery.com/panelgen/shape"
)

type ctrl interface {
	Slug() string
	Faceplate(f control.Frame, x, y float32)
	Overlay(o control.Frame, x, y float32)
	Frames(controls []control.Frame, x float32, y float32)
}

type module struct {
	faceplate control.Frame
	overlay   control.Frame
	controls  []control.Frame
}

func (m module) Slug() string {
	return "cubic"
}

func (m module) Controls() []Slugger {
	var s []Slugger
	for _, c := range m.controls {
		s = append(s, c)
	}
	return s
}

func (m *module) Add(c ctrl, x, y float32) {
	c.Faceplate(m.faceplate, x, y)
	c.Overlay(m.overlay, x, y)
	c.Frames(m.controls, x, y)
}

func newCubic() module {
	const (
		hp    = 5
		width = hp * 5.08
		left  = width/4 + 1/3
		right = width - left
		dy    = 15
		top   = 20
	)
	var (
		fg = color.RGBA{R: 0, G: 0, B: 0, A: 255}
		bg = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	)

	var m module
	name := shape.Label{
		FontSize: 15,
	}
	name.AddTo(m, 0, 0)
	fpBox := shape.Box{
		Fill:   bg,
		Stroke: fg,
	}
	m.faceplate.Add(name, 0, 0)
	for row := 0; row < 4; row++ {
		label := shape.Label{
			Text: fmt.Sprintf(`X<tspan baseline-shift="super">%d</tspan>`, 3-row),
		}
		knob := control.Knob{
			Size:  control.KnobSizeSmall,
			Label: label,
		}
		cv := control.Port{
			Label: shape.Label{Text: "CV"},
		}
		y := top + dy*float32(row)
		m.Add(cv, left, y)
		m.Add(knob, right, y)
	}
	return m
}

func init() {
	Modules = append(Modules, newCubic())
}
