package filter

import (
	"fmt"

	"github.com/melias122/engine/engine"
)

type filterSmernica struct {
	n, m     int
	min, max float64
}

func NewFilterSmernica(min, max float64, n, m int) Filter {
	if min < 0 {
		min = 0
	}
	if max > 2 {
		max = 2
	}
	return &filterSmernica{
		n:   n,
		m:   m,
		min: nextLSS(min),
		max: nextGRT(max),
	}
}

func (s *filterSmernica) Check(k engine.Kombinacia) bool {
	if len(k) == s.n {
		smernica := engine.Smernica(k, s.n, s.m)
		if smernica < s.min || smernica > s.max {
			return false
		}
	}
	return true
}

func (s *filterSmernica) CheckSkupina(skupina engine.Skupina) bool {
	return true
}

func (s *filterSmernica) String() string {
	return fmt.Sprintf("Sm: %s-%s", engine.Ftoa(s.min), engine.Ftoa(s.max))
}
