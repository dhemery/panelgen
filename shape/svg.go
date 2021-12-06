package shape

import (
	"encoding/xml"
	"fmt"
)

type SVG struct {
	Content []Bounded
}

func (s SVG) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	b := boundsOf(s.Content)
	version := xml.Attr{Name: xml.Name{Local: "version"}, Value: "1.1"}
	xmlns := xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://www.w3.org/2000/svg"}
	widthMM := fmt.Sprintf("%fmm", b.Width())
	width := xml.Attr{Name: xml.Name{Local: "width"}, Value: widthMM}
	heightMM := fmt.Sprintf("%fmm", b.Height())
	height := xml.Attr{Name: xml.Name{Local: "height"}, Value: heightMM}
	vb := fmt.Sprintf("%f %f %f %f", b.Left, b.Top, b.Width(), b.Height())
	viewBox := xml.Attr{Name: xml.Name{Local: "viewBox"}, Value: vb}

	start.Attr = append(start.Attr, version, xmlns, width, height, viewBox)
	start.Name = xml.Name{Local: "svg"}

	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.Encode(s.Content); err != nil {
		return err
	}
	return e.EncodeToken(start.End())
}
