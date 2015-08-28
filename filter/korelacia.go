package filter

import (
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type korelacia struct {
	n, m     int
	min, max float64
	k0       komb.Kombinacia
}

func NewKorelacia(n, m int, min, max float64, k0 komb.Kombinacia) Filter {
	if min < -1 {
		min = -1
	}
	if max > 1 {
		max = 1
	}
	return korelacia{
		n:   n,
		m:   m,
		min: nextLSS(min),
		max: nextGRT(max),
		k0:  k0,
	}
}

func (k korelacia) String() string {
	return fmt.Sprintf("Kk: %f-%f", k.min, k.max)
}

func (k korelacia) Check(k1 komb.Kombinacia) bool {
	if len(k1) == k.n {
		korelacia := komb.Korelacia(k.n, k.m, k.k0, k1)
		if korelacia < k.min || korelacia > k.max {
			return false
		}
	}
	return true
}

func (k korelacia) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}
