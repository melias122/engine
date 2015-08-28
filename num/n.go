package num

import "strconv"

type N struct {
	n, m  byte
	cislo byte
	r     ph
	s     []ph
}

func New(x, n, m int) *N {
	new := &N{
		n:     byte(n),
		m:     byte(m),
		cislo: byte(x),
		r:     newph(1, 1, n, m),
		s:     make([]ph, n),
	}
	for i := range new.s {
		new.s[i] = newph(x, i+1, n, m)
	}
	return new
}

func (n *N) Cislo() int {
	return int(n.cislo)
}

func (n *N) Inc(i int) {
	n.r.inc()
	n.s[i].inc()
}

func (n *N) R() float64 {
	return n.r.Hodnota()
}

func (n *N) RNext() float64 {
	return n.r.HodnotaNext()
}

func (n *N) HasSTL(i int) bool {
	if n.Cislo() < int(n.n) {
		return i < n.Cislo()
	}
	if n.Cislo() > int(n.m)-int(n.n)+1 {
		return i+1 >= int(n.m)-n.Cislo()-1
	}
	return true
}

func (n *N) S(i int) float64 {
	return n.s[i-1].Hodnota()
}

func (n *N) SNext(i int) float64 {
	return n.s[i-1].HodnotaNext()
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
