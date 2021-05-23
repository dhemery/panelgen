package shape

import "image/color"

const (
	vAlignAbove = iota
	vAlignCenter
	vAlignBelow
)
const (
	hAlignCenter = iota
	hAlignRight
	hAlighLeft
)

type Label struct {
	Color    color.RGBA
	FontSize float32
	Text     string
	VAlign   int
	HAlign   int
}

func (l Label) AddTo(m interface{}, i int, i2 int) {

}

