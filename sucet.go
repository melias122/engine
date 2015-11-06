package psl

import (
	"fmt"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
)

type sucetFilter struct {
	n        int
	min, max int
}

func NewSucet(n int, min, max int) Filter {
	if min < 0 {
		min = 0
	}
	return sucetFilter{
		n:   n,
		min: min,
		max: max,
	}
}

func (s sucetFilter) Check(k Kombinacia) bool {
	var sum int
	for _, cislo := range k {
		sum += int(cislo)
	}
	if (len(k) == s.n && sum < s.min) || sum > s.max {
		return false
	}
	return true
}

func (s sucetFilter) CheckSkupina(skupina Skupina) bool {
	if int(skupina.Sucet[0]) > s.max || int(skupina.Sucet[1]) < s.min {
		return false
	} else {
		return true
	}
}

func (s sucetFilter) String() string {
	return fmt.Sprintf("Sucet: %d-%d", s.min, s.max)
}
