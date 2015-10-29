package hrx

import (
	"math"

	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

type H struct {
	n, m   int
	xcisla Presun
	Cisla  num.Nums
	max    int
}

func NewHrx(n, m int) *H {
	return &H{
		n:      n,
		m:      m,
		xcisla: NewPresun(m),
		Cisla:  make(num.Nums, m),

		max: 19,
	}
}

func NewHHrx(n, m int) *H {
	return &H{
		n:      n,
		m:      m,
		xcisla: NewPresun(m),
		Cisla:  make(num.Nums, m),
	}
}

func (h *H) Add(x, y int) {

	if x <= 0 || y < 0 {
		panic("hrx.Add: x <= 0")
	}

	// Ak N nie je v vytvorene
	N := h.Cisla[x-1]
	if N == nil {
		N = num.New(x, h.n, h.m)
		h.Cisla[x-1] = N
	}

	// Presun Hrx/HHrx zo skupiny PocetR, do skupiny aktualnej pocetnosti cisla N
	h.xcisla.move(1, N.PocetR(), N.PocetR()+1)

	// Incrementuj pocetnost cisla x
	N.Inc(y)
}

func (h *H) Is101() bool {
	return h.Cisla.Is101()
}

func (h *H) GetN(x int) *num.N {
	N := h.Cisla[x-1]
	if N == nil {
		return num.New(x, h.n, h.m)
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

func (h *H) ValueKombinacia(k komb.Kombinacia) float64 {
	p := h.xcisla.copy() // use chan xcisla to reduce presure on GC
	for _, cislo := range k {
		sk := h.GetN(int(cislo)).PocetR()
		p.move(1, sk, sk+1)
	}
	return h.valuePresun(p)
}

//Vypocita hodnotu Presun p
func (h *H) valuePresun(p Presun) float64 {
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

func (h *H) Presun() Presun {
	return h.xcisla.copy()
}
