package filter

import (
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type hrxfilter struct {
	n        int
	min, max float64
	hrx      *hrx.H
	fname    string
}

func NewHrx(n int, min, max float64, h *hrx.H, fname string) Filter {
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

func (h hrxfilter) Check(k komb.Kombinacia) bool {
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

func (h hrxfilter) CheckSkupina(skupina hrx.Skupina) bool {
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
