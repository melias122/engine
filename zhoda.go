package psl

import (
	"sort"
	"strconv"
	"strings"
)

func Zhoda(k0, k1 Kombinacia) int {
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

// presun urcuje poziciu presunu cisla
// z Kombinacie k0 do Kombinacie k1
type presun [][2]byte

func ZhodaPresun(k0, k1 Kombinacia) presun {
	var p presun
	for i, j := 0, 0; i < len(k0) && j < len(k1); {
		if k0[i] == k1[j] {
			p = append(p, [2]byte{byte(i + 1), byte(j + 1)})
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return p
}

func (p presun) String() string {
	s := make([]string, len(p))
	for i, p := range p {
		s[i] = strconv.Itoa(int(p[0])) + "/" + strconv.Itoa(int(p[1]))
	}
	return strings.Join(s, ", ")
}

type filterZhoda struct {
	n, min, max int
	kombinacia  Kombinacia
	exact       []bool
}

func NewFilterZhodaRange(n, min, max int, kombinacia Kombinacia) Filter {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return filterZhoda{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: kombinacia,
	}
}

func NewFilterZhodaExact(n int, ints []int, kombinacia Kombinacia) Filter {
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
		kombinacia: kombinacia,
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
	return true
}

func (f filterZhoda) String() string {
	return "Zh: " + strconv.Itoa(f.min) + "-" + strconv.Itoa(f.max)
}
