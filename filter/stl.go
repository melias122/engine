package filter

import (
	"fmt"

	"github.com/melias122/engine/engine"
)

func NewFilterSTL1(min, max float64, HHrxNums engine.Nums, n int) Filter {
	return &filterSTL1{newFilterSTL("ƩSTL 1-DO", min, max, HHrxNums, n)}
}

func NewFilterSTL2(min, max float64, HrxNums engine.Nums, n int) Filter {
	return &filterSTL2{newFilterSTL("ƩSTL OD-DO", min, max, HrxNums, n)}
}

type filterSTL1 struct{ *filterSTL }

func (f *filterSTL1) CheckSkupina(s engine.Skupina) bool {
	return f.filterSTL.chceckSkupina(s.S1[0], s.S1[1])
}

type filterSTL2 struct{ *filterSTL }

func (f *filterSTL2) CheckSkupina(s engine.Skupina) bool {
	return f.filterSTL.chceckSkupina(s.S2[0], s.S2[1])
}

type filterSTL struct {
	n        int
	min, max float64
	nums     engine.Nums
	fname    string
}

func newFilterSTL(fname string, min, max float64, nums engine.Nums, n int) *filterSTL {
	if min <= 0 {
		min = 0.1
	}
	return &filterSTL{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		nums:  nums,
		fname: fname,
	}
}

func (s *filterSTL) Check(k engine.Kombinacia) bool {
	var sum float64
	for i, c := range k {
		sum += s.nums[c-1].SNext(i + 1)
	}
	if (len(k) == s.n && sum < s.min) || sum > s.max {
		return false
	}
	return true
}

func (s *filterSTL) chceckSkupina(min, max float64) bool {
	return !outOfRangeFloats64(s.min, s.max, min, max)
}

func (s *filterSTL) String() string {
	return fmt.Sprintf("%s: %s-%s", s.fname, engine.Ftoa(s.min), engine.Ftoa(s.max))
}
