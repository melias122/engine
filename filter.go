package psl

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (f Filters) CheckSkupina(skupina Skupina) bool {
	for _, filter := range f {
		if !filter.CheckSkupina(skupina) {
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

func dt(f float64) float64 {
	var (
		s  = strconv.FormatFloat(f, 'f', -1, 64)
		dt = 1.0
	)
	if strings.Contains(s, ".") {
		s = strings.Split(s, ".")[1]
		for i := 0; i < len(s); i++ {
			dt /= 10
		}
	}
	return dt
}

func nextGRT(f float64) float64 {
	return f + dt(f)
}

func nextLSS(f float64) float64 {
	return f - dt(f)
}
