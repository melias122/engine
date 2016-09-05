package filter

import (
	"fmt"

	"gitlab.com/melias122/engine"
)

type filterSucet struct {
	n        int
	min, max int
}

func NewFilterSucet(min, max, n int) Filter {
	if min < 0 {
		min = 0
	}
	return &filterSucet{
		n:   n,
		min: min,
		max: max,
	}
}

func (s *filterSucet) Check(k engine.Kombinacia) bool {
	var sum int
	for _, cislo := range k {
		sum += int(cislo)
	}
	if (len(k) == s.n && sum < s.min) || sum > s.max {
		return false
	}
	return true
}

func (s *filterSucet) CheckSkupina(skupina engine.Skupina) bool {
	if int(skupina.Sucet[0]) > s.max || int(skupina.Sucet[1]) < s.min {
		return false
	} else {
		return true
	}
}

func (s *filterSucet) String() string {
	return fmt.Sprintf("Sucet: %d-%d", s.min, s.max)
}
