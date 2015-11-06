package psl

import (
	"fmt"
)

type filterR struct {
	n        int
	min, max float64
	cisla    Nums
	fname    string
}

func NewFilterR(n int, min, max float64, cisla Nums, fname string) Filter {
	if min < 0 {
		min = 0
	}
	//TODO: max... asi 1
	return filterR{
		n:     n,
		min:   nextLSS(min),
		max:   nextGRT(max),
		cisla: cisla,
		fname: fname,
	}
}

func (r filterR) Check(k Kombinacia) bool {
	var sum float64
	for _, cislo := range k {
		sum += r.cisla[cislo-1].RNext()
	}
	if (len(k) == r.n && sum < r.min) || sum > r.max {
		return false
	}
	return true
}

func (r filterR) CheckSkupina(skupina Skupina) bool {
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

func (r filterR) String() string {
	return fmt.Sprintf("%s: %f-%f", r.fname, r.min, r.max)
}
