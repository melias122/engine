package engine

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Cifrovacky [10]byte

func CifrovackyTeorMax(n, m int) Cifrovacky {
	var c Cifrovacky
	for i := 1; i <= m; i++ {
		c.set(i)
	}
	return c
}

func MakeCifrovacky(k Kombinacia) Cifrovacky {
	var c Cifrovacky
	for _, n := range k {
		c.set(int(n))
	}
	return c
}

func (c *Cifrovacky) set(n int) {
	n = n % 10
	switch n {
	case 0:
		c[9]++
	default:
		c[n-1]++
	}
}

func (c Cifrovacky) Strings() []string {
	s := make([]string, len(c))
	for i, c := range c {
		s[i] = strconv.Itoa(int(c))
	}
	return s
}

func (c Cifrovacky) String() string {
	return strings.Join(c.Strings(), " ")
}

func (k Kombinacia) Cifrovacky() Cifrovacky {
	return MakeCifrovacky(k)
}

type filterCifrovacky struct {
	n int
	c Cifrovacky
}

func NewFilterCifrovacky(c Cifrovacky, n, m int) (Filter, error) {
	var sum int
	for i := range c {
		sum += int(c[i])
	}
	if sum != n {
		return nil, fmt.Errorf("súčet cifrovaciek != %d", n)
	}
	tmax := CifrovackyTeorMax(n, m)
	for i := range c {
		if c[i] > tmax[i] {
			return nil, fmt.Errorf("cifra(%d): %d je viac ako maximum %d", (i+1)%10, c[i], tmax[i])
		}
	}
	return filterCifrovacky{n: n, c: c}, nil
}

func (c filterCifrovacky) Check(k Kombinacia) bool {
	cifrovacky := k.Cifrovacky()
	cmp := bytes.Compare(cifrovacky[:], c.c[:])
	if cmp > 0 || (len(k) == c.n && cmp != 0) {
		return false
	}
	return true
}

func (c filterCifrovacky) CheckSkupina(s Skupina) bool {
	return true
}

func (c filterCifrovacky) String() string {
	return fmt.Sprintf("Cifrovacky: %s", c.c[:])
}
