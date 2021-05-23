package module

type Slugger interface {
	Slug() string
}

type Module interface {
	Faceplate() Slugger
	Image() Slugger
	Controls() []Slugger
}

var Modules []Module
