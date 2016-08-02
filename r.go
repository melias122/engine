package engine

import (
	"fmt"
)

func NewFilterR1(min, max float64, HHrxCisla Nums, n int) Filter {
	return filterR1{newFilterR("ƩR 1-DO", min, max, HHrxCisla, n)}
}

func NewFilterR2(min, max float64, HrxCisla Nums, n int) Filter {
	return filterR2{newFilterR("ƩR OD-DO", min, max, HrxCisla, n)}
}

type filterR1 struct{ filterR }

func (f filterR1) CheckSkupina(s Skupina) bool {
	return f.checkSkupina(s.R1[0], s.R1[1])
}

type filterR2 struct{ filterR }

func (f filterR2) Check(k Kombinacia) bool { return true }

func (f filterR2) CheckSkupina(s Skupina) bool {
	return f.checkSkupina(s.R2, s.R2)
}

type filterR struct {
	n        int
	min, max float64
	cisla    Nums
	fname    string
}

func newFilterR(fname string, min, max float64, cisla Nums, n int) filterR {
	if min <= 0 {
		min = 0.1
	}
	//TODO: max... asi 1
	return filterR{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		cisla: cisla,
		fname: fname,
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

func (r filterR) checkSkupina(min, max float64) bool {
	return !outOfRangeFloats64(r.min, r.max, min, max)
}

func (r filterR) String() string {
	return fmt.Sprintf("%s: %s-%s", r.fname, ftoa(r.min), ftoa(r.max))
}
