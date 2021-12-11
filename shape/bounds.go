package shape

type Bounded interface {
	Top() float32
	Right() float32
	Bottom() float32
	Left() float32
	Width() float32
	Height() float32
}

type bounds struct {
	Top, Right, Bottom, Left float32
}

func (b bounds) Width() float32 {
	return b.Right - b.Left
}

func (b bounds) Height() float32 {
	return b.Bottom - b.Top
}

func BoundsOf(shapes ...Bounded) bounds {
	if len(shapes) < 1 {
		return bounds{}
	}
	first := shapes[0]
	b := bounds{
		Top:    first.Top(),
		Right:  first.Right(),
		Bottom: first.Bottom(),
		Left:   first.Left(),
	}
	for _, s := range shapes[1:] {
		if v := s.Top(); v < b.Top {
			b.Top = v
		}
		if v := s.Right(); v > b.Right {
			b.Right = v
		}
		if v := s.Bottom(); v > b.Bottom {
			b.Bottom = v
		}
		if v := s.Left(); v < b.Left {
			b.Left = v
		}
	}
	return b
}
