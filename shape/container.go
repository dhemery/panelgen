package shape

type Bounded interface {
	Top() float32
	Right() float32
	Bottom() float32
	Left() float32
	Width() float32
	Height() float32
}

type container struct {
	t, r, b, l float32
	Content    []Bounded
}

func (c *container) Add(shapes ...Bounded) {
	for _, s := range shapes {
		c.Content = append(c.Content, s)
		if s.Top() < c.t {
			c.t = s.Top()
		}
		if s.Right() < c.r {
			c.r = s.Right()
		}
		if s.Bottom() > c.b {
			c.b = s.Bottom()
		}
		if s.Left() < c.l {
			c.l = s.Left()
		}
	}
}

func (c container) Top() float32 {
	return c.t
}

func (c container) Right() float32 {
	return c.r
}

func (c container) Bottom() float32 {
	return c.b
}

func (c container) Left() float32 {
	return c.l
}

func (c container) Width() float32 {
	return c.r - c.l
}

func (c container) Height() float32 {
	return c.b - c.t
}

type SVG struct {
	XMLName string `xml:"svg"`
	container
}

type G struct {
	XMLName string `xml:"g"`
	container
}

func (g G) Translate(x, y float32) G {
	return g
}
