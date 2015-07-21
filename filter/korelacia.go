package filter

import "github.com/melias122/psl/komb"

type korelacia struct {
	n, m     int
	min, max float64
	k0       komb.Kombinacia
}

func NewKorelacia(n, m int, min, max float64, k0 komb.Kombinacia) Filter {
	return korelacia{
		n:   n,
		m:   m,
		min: min,
		max: max,
		k0:  k0,
	}
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
