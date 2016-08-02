package engine

type Delta int

const (
	POSSITIVE Delta = iota
	NEGATIVE
)

type filterRMinusSTL struct {
	n     int
	nums  Nums
	delta Delta
	fname string

	filterPriority
}

func NewFilterR1MinusSTL1(d Delta, nums Nums, n int) Filter {
	return filterRMinusSTL{
		n:              n,
		nums:           nums,
		delta:          d,
		fname:          "Δ(ƩR1-DO-ƩSTL1-DO)",
		filterPriority: P3,
	}
}

func NewFilterR2MinusSTL2(d Delta, nums Nums, n int) Filter {
	return filterRMinusSTL{
		n:     n,
		nums:  nums,
		delta: d,
		fname: "Δ(ƩROD-DO-ƩSTLOD-DO)",
	}
}

func (r filterRMinusSTL) Check(k Kombinacia) bool {
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

func (r filterRMinusSTL) CheckSkupina(s Skupina) bool {
	return true
}

func (r filterRMinusSTL) String() string {
	var s string
	switch r.delta {
	case POSSITIVE:
		s = "+Δ"
	case NEGATIVE:
		s = "-Δ"
	}
	return r.fname + ": " + s
}
