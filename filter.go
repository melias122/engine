package psl

import "fmt"

type Filter interface {
	fmt.Stringer
	Check(Kombinacia) bool
	CheckSkupina(Skupina) bool
}

type Filters []Filter

func (f Filters) Check(k Kombinacia) bool {
	for _, filter := range f {
		if !filter.Check(k) {
			return false
		}
	}
	return true
}

func (f Filters) CheckSkupina(s Skupina) bool {
	for _, filter := range f {
		if !filter.CheckSkupina(s) {
			return false
		}
	}
	return true
}

func (f Filters) String() string {
	var s string
	for _, filter := range f {
		s += filter.String() + "\n"
	}
	return s
}
