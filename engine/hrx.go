package engine

import "math"

type H struct {
	n, m   int
	xcisla Xcisla
	Cisla  Nums
	max    int
}

func NewHrx(n, m int) *H {
	return &H{
		n:      n,
		m:      m,
		xcisla: NewXcisla(m),
		Cisla:  make(Nums, m),

		max: 19,
	}
}

func NewHHrx(n, m int) *H {
	return &H{
		n:      n,
		m:      m,
		xcisla: NewXcisla(m),
		Cisla:  make(Nums, m),
	}
}

func (h *H) Add(x, y int) {

	if x <= 0 || y < 0 {
		panic("hrx.Add: x <= 0")
	}

	// Ak N nie je v vytvorene
	N := h.Cisla[x-1]
	if N == nil {
		N = NewNum(x, h.n, h.m)
		h.Cisla[x-1] = N
	}

	// Presun Hrx/HHrx zo skupiny PocetR, do skupiny aktualnej pocetnosti cisla N
	h.xcisla.move(1, N.PocetR(), N.PocetR()+1)

	// Incrementuj pocetnost cisla x
	N.Inc(y)
}

func (h *H) GetNum(x int) *Num {
	N := h.Cisla[x-1]
	if N == nil {
		return NewNum(x, h.n, h.m)
	} else {
		return h.Cisla[x-1]
	}
}

func (h *H) Value(k Kombinacia) float64 {
	if k == nil {
		return h.valuePresun(h.xcisla)
	}
	xcisla := h.xcisla.copy()
	// move
	for _, cislo := range k {
		sk := h.GetNum(int(cislo)).PocetR()
		xcisla.move(1, sk, sk+1)
	}
	// compute
	return h.valuePresun(xcisla)
}

//Vypocita hodnotu Presun p
func (h *H) valuePresun(p Xcisla) float64 {
	if p.Max() == 0 {
		return 100
	}
	var (
		max float64
		hrx float64
	)
	if h.max == 0 {
		max = float64(p.Max())
	} else {
		max = float64(h.max)
	}
	for _, p := range p {

		x := (max - float64(p.Sk)) / max
		x *= x // x^2
		x *= x // x^4
		x *= x // x^8
		x *= x // x^16
		x *= (float64(p.Max) / float64(h.m))

		hrx += x
	}
	return math.Sqrt(math.Sqrt(hrx)) * 100
}
