package control

type Shape interface{}

type Frame struct {
	slug   string
	shapes []Shape
}

func (s *Frame) Add(a []Shape) {
	s.shapes = append(s.shapes, a...)
}
