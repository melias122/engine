package filter

import "github.com/melias122/psl/komb"

type cislovacky struct {
	n        int
	min, max int
	fun      func(int) bool
}

func NewCislovacky(n, min, max int, fun func(int) bool) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: fun,
	}
}

func (c cislovacky) Check(combination komb.Kombinacia) bool {
	var count int
	for _, number := range combination {
		if c.fun(int(number)) {
			count++
		}
	}
	if len(combination) == c.n {
		if !(count >= c.min && count <= c.max) {
			return false
		}
	} else {
		if count > c.max {
			return false
		}
	}
	return true
}
