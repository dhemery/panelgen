package control

import (
	"encoding/xml"
	"image/color"

	"dhemery.com/panelgen/shape"
)

type Port struct {
	BG    color.RGBA
	FG    color.RGBA
	Label shape.Label
}

func (p Port) Slug() string {
	return "port"
}

func (p Port) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (p Port) MarshalFaceplate(e xml.Encoder) error {
	return nil
}

func (p Port) MarshalOverlay(e xml.Encoder) error {
	return nil
}
