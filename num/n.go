package num

import "strconv"

type N struct {
	n byte
	m byte
	c C
	x uint32
	r ph
	s []ph
}

func Zero(n, m int) *N {
	zero := &N{
		n: byte(n),
		m: byte(m),
		s: make([]ph, n),
	}
	return zero
}

func New(x, n, m int) *N {
	new := &N{
		n: byte(n),
		m: byte(m),
		c: NewC(x),
		x: uint32(x),
		r: newph(1, 1, n, m),
		s: make([]ph, n),
	}
	for i := range new.s {
		new.s[i] = newph(x, i+1, n, m)
	}
	return new
}

func (old *N) MakeCopy() *N {
	new := &N{
		n: old.n,
		m: old.m,
		c: old.c,
		x: old.x,
		r: old.r,
		s: make([]ph, int(old.n)),
	}
	for i := range old.s {
		new.s[i] = old.s[i]
	}
	return new

}

func (n *N) C() C {
	return n.c
}

func (n *N) Cislo() int {
	return int(n.x)
}

func (n *N) Plus(m *N) {
	n.c.Plus(m.c)
	n.x += m.x
	n.r.plus(m.r)
	for i := range n.s {
		n.s[i].plus(m.s[i])
	}
}

func (n *N) Minus(m *N) {
	n.c.Minus(m.c)
	n.x -= m.x
	n.r.minus(m.r)
	for i := range n.s {
		n.s[i].minus(m.s[i])
	}
}

func (n *N) Inc(i int) {
	n.r.inc()
	n.s[i].inc()
}

func (n *N) R() float64 {
	return n.r.Hodnota()
}

func (n *N) S(i int) float64 {
	return n.s[i-1].Hodnota()
}

func (n *N) PocetR() int {
	return n.r.Pocet()
}

func (n *N) PocetS(i int) int {
	return n.s[i-1].Pocet()
}

func (n *N) String() string {
	return strconv.Itoa(n.Cislo())
}

type Nums []*N

func (n Nums) Is101() bool {
	for _, N := range n {
		if N == nil {
			return false
		}
	}
	return true
}

func (c Nums) Len() int           { return len(c) }
func (c Nums) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Nums) Less(i, j int) bool { return c[i].Cislo() < c[j].Cislo() }

type ByPocetR struct {
	Nums
}

func (by ByPocetR) Less(i, j int) bool { return by.Nums[i].PocetR() < by.Nums[j].PocetR() }
