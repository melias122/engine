package filter

import (
	"fmt"
	"sort"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

// Cislovacky implementuju Filter pre P, N, Pr, Mc, Vc, C19, C0, cC, Cc, CC
type Cislovacky struct {
	n, min, max int
	fname       string
	f           num.CislovackaFunc
	exact       []bool
}

func CislovackyRange(n, min, max int, c num.Cislovacka) *Cislovacky {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return &Cislovacky{
		n:     n,
		min:   min,
		max:   max,
		f:     c.Func(),
		fname: c.String(),
	}
}

func CislovackyExact(n int, ints []int, c num.Cislovacka) *Cislovacky {
	sort.Ints(ints)
	min := ints[0]
	max := ints[len(ints)-1]
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	exact := make([]bool, n+1)
	for _, i := range ints {
		if i >= 0 && i <= n {
			exact[i] = true
		}
	}
	return &Cislovacky{
		n:     n,
		min:   min,
		max:   max,
		f:     c.Func(),
		fname: c.String(),
		exact: exact,
	}
}

func (c *Cislovacky) String() string {
	return fmt.Sprintf("%s: %d-%d", c.fname, c.min, c.max)
}

func (c *Cislovacky) Check(k komb.Kombinacia) bool {
	var count int
	for _, n := range k {
		if c.f(int(n)) {
			count++
		}
	}
	if count > c.max || (len(k) == c.n && count < c.min) {
		return false
	}
	if c.exact != nil && len(k) == c.n {
		return c.exact[count]
	}
	return true
}

func (c *Cislovacky) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}
