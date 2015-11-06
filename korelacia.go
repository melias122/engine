package psl

import (
	"fmt"
	"math"
)

func Korelacia(n, m int, k0, k1 Kombinacia) float64 {
	if len(k0) != n || len(k1) != n {
		return 0.0
	}

	var kk float64
	for i := 0; i < n; i++ {
		x := (float64(k1[i]) - float64(k0[i])) / float64(m)
		x *= x //^2
		x *= x //^4
		kk += x / float64(n)
	}
	kk = float64(1) - math.Sqrt(kk)
	kk *= kk //^2
	kk *= kk //^4
	kk *= kk //^8
	return kk
}

type filterKorelacia struct {
	n, m     int
	min, max float64
	k0       Kombinacia
}

func NewFilterKorelacia(n, m int, min, max float64, k0 Kombinacia) Filter {
	if min < -1 {
		min = -1
	}
	if max > 1 {
		max = 1
	}
	return filterKorelacia{
		n:   n,
		m:   m,
		min: nextLSS(min),
		max: nextGRT(max),
		k0:  k0,
	}
}

func (k filterKorelacia) String() string {
	return fmt.Sprintf("Kk: %f-%f", k.min, k.max)
}

func (k filterKorelacia) Check(k1 Kombinacia) bool {
	if len(k1) == k.n {
		korelacia := Korelacia(k.n, k.m, k.k0, k1)
		if korelacia < k.min || korelacia > k.max {
			return false
		}
	}
	return true
}

func (k filterKorelacia) CheckSkupina(skupina Skupina) bool {
	return true
}
