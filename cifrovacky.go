package psl

import (
	"bytes"
	"fmt"
	"strconv"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
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

func (k Kombinacia) Cifrovacky() Cifrovacky {
	return MakeCifrovacky(k)
}

type cifrovackyFilter struct {
	n int
	c Cifrovacky
}

func NewCifrovacky(n, m int, c Cifrovacky) (Filter, error) {
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
	return cifrovackyFilter{n: n, c: c}, nil
}

func (c cifrovackyFilter) Check(k Kombinacia) bool {
	cifrovacky := k.Cifrovacky()
	cmp := bytes.Compare(cifrovacky[:], c.c[:])
	if cmp > 0 || (len(k) == c.n && cmp != 0) {
		return false
	}
	return true
}

func (c cifrovackyFilter) CheckSkupina(s Skupina) bool {
	return true
}

func (c cifrovackyFilter) String() string {
	return fmt.Sprintf("Cifrovacky: %s", c.c[:])
}
