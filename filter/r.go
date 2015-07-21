package filter

import (
	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

type r struct {
	n        int
	min, max float64
	cisla    num.Nums
}

func NewR(n int, min, max float64, cisla num.Nums) Filter {
	return r{
		n:     n,
		min:   min,
		max:   max,
		cisla: cisla,
	}
}

func (r r) Check(k komb.Kombinacia) bool {
	var sum float64
	for _, cislo := range k {
		sum += r.cisla[cislo-1].R()
	}
	if len(k) == r.n {
		if sum < r.min || sum > r.max {
			return false
		}
	} else if sum > r.max {
		return false
	}
	return true
}
