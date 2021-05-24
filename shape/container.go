package shape

type SVG struct {
	XMLName string `xml:"svg"`
	Content []interface{}
}

type G struct {
	XMLName string `xml:"g"`
	Content []interface{}
}
