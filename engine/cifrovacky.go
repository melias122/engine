package engine

import (
	"strconv"
	"strings"
)

func CifrovackyTeorMax(n, m int) Cifrovacky {
	var c Cifrovacky
	for i := 1; i <= m; i++ {
		c.set(i)
	}
	return c
}

func NewCifrovacky(k Kombinacia) Cifrovacky {
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
