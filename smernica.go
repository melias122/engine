package psl

import (
	"fmt"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
)

func Smernica(n, m int, k Kombinacia) float64 {
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

type smernicaFilter struct {
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
	return smernicaFilter{
		n:   n,
		m:   m,
		min: nextLSS(min),
		max: nextGRT(max),
	}
}

func (s smernicaFilter) Check(k Kombinacia) bool {
	if len(k) == s.n {
		smernica := Smernica(s.n, s.m, k)
		if smernica < s.min || smernica > s.max {
			return false
		}
	}
	return true
}

func (s smernicaFilter) CheckSkupina(skupina Skupina) bool {
	return true
}

func (s smernicaFilter) String() string {
	return fmt.Sprintf("Sm: %f-%f", s.min, s.max)
}
