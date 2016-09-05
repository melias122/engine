package filter

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"gitlab.com/melias122/engine"
)

type filterZhoda struct {
	n, min, max int
	kombinacia  engine.Kombinacia
	exact       []bool
}

func NewFilterZhodaRange(min, max int, k engine.Kombinacia, n int) Filter {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return &filterZhoda{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: k,
	}
}

func NewFilterZhodaExactFromString(s string, k engine.Kombinacia, n, m int) (Filter, error) {
	r := strings.NewReader(s)
	p := NewParser(r, n, m)
	ints, err := p.ParseInts()
	if err != nil {
		return nil, err
	}
	return NewFilterZhodaExact(ints, k, n), nil
}

func NewFilterZhodaExact(ints []int, k engine.Kombinacia, n int) Filter {
	sort.Ints(ints)
	min := ints[0]
	max := ints[len(ints)-1]
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	exact := make([]bool, n+1)
	for _, i := range ints {
		if i >= 0 && i <= n {
			exact[i] = true
		}
	}
	return &filterZhoda{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: k,
		exact:      exact,
	}
}

func (f *filterZhoda) Check(k engine.Kombinacia) bool {
	count := engine.Zhoda(f.kombinacia, k)
	if (len(k) == f.n && count < f.min) || count > f.max {
		return false
	}
	if f.exact != nil && len(k) == f.n {
		return f.exact[count]
	}
	return true
}

func (f *filterZhoda) CheckSkupina(s engine.Skupina) bool {
	min := int(s.Zh[0])
	max := int(s.Zh[1])
	if min > f.max || max < f.min {
		return false
	}
	return true
}

func (f *filterZhoda) String() string {
	fname := "Zh"
	if f.exact != nil {
		var s []string
		for i, ok := range f.exact {
			if ok {
				s = append(s, strconv.Itoa(i))
			}
		}
		return fmt.Sprintf("%s: %s", fname, strings.Join(s, ", "))
	}
	return fmt.Sprintf("%s: %d-%d", fname, f.min, f.max)
}
