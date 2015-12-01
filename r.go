package psl

import (
	"fmt"
)

type filterR struct {
	n        int
	min, max float64
	cisla    Nums

	typ int // 1 alebo 2
}

func NewFilterR1(min, max float64, HHrxCisla Nums, n int) Filter {
	return newFilterR(min, max, HHrxCisla, n, 1)
}

func NewFilterR2(min, max float64, HrxCisla Nums, n int) Filter {
	return newFilterR(min, max, HrxCisla, n, 2)
}

func newFilterR(min, max float64, cisla Nums, n int, typ int) Filter {
	if min < 0 {
		min = 0
	}
	//TODO: max... asi 1
	return filterR{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		cisla: cisla,
		typ:   typ,
	}
}

func (r filterR) Check(k Kombinacia) bool {
	var sum float64
	for _, cislo := range k {
		sum += r.cisla[cislo-1].RNext()
	}
	if (len(k) == r.n && sum < r.min) || sum > r.max {
		return false
	}
	return true
}

func (r filterR) CheckSkupina(skupina Skupina) bool {
	if r.typ == 1 {
		// R 1-DO
		if skupina.R1[0] > r.max || skupina.R1[1] < r.min {
			return false
		}
	} else if r.typ == 2 {
		// R OD-DO
		if skupina.R2 > r.max || skupina.R2 < r.min {
			return false
		}
	}
	return true
}

func (r filterR) String() string {
	var fname string
	if r.typ == 1 {
		fname = "ƩR 1-DO"
	} else if r.typ == 2 {
		fname = "ƩR OD-DO"
	}
	return fmt.Sprintf("%s: %f-%f", fname, r.min, r.max)
}
