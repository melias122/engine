package psl

import (
	"fmt"
)

type filterSucet struct {
	n        int
	min, max int
	filterPriority
}

func NewFilterSucet(min, max, n int) Filter {
	if min < 0 {
		min = 0
	}
	return filterSucet{
		n:              n,
		min:            min,
		max:            max,
		filterPriority: P2,
	}
}

func (s filterSucet) Check(k Kombinacia) bool {
	var sum int
	for _, cislo := range k {
		sum += int(cislo)
	}
	if (len(k) == s.n && sum < s.min) || sum > s.max {
		return false
	}
	return true
}

func (s filterSucet) CheckSkupina(skupina Skupina) bool {
	if int(skupina.Sucet[0]) > s.max || int(skupina.Sucet[1]) < s.min {
		return false
	} else {
		return true
	}
}

func (s filterSucet) String() string {
	return fmt.Sprintf("Sucet: %d-%d", s.min, s.max)
}
