package shape

import (
	"encoding/xml"
	"fmt"
	"math"
)

type HSL struct {
	H float64
	S float64
	L float64
}

func (hsl HSL) RGBString() string {
	twoL := 2 * hsl.L
	chroma := (1 - math.Abs(twoL-1)) * hsl.S
	hp := hsl.H / 60
	hpMod2 := math.Mod(hp, 2)
	x := chroma * (1 - math.Abs(hpMod2-1))

	var rp, gp, bp float64

	switch int64(hp) {
	case 0:
		rp, gp, bp = chroma, x, 0
	case 1:
		rp, gp, bp = x, chroma, 0
	case 2:
		rp, gp, bp = 0, chroma, x
	case 3:
		rp, gp, bp = 0, x, chroma
	case 4:
		rp, gp, bp = x, 0, chroma
	case 5:
		rp, gp, bp = chroma, 0, x
	default:
		panic(fmt.Sprintf("HSL %#+v out of range", hsl))
	}
	m := hsl.L - chroma/2
	return fmt.Sprintf("#%02x%02x%02x", uint8((rp+m)*255+0.5), uint8((gp+m)*255+0.5), uint8((bp+m)*255+0.5))
}

func (hsl HSL) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: hsl.RGBString(),
	}, nil
}
