package filter

import (
	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

type Delta int

const (
	POSSITIVE Delta = iota
	NEGATIVE
)

type rMinusSTL struct {
	n     int
	nums  num.Nums
	delta Delta
	fname string
}

func R1MinusSTL1(d Delta, nums num.Nums, n int) Filter {
	return rMinusSTL{
		n:     n,
		nums:  nums,
		delta: d,
		fname: "Δ(ƩR1-DO-ƩSTL1-DO)",
	}
}

func R2MinusSTL2(d Delta, nums num.Nums, n int) Filter {
	return rMinusSTL{
		n:     n,
		nums:  nums,
		delta: d,
		fname: "Δ(ƩROD-DO-ƩSTLOD-DO)",
	}
}

func (r rMinusSTL) Check(k komb.Kombinacia) bool {
	if len(k) == r.n {
		sum1, sum2 := k.SucetRSNext(r.nums)
		switch r.delta {
		case POSSITIVE:
			return sum1 >= sum2
		case NEGATIVE:
			return sum1 < sum2
		}
	}
	return true
}

func (r rMinusSTL) CheckSkupina(s hrx.Skupina) bool {
	return true
}

func (r rMinusSTL) String() string {
	var s string
	switch r.delta {
	case POSSITIVE:
		s = "+"
	case NEGATIVE:
		s = "-"
	}
	return r.fname + ": " + s
}
