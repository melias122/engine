package filter

import (
	Hrx "github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type hrx struct {
	n        int
	min, max float64
	hrx      *Hrx.H
}

func NewHrx(n int, min, max float64, h *Hrx.H) Filter {
	return hrx{
		n:   n,
		min: min,
		max: max,
		hrx: h,
	}
}

func (h hrx) Check(k komb.Kombinacia) bool {
	value := h.hrx.ValueKombinacia(k)
	if len(k) == h.n {
		if value < h.min || value > h.max {
			return false
		}
	} else if value > h.max {
		return false
	}
	return true
}
