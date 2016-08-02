package engine

import (
	"sort"
	"strconv"
)

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

func (n *Num) Copy() *Num {
	new := &Num{
		n:     n.n,
		m:     n.m,
		cislo: n.cislo,
		r:     n.r,
		s:     make([]ph, n.n),
	}
	for i := range n.s {
		new.s[i] = n.s[i]
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

type Nums []*Num

func (n Nums) rplus1() Nums {
	nums := make(Nums, n.Len())
	for i, num := range n {
		cp := num.Copy()
		cp.r.inc()
		for j := range cp.s {
			cp.s[j].inc()
		}
		nums[i] = cp
	}
	return nums
}

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
func (n Nums) Sort()              { sort.Sort(n) }

type ByPocetR struct {
	Nums
}

func (by ByPocetR) Less(i, j int) bool {
	return by.Nums[i].PocetR() < by.Nums[j].PocetR()
}
