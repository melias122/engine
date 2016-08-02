package engine

import (
	"fmt"
	"sort"
	"strings"
)

func Zhoda(k0, k1 Kombinacia) int {
	if k0 == nil || k1 == nil {
		return 0
	}
	var zhoda int
	for i, j := 0, 0; i < len(k0) && j < len(k1); {
		if k0[i] == k1[j] {
			zhoda++
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return zhoda
}

// ZhodaPresun urcuje poziciu presunu cisla
// z Kombinacie k0 do Kombinacie k1
type ZhodaPresun struct {
	p [][2]int
}

func NewZhodaPresun(k0, k1 Kombinacia) ZhodaPresun {
	if k0 == nil || k1 == nil {
		return ZhodaPresun{}
	}
	var zp ZhodaPresun
	for i, j := 0, 0; i < len(k0) && j < len(k1); {
		if k0[i] == k1[j] {
			zp.p = append(zp.p, [2]int{i + 1, j + 1})
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return zp
}

func (zp ZhodaPresun) String() string {
	s := make([]string, len(zp.p))
	for i, p := range zp.p {
		s[i] = itoa(int(p[0])) + "|" + itoa(int(p[1]))
	}
	return strings.Join(s, ", ")
}

type filterZhoda struct {
	n, min, max int
	kombinacia  Kombinacia
	exact       []bool
	filterPriority
}

func NewFilterZhodaRange(min, max int, k Kombinacia, n int) Filter {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return filterZhoda{
		n:              n,
		min:            min,
		max:            max,
		kombinacia:     k,
		filterPriority: P1,
	}
}

func NewFilterZhodaExactFromString(s string, k Kombinacia, n, m int) (Filter, error) {
	r := strings.NewReader(s)
	p := NewParser(r, n, m)
	ints, err := p.ParseInts()
	if err != nil {
		return nil, err
	}
	return NewFilterZhodaExact(ints, k, n), nil
}

func NewFilterZhodaExact(ints []int, k Kombinacia, n int) Filter {
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
	return filterZhoda{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: k,
		exact:      exact,
	}
}

func (f filterZhoda) Check(k Kombinacia) bool {
	count := Zhoda(f.kombinacia, k)
	if (len(k) == f.n && count < f.min) || count > f.max {
		return false
	}
	if f.exact != nil && len(k) == f.n {
		return f.exact[count]
	}
	return true
}

func (f filterZhoda) CheckSkupina(s Skupina) bool {
	min := int(s.Zh[0])
	max := int(s.Zh[1])
	if min > f.max || max < f.min {
		return false
	}
	return true
}

func (f filterZhoda) String() string {
	fname := "Zh"
	if f.exact != nil {
		var s []string
		for i, ok := range f.exact {
			if ok {
				s = append(s, itoa(i))
			}
		}
		return fmt.Sprintf("%s: %s", fname, strings.Join(s, ", "))
	}
	return fmt.Sprintf("%s: %d-%d", fname, f.min, f.max)
}
