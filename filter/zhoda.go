package filter

import (
	"fmt"
	"sort"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

func ZhodaRange(n, min, max int, kombinacia komb.Kombinacia) Filter {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return zhoda{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: kombinacia,
	}
}

func ZhodaExact(n int, ints []int, kombinacia komb.Kombinacia) Filter {
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
	return zhoda{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: kombinacia,
		exact:      exact,
	}
}

type zhoda struct {
	n, min, max int
	kombinacia  komb.Kombinacia
	exact       []bool
}

func (z zhoda) Check(k komb.Kombinacia) bool {
	count := komb.Zhoda(z.kombinacia, k)
	if (len(k) == z.n && count < z.min) || count > z.max {
		return false
	}
	if z.exact != nil && len(k) == z.n {
		return z.exact[count]
	}
	return true
}

func (z zhoda) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (z zhoda) String() string {
	return fmt.Sprintf("Zh: %d-%d", z.min, z.max)
}
