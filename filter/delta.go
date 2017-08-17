package filter

import (
	"github.com/melias122/engine/engine"
	"github.com/melias122/engine/hrx"
)

type Delta int

const (
	POSSITIVE Delta = iota
	NEGATIVE
)

type filterRMinusSTL struct {
	n     int
	sum   engine.RSTLk
	delta Delta
	fname string
}

func NewFilterR1MinusSTL1(d Delta, sum engine.RSTLk, n int) Filter {
	return &filterRMinusSTL{
		n:     n,
		sum:   sum,
		delta: d,
		fname: "Δ(ƩR1-DO-ƩSTL1-DO)",
	}
}

func NewFilterR2MinusSTL2(d Delta, sum engine.RSTLk, n int) Filter {
	return &filterRMinusSTL{
		n:     n,
		sum:   sum,
		delta: d,
		fname: "Δ(ƩROD-DO-ƩSTLOD-DO)",
	}
}

func (r *filterRMinusSTL) Check(k engine.Kombinacia) bool {
	if len(k) == r.n {
		sum1 := r.rssum.R(k)
		sum2 := r.rssum.S(k)
		switch r.delta {
		case POSSITIVE:
			return sum1 >= sum2
		case NEGATIVE:
			return sum1 < sum2
		}
	}
	return true
}

func (r *filterRMinusSTL) CheckSkupina(s hrx.Skupina) bool {
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
