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
	top, right, bottom, left float32
}

func (b bounds) Top() float32 {
	return b.top
}
func (b bounds) Right() float32 {
	return b.right
}
func (b bounds) Bottom() float32 {
	return b.bottom
}
func (b bounds) Left() float32 {
	return b.left
}
func (b bounds) Width() float32 {
	return b.Right() - b.Left()
}
func (b bounds) Height() float32 {
	return b.Bottom() - b.Top()
}

func Bounds(shapes ...Bounded) Bounded {
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
