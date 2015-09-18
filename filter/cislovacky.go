package filter

import (
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type cislovacky struct {
	n        int
	min, max int
	fun      func(int) bool
	fname    string
}

func NewCislovacky(n, min, max int, fun func(int) bool, fname string) Filter {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return cislovacky{
		n:     n,
		min:   min,
		max:   max,
		fun:   fun,
		fname: fname,
	}
}

// IsP, IsN, IsPr, IsMc, IsVc, IsC19, IsC0, IscC, IsCc, IsCC
func (c cislovacky) String() string {
	return fmt.Sprintf("%s: %d-%d", c.fname, c.min, c.max)
}

func (c cislovacky) Check(k komb.Kombinacia) bool {
	var count int
	for _, number := range k {
		if c.fun(int(number)) {
			count++
		}
	}

	if count > c.max || (len(k) == c.n && count < c.min) {
		return false
	}
	return true
}

func (c cislovacky) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}
