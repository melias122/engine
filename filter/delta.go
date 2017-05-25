package filter

import "github.com/melias122/engine"

type Delta int

const (
	POSSITIVE Delta = iota
	NEGATIVE
)

type filterRMinusSTL struct {
	n     int
	nums  engine.Nums
	delta Delta
	fname string
}

func NewFilterR1MinusSTL1(d Delta, nums engine.Nums, n int) Filter {
	return &filterRMinusSTL{
		n:     n,
		nums:  nums,
		delta: d,
		fname: "Δ(ƩR1-DO-ƩSTL1-DO)",
	}
}

func NewFilterR2MinusSTL2(d Delta, nums engine.Nums, n int) Filter {
	return &filterRMinusSTL{
		n:     n,
		nums:  nums,
		delta: d,
		fname: "Δ(ƩROD-DO-ƩSTLOD-DO)",
	}
}

func (r *filterRMinusSTL) Check(k engine.Kombinacia) bool {
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

func (r *filterRMinusSTL) CheckSkupina(s engine.Skupina) bool {
	return true
}

func (r *filterRMinusSTL) String() string {
	var s string
	switch r.delta {
	case POSSITIVE:
		s = "+Δ"
	case NEGATIVE:
		s = "-Δ"
	}
	return r.fname + ": " + s
}
