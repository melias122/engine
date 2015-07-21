package filter

import (
	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

type stl struct {
	n        int
	min, max float64
	cisla    num.Nums
}

func NewStl(n int, min, max float64, cisla num.Nums) Filter {
	return stl{
		n:     n,
		min:   min,
		max:   max,
		cisla: cisla,
	}
}

func (s stl) Check(k komb.Kombinacia) bool {
	var sum float64
	for i, cislo := range k {
		sum += s.cisla[cislo-1].S(i + 1)
	}
	if len(k) == s.n {
		if sum < s.min || sum > s.max {
			return false
		}
	} else if sum > s.max {
		return false
	}
	return true
}
