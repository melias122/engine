package engine

import (
	"fmt"
)

func Smernica(k Kombinacia, n, m int) float64 {
	if len(k) < 2 {
		return .0
	}
	var (
		sm  float64
		nSm float64
		M   = float64(m - 1)
		N   = float64(n - 1)
	)
	for i, n0 := range k[:len(k)-1] {
		for j, n1 := range k[i+1:] {
			sm += (float64(n1-n0) / M) / (float64(j+1) / N)
			nSm++
		}
	}
	return sm / nSm
}

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
	return filterSmernica{
		n:   n,
		m:   m,
		min: nextLSS(min),
		max: nextGRT(max),
	}
}

func (s filterSmernica) Check(k Kombinacia) bool {
	if len(k) == s.n {
		smernica := Smernica(k, s.n, s.m)
		if smernica < s.min || smernica > s.max {
			return false
		}
	}
	return true
}

func (s filterSmernica) CheckSkupina(skupina Skupina) bool {
	return true
}

func (s filterSmernica) String() string {
	return fmt.Sprintf("Sm: %s-%s", ftoa(s.min), ftoa(s.max))
}
