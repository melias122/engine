package filter

import (
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

type r struct {
	n        int
	min, max float64
	cisla    num.Nums
	fname    string
}

func NewR(n int, min, max float64, cisla num.Nums, fname string) Filter {
	if min < 0 {
		min = 0
	}
	//TODO: max... asi 1
	return r{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		cisla: cisla,
		fname: fname,
	}
}

func (r r) Check(k komb.Kombinacia) bool {
	var sum float64
	for _, cislo := range k {
		sum += r.cisla[cislo-1].RNext()
	}
	if (len(k) == r.n && sum < r.min) || sum > r.max {
		return false
	}
	return true
}

func (r r) CheckSkupina(skupina hrx.Skupina) bool {
	switch r.fname {
	case "ƩR 1-DO":
		if skupina.R1[0] > r.max || skupina.R1[1] < r.min {
			return false
		}
	case "ƩR OD-DO":
		if skupina.R2 > r.max || skupina.R2 < r.min {
			return false
		}
	}
	return true
}

func (r r) String() string {
	return fmt.Sprintf("%s: %f-%f", r.fname, r.min, r.max)
}
