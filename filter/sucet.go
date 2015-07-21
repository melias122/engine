package filter

import "github.com/melias122/psl/komb"

type sucet struct {
	n        int
	min, max int
}

func NewSucet(n int, min, max int) Filter {
	return sucet{
		n:   n,
		min: min,
		max: max,
	}
}

func (s sucet) Check(k komb.Kombinacia) bool {
	var sucetK int
	for _, cislo := range k {
		sucetK += int(cislo)
	}
	if len(k) == s.n {
		if sucetK < s.min || sucetK > s.max {
			return false
		}
	} else if sucetK > s.max {
		return false
	}
	return true
}
