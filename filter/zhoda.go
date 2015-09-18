package filter

import (
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

func NewZhoda(n, min, max int, kombinacia komb.Kombinacia) Filter {
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

type zhoda struct {
	n, min, max int
	kombinacia  komb.Kombinacia
}

func (z zhoda) Check(kombinacia komb.Kombinacia) bool {
	count := komb.Zhoda(z.kombinacia, kombinacia)
	if (len(kombinacia) == z.n && count < z.min) || count > z.max {
		return false
	}
	return true
}

func (z zhoda) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (z zhoda) String() string {
	return fmt.Sprintf("Zh: %d-%d", z.min, z.max)
}
