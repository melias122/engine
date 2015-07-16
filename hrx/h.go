package hrx

import (
	"math"

	"github.com/melias122/psl/num"
)

type H struct {
	n, m    int
	max     int
	skupiny []int8
	Cisla   num.Nums
}

func New(n, m int) *H {
	h := &H{
		n:       n,
		m:       m,
		skupiny: make([]int8, m),
		Cisla:   make(num.Nums, m),
	}
	h.skupiny[0] = int8(m)
	return h
}

func (h *H) Add(x, y int) {

	if x <= 0 {
		panic("x <= 0")
	}

	// Ak N nie je v vytvorene
	N := h.Cisla[x-1]
	if N == nil {
		N = num.New(x, h.n, h.m)
		h.Cisla[x-1] = N
	}

	// Presun Hrx/HHrx zo skupiny PocetR, do skupiny aktualnej pocetnosti cisla N
	h.Move(1, N.PocetR(), N.PocetR()+1)

	// Incrementuj pocetnost cisla x
	N.Inc(y)
}

func (h *H) Is101() bool {
	return h.Cisla.Is101()
}

// Presun
func (h *H) Move(pocet, from, to int) {
	// keby som odrataval napr z 2->1 (5) je to iste ako 1->2 (-5)
	if from > to {
		from, to = to, from
		pocet = -pocet
	}

	if to >= len(h.skupiny) {
		h.skupiny = append(h.skupiny, make([]int8, h.m)...)
	}
	// priratanie odratanie do danej skupiny
	h.skupiny[from] -= int8(pocet)
	h.skupiny[to] += int8(pocet)

	if to == h.max && h.skupiny[to] == 0 {
		h.max = from
	} else if to > h.max {
		h.max = to
	}
}

// Vrati N
func (h *H) GetN(x int) *num.N {
	if x <= 0 {
		panic("x <= 0")
	}
	return h.Cisla[x-1]
}

// Squaring je asi 10x rychlejsi ako math.Pow ...
func (h *H) value(skupina, pocet, max, m float64) float64 {
	x := (max - skupina) / max
	x *= x // x^2
	x *= x // x^4
	x *= x // x^8
	x *= x // x^16
	return (pocet / m) * x
}

// Hodnota aktualnej zostavy Hrx
func (h *H) Value() float64 {
	if h.max == 0 {
		return 100.0
	}
	var hrx float64
	for skupina, pocet := range h.skupiny {
		if pocet > 0 {
			hrx += h.value(float64(skupina), float64(pocet), float64(h.max), float64(h.m))
		}
	}
	return math.Sqrt(math.Sqrt(hrx)) * 100
}

//Vypocita hodnotu Presun p
func (h *H) ValuePresun(p Presun) float64 {
	// z aktualnej skupiny potrebujem preniest t.Max
	// do dalsej skupiny sk+1
	for _, t := range p {
		h.Move(t.Max, t.Sk, t.Sk+1)
	}
	// Vypocitaj hrx pre zostavu p
	hrx := h.Value()
	// Obnov povodny stav
	for _, t := range p {
		h.Move(t.Max, t.Sk+1, t.Sk)
	}
	return hrx
}

func (h *H) Presun() Presun {
	p := Presun{}
	for sk, max := range h.skupiny {
		if max > 0 {
			p = append(p, Tab{sk, int(max)})
		}
	}
	return p
}
