package filter

import (
	"fmt"

	"github.com/melias122/engine/engine"
)

func NewFilterR1(min, max float64, hhrxNums engine.RSum, n int) Filter {
	return &filterR1{newFilterR("ƩR 1-DO", min, max, hhrxNums, n)}
}

func NewFilterR2(min, max float64, hrxNums engine.RSum, n int) Filter {
	return &filterR2{newFilterR("ƩR OD-DO", min, max, hrxNums, n)}
}

type filterR1 struct{ *filterR }

func (f *filterR1) CheckSkupina(s engine.Skupina) bool {
	return f.checkSkupina(s.R1[0], s.R1[1])
}

type filterR2 struct{ *filterR }

func (f *filterR2) Check(k engine.Kombinacia) bool { return true }

func (f *filterR2) CheckSkupina(s engine.Skupina) bool {
	return f.checkSkupina(s.R2, s.R2)
}

type filterR struct {
	n        int
	min, max float64
	rsum     engine.RSum
	fname    string
}

func newFilterR(fname string, min, max float64, rsum engine.RSum, n int) *filterR {
	if min <= 0 {
		min = 0.1
	}
	//TODO: max... asi 1
	return &filterR{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		rsum:  rsum,
		fname: fname,
	}
}

func (r *filterR) Check(k engine.Kombinacia) bool {
	sum := r.rsum.R(k)
	if (len(k) == r.n && sum < r.min) || sum > r.max {
		return false
	}
	return true
}

func (r *filterR) checkSkupina(min, max float64) bool {
	return !outOfRangeFloats64(r.min, r.max, min, max)
}

func (r *filterR) String() string {
	return fmt.Sprintf("%s: %s-%s", r.fname, engine.Ftoa(r.min), engine.Ftoa(r.max))
}
