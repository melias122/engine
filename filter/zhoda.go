package filter

import "github.com/melias122/psl/komb"

func NewZhoda(combination []int, n, min, max int) Filter {
	return zhoda{
		n:           n,
		min:         min,
		max:         max,
		combination: combination,
	}
}

type zhoda struct {
	n, min, max int
	combination []int
}

func (z zhoda) Check(combination []int) bool {
	count := komb.Zhoda(z.combination, combination)

	if len(combination) == z.n {
		if !(count >= z.min && count <= z.max) {
			return false
		}
	} else {
		if count > z.max {
			return false
		}
	}
	return true
}
