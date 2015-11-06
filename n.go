package psl

import "strconv"

type Num struct {
	n, m  byte
	cislo byte
	r     ph
	s     []ph
}

func NewNum(x, n, m int) *Num {
	new := &Num{
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

func (n *Num) Cislo() int {
	return int(n.cislo)
}

func (n *Num) Inc(i int) {
	n.r.inc()
	n.s[i].inc()
}

func (n *Num) R() float64 {
	return n.r.Hodnota()
}

func (n *Num) RNext() float64 {
	return n.r.HodnotaNext()
}

func (n *Num) HasSTL(i int) bool {
	if n.Cislo() < int(n.n) {
		return i < n.Cislo()
	}
	if n.Cislo() > int(n.m)-int(n.n)+1 {
		return i+1 >= int(n.m)-n.Cislo()-1
	}
	return true
}

func (n *Num) S(i int) float64 {
	return n.s[i-1].Hodnota()
}

func (n *Num) SNext(i int) float64 {
	return n.s[i-1].HodnotaNext()
}

func (n *Num) PocetR() int {
	return n.r.Pocet()
}

func (n *Num) PocetS(i int) int {
	return n.s[i-1].Pocet()
}

func (n *Num) String() string {
	return strconv.Itoa(n.Cislo())
}
