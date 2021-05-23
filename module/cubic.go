package module

type cubic struct{}

func (c cubic) Slug() string {
	return "cubic"
}

func (c cubic) Faceplate() Slugger {
	return c
}

func (c cubic) Image() Slugger {
	return c
}

func (c cubic) Controls() []Slugger {
	return make([]Slugger, 0)
}

func newCubic() cubic {
	return cubic{}
}

func init() {
	Modules = append(Modules, newCubic())
}
