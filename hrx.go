package psl

import (
	"fmt"
	"math"
)

type H struct {
	n, m   int
	xcisla Presun
	Cisla  Nums
	max    int
}

func NewHrx(n, m int) *H {
	return &H{
		n:      n,
		m:      m,
		xcisla: NewPresun(m),
		Cisla:  make(Nums, m),

		max: 19,
	}
}

func NewHHrx(n, m int) *H {
	return &H{
		n:      n,
		m:      m,
		xcisla: NewPresun(m),
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

func (h *H) Is101() bool {
	return h.Cisla.Is101()
}

func (h *H) GetN(x int) *Num {
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

func (h *H) ValueKombinacia(k Kombinacia) float64 {
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

type hrxfilter struct {
	n        int
	min, max float64
	hrx      *H
	fname    string
}

func NewHrxFilter(n int, min, max float64, h *H, fname string) Filter {
	if min < 0 {
		min = 0
	}
	if max > 100 {
		max = 99.99999999999
	}
	return hrxfilter{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		hrx:   h,
		fname: fname,
	}
}

func (h hrxfilter) Check(k Kombinacia) bool {
	switch h.fname {
	case "HRX":
		return true
	case "HHRX":
		value := h.hrx.ValueKombinacia(k)
		if value < h.min || (len(k) == h.n && value > h.max) {
			return false
		}
	}
	return true
}

func (h hrxfilter) CheckSkupina(skupina Skupina) bool {
	switch h.fname {
	case "HRX":
		if skupina.Hrx > h.max || skupina.Hrx < h.min {
			return false
		}
	case "HHRX":
		if skupina.HHrx[0] > h.max || skupina.HHrx[1] < h.min {
			return false
		}
	}
	return true
}

func (h hrxfilter) String() string {
	return fmt.Sprintf("%s: %f-%f", h.fname, h.min, h.max)
}
