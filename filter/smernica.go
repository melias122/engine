package filter

import (
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type smernica struct {
	n, m     int
	min, max float64
}

func NewSmernica(n, m int, min, max float64) Filter {
	if min < 0 {
		min = 0
	}
	if max > 2 {
		max = 2
	}
	return smernica{
		n:   n,
		m:   m,
		min: nextLSS(min),
		max: nextGRT(max),
	}
}

func (s smernica) Check(k komb.Kombinacia) bool {
	if len(k) == s.n {
		smernica := komb.Smernica(s.n, s.m, k)
		if smernica < s.min || smernica > s.max {
			return false
		}
	}
	return true
}

func (s smernica) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (s smernica) String() string {
	return fmt.Sprintf("Sm: %f-%f", s.min, s.max)
}
