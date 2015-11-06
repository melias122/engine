package psl

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
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

func ZhodaRange(n, min, max int, kombinacia Kombinacia) Filter {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return zhodaFilter{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: kombinacia,
	}
}

func ZhodaExact(n int, ints []int, kombinacia Kombinacia) Filter {
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
	return zhodaFilter{
		n:          n,
		min:        min,
		max:        max,
		kombinacia: kombinacia,
		exact:      exact,
	}
}

type zhodaFilter struct {
	n, min, max int
	kombinacia  Kombinacia
	exact       []bool
}

func (z zhodaFilter) Check(k Kombinacia) bool {
	count := Zhoda(z.kombinacia, k)
	if (len(k) == z.n && count < z.min) || count > z.max {
		return false
	}
	if z.exact != nil && len(k) == z.n {
		return z.exact[count]
	}
	return true
}

func (z zhodaFilter) CheckSkupina(skupina Skupina) bool {
	return true
}

func (z zhodaFilter) String() string {
	return fmt.Sprintf("Zh: %d-%d", z.min, z.max)
}
