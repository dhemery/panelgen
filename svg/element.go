package svg

type Bounded interface {
	Top() float64
	Right() float64
	Bottom() float64
	Left() float64
	Width() float64
	Height() float64
}

type Element interface {
	Bounded
	Translate(dx, dy float64) Element
}

type bounds struct {
	top, right, bottom, left float64
}

func (b bounds) Top() float64 {
	return b.top
}
func (b bounds) Right() float64 {
	return b.right
}
func (b bounds) Bottom() float64 {
	return b.bottom
}
func (b bounds) Left() float64 {
	return b.left
}
func (b bounds) Width() float64 {
	return b.Right() - b.Left()
}
func (b bounds) Height() float64 {
	return b.Bottom() - b.Top()
}

func Bounds(shapes ...Bounded) bounds {
	if len(shapes) < 1 {
		return bounds{}
	}
	first := shapes[0]
	b := bounds{
		top:    first.Top(),
		right:  first.Right(),
		bottom: first.Bottom(),
		left:   first.Left(),
	}
	for _, s := range shapes[1:] {
		if v := s.Top(); v < b.top {
			b.top = v
		}
		if v := s.Right(); v > b.right {
			b.right = v
		}
		if v := s.Bottom(); v > b.bottom {
			b.bottom = v
		}
		if v := s.Left(); v < b.left {
			b.left = v
		}
	}
	return b
}
