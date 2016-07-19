package psl

import (
	"fmt"
	"math"
)

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

// Squaring je asi 10x rychlejsi ako math.Pow ...
func (h *H) value(skupina, pocet, max float64) float64 {
	var (
		m = float64(h.m)
	)
	x := (max - skupina) / max
	x *= x // x^2
	x *= x // x^4
	x *= x // x^8
	x *= x // x^16
	return (pocet / m) * x
}

// Hodnota aktualnej zostavy Hrx
func (h *H) Value() float64 {
	return h.valuePresun(h.xcisla)
}

// var xcislaPool = sync.Pool{
// 	New: func() interface{} {
// 		return make(Xcisla, 0, 30)
// 	},
// }

func (h *H) ValueKombinacia(k Kombinacia) float64 {
	// Get xcisla from pool
	// xcisla := xcislaPool.Get().(Xcisla)
	// make copy
	// xcisla = xcisla[:0]
	// for _, x := range h.xcisla {
	// xcisla = append(xcisla, x)
	// }
	xcisla := h.xcisla.copy()
	// move
	for _, cislo := range k {
		sk := h.GetNum(int(cislo)).PocetR()
		xcisla.move(1, sk, sk+1)
	}
	// compute
	valuePresun := h.valuePresun(xcisla)
	// put xcisla back for later use
	// xcislaPool.Put(xcisla)
	return valuePresun
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
		hrx += h.value(float64(p.Sk), float64(p.Max), max)
	}
	return math.Sqrt(math.Sqrt(hrx)) * 100
}

func (h *H) Xcisla() Xcisla {
	return h.xcisla.copy()
}

func NewFilterHrx(min, max float64, Hrx *H, n int) Filter {
	f := filterHrx{newFilterH("Hrx", min, max, Hrx, n)}
	f.filterPriority = P2
	return f
}

func NewFilterHHrx(min, max float64, HHrx *H, n int) Filter {
	f := filterHHrx{newFilterH("HHrx", min, max, HHrx, n)}
	f.filterPriority = P4
	return f
}

type filterHrx struct{ filterH }

func (f filterHrx) Check(Kombinacia) bool { return true }

func (f filterHrx) CheckSkupina(s Skupina) bool {
	return f.checkSkupina(s.Hrx, s.Hrx)
}

type filterHHrx struct{ filterH }

func (f filterHHrx) CheckSkupina(s Skupina) bool {
	return f.checkSkupina(s.HHrx[0], s.HHrx[1])
}

type filterH struct {
	n        int
	min, max float64
	h        *H
	fname    string
	filterPriority
}

func newFilterH(fname string, min, max float64, h *H, n int) filterH {
	if min <= 0 {
		min = 0.1
	}
	if max > 100 {
		max = 99.99999999999
	}
	return filterH{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		h:     h,
		fname: fname,
	}
}

func (h filterH) Check(k Kombinacia) bool {
	value := h.h.ValueKombinacia(k)
	if len(k) == h.n {
		if value < h.min || value > h.max {
			return false
		}
	}
	return true
}

func (h filterH) checkSkupina(min, max float64) bool {
	return !outOfRangeFloats64(h.min, h.max, min, max)
}

func (h filterH) String() string {
	return fmt.Sprintf("%s: %s-%s", h.fname, ftoa(h.min), ftoa(h.max))
}
