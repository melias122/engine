package filter

import (
	"fmt"

	"github.com/melias122/engine/engine"
)

func NewFilterSTL1(min, max float64, hhrxNums engine.SSum, n int) Filter {
	return &filterSTL1{newFilterSTL("ƩSTL 1-DO", min, max, hhrxNums, n)}
}

func NewFilterSTL2(min, max float64, hrxNums engine.SSum, n int) Filter {
	return &filterSTL2{newFilterSTL("ƩSTL OD-DO", min, max, hrxNums, n)}
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
	ssum     engine.SSum
	fname    string
}

func newFilterSTL(fname string, min, max float64, ssum engine.SSum, n int) *filterSTL {
	if min <= 0 {
		min = 0.1
	}
	return &filterSTL{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		ssum:  ssum,
		fname: fname,
	}
}

func (s *filterSTL) Check(k engine.Kombinacia) bool {
	sum := s.ssum.S(k) // next
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
