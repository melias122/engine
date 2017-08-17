package hrx

import (
	"math"

	"github.com/melias122/engine/engine"
)

type cislo struct {
	pocet   int
	hodnota float64
}

type Cislo struct {
	n, m  int
	r     []cislo
	stl   [][]cislo
	cache *hCache

	skupina map[int]int
	max     int
}

func newCislo(cache *hCache, n, m int) *Cislo {
	stl := make([][]cislo, m)
	for i := range stl {
		stl[i] = make([]cislo, n)
	}
	if cache == nil {
		cache = newHCache(n, m)
	}
	return &Cislo{
		n:       n,
		m:       m,
		r:       make([]cislo, m),
		stl:     stl,
		cache:   cache,
		skupina: map[int]int{0: m},
		max:     0,
	}
}

func (c *Cislo) Rp(cislo int) int {
	if cislo < 1 || cislo > c.m {
		panic("Rp: out of bounds")
	}
	return c.r[cislo-1].pocet
}

func (c *Cislo) Rh(cislo int) float64 {
	if cislo < 1 || cislo > c.m {
		panic("Rh: out of bounds")
	}
	return c.r[cislo-1].hodnota
}

func (c *Cislo) STLp(cislo, pozicia int) int {
	if cislo < 1 || cislo > c.m || pozicia < 1 || pozicia > c.n {
		panic("STLp: out of bounds")
	}
	return c.stl[cislo-1][pozicia-1].pocet
}

func (c *Cislo) STLh(cislo, pozicia int) float64 {
	if cislo < 1 || cislo > c.m || pozicia < 1 || pozicia > c.n {
		panic("STLh: out of bounds")
	}
	return c.stl[cislo-1][pozicia-1].hodnota
}

func (x *Cislo) R(k engine.Kombinacia) float64 {
	var sum float64
	for _, c := range k {
		sum += x.Rh(c)
	}
	return sum
}

func (x *Cislo) STL(k engine.Kombinacia) float64 {
	var sum float64
	for i, c := range k {
		sum += x.STLh(c, i+1)
	}
	return sum
}

func (x *Cislo) add(k engine.Kombinacia) {
	for p, c := range k {
		r := &x.r[c-1]
		r.pocet++
		r.hodnota = x.cache.H(c, 1, r.pocet)

		s := &x.stl[c-1][p]
		s.pocet++
		s.hodnota = x.cache.H(c, p+1, s.pocet)

		z, do := r.pocet-1, r.pocet
		x.skupina[z]--
		x.skupina[do]++
		pocet := x.skupina[z]
		if pocet == 0 {
			delete(x.skupina, z)
		}
		if do > x.max {
			x.max = do
		}
	}
}

func (n *Cislo) x(max int, k engine.Kombinacia) float64 {
	if max == 0 && len(k) == 0 {
		return 100
	}
	hrx := .0
	for i, j := range n.skupina {
		for _, c := range k {
			s := n.Rp(c)
			if s == i {
				j--
			} else if s+1 == i {
				j++
			}
		}
		x := float64(max-i) / float64(max)
		x *= x // x^2
		x *= x // x^4
		x *= x // x^8
		x *= x // x^16
		x *= (float64(j) / float64(n.m))

		hrx += x
	}
	return math.Sqrt(math.Sqrt(hrx)) * 100
}
