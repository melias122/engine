package filter

import "github.com/melias122/psl/komb"

func NewZhoda(n, min, max int, kombinacia komb.Kombinacia) Filter {
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

	if len(kombinacia) == z.n {
		if count < z.min || count > z.max {
			return false
		}
	} else {
		if count > z.max {
			return false
		}
	}
	return true
}
