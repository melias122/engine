package filter

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/melias122/engine/engine"
)

// Cislovacky implementuju Filter pre P, N, Pr, Mc, Vc, C19, C0, cC, Cc, CC
type cislovacky struct {
	n        int
	min, max int
	c        engine.Cislovacka
}

func NewFilterCislovackyRange(min, max int, c engine.Cislovacka, n int) Filter {
	return newFilterCislovackyRange(min, max, c, n)
}

func newFilterCislovackyRange(min, max int, c engine.Cislovacka, n int) *cislovacky {
	if min < 0 {
		min = 0
	}
	if max > n {
		max = n
	}
	return &cislovacky{
		n:   n,
		min: min,
		max: max,
		c:   c,
	}
}

func (c *cislovacky) Check(k engine.Kombinacia) bool {
	_, ok := c.check(k)
	return ok
}

func (c *cislovacky) check(k engine.Kombinacia) (int, bool) {
	var (
		fun   = c.c.Func()
		count int
	)
	for _, n := range k {
		if fun(int(n)) {
			count++
		}
	}
	if count > c.max || (len(k) == c.n && count < c.min) {
		return count, false
	}
	return count, true
}

func (f *cislovacky) CheckSkupina(s engine.Skupina) bool {
	min := int(s.Cislovacky[0][f.c])
	max := int(s.Cislovacky[1][f.c])
	if min > f.max || max < f.min {
		return false
	}
	return true
}

func (c *cislovacky) String() string {
	return fmt.Sprintf("%s: %d-%d", c.c.String(), c.min, c.max)
}

type cislovackyExact struct {
	*cislovacky
	exact []bool
}

func NewFilterCislovackyExactFromString(s string, c engine.Cislovacka, n, m int) (Filter, error) {
	r := strings.NewReader(s)
	p := NewParser(r, n, m)
	ints, err := p.ParseInts()
	if err != nil {
		return nil, err
	}
	return NewFilterCislovackyExact(ints, c, n)
}

func NewFilterCislovackyExact(ints []int, c engine.Cislovacka, n int) (Filter, error) {
	if ints == nil || len(ints) == 0 {
		return nil, errors.New("NewFilterCislovackyExact: aspon 1 cislo musi byt zadane")
	}
	sort.Ints(ints)
	min := ints[0]
	max := ints[len(ints)-1]
	exact := make([]bool, n+1)
	for _, i := range ints {
		if i >= 0 && i <= n {
			exact[i] = true
		}
	}
	return &cislovackyExact{
		cislovacky: newFilterCislovackyRange(min, max, c, n),
		exact:      exact,
	}, nil
}

func (f *cislovackyExact) Check(k engine.Kombinacia) bool {
	count, ok := f.cislovacky.check(k)
	if len(k) < f.n {
		return ok
	}
	return ok && f.exact[count]
}

func (f *cislovackyExact) String() string {
	var s []string
	for i, ok := range f.exact {
		if ok {
			s = append(s, strconv.Itoa(i))
		}
	}
	return fmt.Sprintf("%s: %s", f.c, strings.Join(s, ", "))
}
