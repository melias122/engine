package filter

import "github.com/melias122/psl/komb"

type smernica struct {
	n, m     int
	min, max float64
}

func NewSmernica(n, m int, min, max float64) Filter {
	return smernica{
		n:   n,
		m:   m,
		min: min,
		max: max,
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
