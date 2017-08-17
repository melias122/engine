package filter

import (
	"fmt"

	"github.com/melias122/engine/engine"
)

func NewFilterHrx(min, max float64, hrx engine.Xk, n int) Filter {
	f := &filterHrx{newFilterH("Hrx", min, max, hrx, n)}
	return f
}

func NewFilterHHrx(min, max float64, hhrx engine.Xk, n int) Filter {
	f := &filterHHrx{newFilterH("HHrx", min, max, hhrx, n)}
	return f
}

type filterHrx struct{ filterH }

func (f *filterHrx) Check(engine.Kombinacia) bool { return true }

func (f *filterHrx) CheckSkupina(s engine.Skupina) bool {
	return f.checkSkupina(s.Hrx, s.Hrx)
}

type filterHHrx struct{ filterH }

func (f *filterHHrx) CheckSkupina(s engine.Skupina) bool {
	return f.checkSkupina(s.HHrx[0], s.HHrx[1])
}

type filterH struct {
	n        int
	min, max float64
	x        engine.Xk
	fname    string
}

func newFilterH(fname string, min, max float64, x engine.Xk, n int) filterH {
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
		x:     x,
		fname: fname,
	}
}

func (h *filterH) Check(k engine.Kombinacia) bool {
	x := h.x.X(k)
	if len(k) == h.n {
		if x < h.min || x > h.max {
			return false
		}
	}
	return true
}

func (h *filterH) checkSkupina(min, max float64) bool {
	return !outOfRangeFloats64(h.min, h.max, min, max)
}

func (h *filterH) String() string {
	return fmt.Sprintf("%s: %s-%s", h.fname, engine.Ftoa(h.min), engine.Ftoa(h.max))
}
