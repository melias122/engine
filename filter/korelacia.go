package filter

import (
	"fmt"

	"github.com/melias122/engine/engine"
)

type filterKorelacia struct {
	n, m     int
	min, max float64
	k0       engine.Kombinacia
}

func NewFilterKorelacia(min, max float64, k0 engine.Kombinacia, n, m int) Filter {
	if min < -1 {
		min = -1.1
	}
	if max > 2 {
		max = 1.9
	}
	return &filterKorelacia{
		n:   n,
		m:   m,
		min: nextLSS(min),
		max: nextGRT(max),
		k0:  k0,
	}
}

func (k *filterKorelacia) Check(k1 engine.Kombinacia) bool {
	if len(k1) == k.n {
		korelacia := engine.Korelacia(k.k0, k1, k.n, k.m)
		if korelacia < k.min || korelacia > k.max {
			return false
		}
	}
	return true
}

func (k *filterKorelacia) CheckSkupina(skupina engine.Skupina) bool {
	return true
}

func (k *filterKorelacia) String() string {
	return fmt.Sprintf("Kk: %s-%s", engine.Ftoa(k.min), engine.Ftoa(k.max))
}
