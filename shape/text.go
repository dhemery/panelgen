package shape

type dominantBaseline int

const (
	VAlignAbove dominantBaseline = iota
	VAlignCenter
	VAlignBelow
)

var dominantBaselines = []string{
	VAlignAbove:  "alphabetic",
	VAlignCenter: "middle",
	VAlignBelow:  "hanging",
}

func (b dominantBaseline) String() string {
	return dominantBaselines[b]
}

type textAnchor int

const (
	HAlignCenter textAnchor = iota
	HAlignRight
	HAlignLeft
)

var textAnchors = []string{
	HAlignCenter: "middle",
	HAlignLeft:   "start",
	HAlignRight:  "end",
}

func (a textAnchor) String() string {
	return textAnchors[a]
}

type Text struct {
	XMLName  string `xml:"text"`
	Color    HSL
	FontSize float32
	VAlign   dominantBaseline
	HAlign   textAnchor
	Content  string
}
