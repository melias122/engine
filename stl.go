package psl

import (
	"fmt"
)

type filterSTL struct {
	n        int
	min, max float64
	cisla    Nums
	fname    string
}

func NewFilterSTL(n int, min, max float64, cisla Nums, fname string) Filter {
	if min < 0 {
		min = 0
	}
	return filterSTL{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		cisla: cisla,
		fname: fname,
	}
}

func (s filterSTL) Check(k Kombinacia) bool {
	var sum float64
	for i, cislo := range k {
		sum += s.cisla[cislo-1].SNext(i + 1)
	}
	if (len(k) == s.n && sum < s.min) || sum > s.max {
		return false
	}
	return true
}

func (s filterSTL) CheckSkupina(skupina Skupina) bool {
	return true
}

func (s filterSTL) String() string {
	return fmt.Sprintf("%s: %f-%f", s.fname, s.min, s.max)
}
