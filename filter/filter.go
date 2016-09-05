package filter

import (
	"math"
	"strconv"
	"strings"

	"gitlab.com/melias122/engine"
)

type Filter interface {
	String() string
	Check(engine.Kombinacia) bool
	CheckSkupina(engine.Skupina) bool
}

type Filters []Filter

func (f Filters) Check(k engine.Kombinacia) bool {
	for _, filter := range f {
		if !filter.Check(k) {
			return false
		}
	}
	return true
}

func (f Filters) CheckSkupina(s engine.Skupina) bool {
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

func lessFloat64(a, b float64) bool {
	return math.Float64bits(a) < math.Float64bits(b)
}

func greaterFloat64(a, b float64) bool {
	return math.Float64bits(a) > math.Float64bits(b)
}

func outOfRangeFloats64(x1, x2, y1, y2 float64) bool {
	if greaterFloat64(x1, y2) || lessFloat64(x2, y1) {
		return true
	}
	return false
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
