package filter

import (
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

type stl struct {
	n        int
	min, max float64
	cisla    num.Nums
	fname    string
}

func NewStl(n int, min, max float64, cisla num.Nums, fname string) Filter {
	if min < 0 {
		min = 0
	}
	return stl{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		cisla: cisla,
		fname: fname,
	}
}

func (s stl) Check(k komb.Kombinacia) bool {
	var sum float64
	for i, cislo := range k {
		sum += s.cisla[cislo-1].SNext(i + 1)
	}
	if (len(k) == s.n && sum < s.min) || sum > s.max {
		return false
	}
	return true
}

func (s stl) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (s stl) String() string {
	return fmt.Sprintf("%s: %f-%f", s.fname, s.min, s.max)
}
