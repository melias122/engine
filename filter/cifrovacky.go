package filter

import (
	"bytes"
	"fmt"

	"gitlab.com/melias122/engine"
)

type cifrovacky struct {
	n int
	c engine.Cifrovacky
}

func NewFilterCifrovacky(c engine.Cifrovacky, n, m int) (Filter, error) {
	var sum int
	for i := range c {
		sum += int(c[i])
	}
	if sum != n {
		return nil, fmt.Errorf("súčet cifrovaciek != %d", n)
	}
	tmax := engine.CifrovackyTeorMax(n, m)
	for i := range c {
		if c[i] > tmax[i] {
			return nil, fmt.Errorf("cifra(%d): %d je viac ako maximum %d", (i+1)%10, c[i], tmax[i])
		}
	}
	return &cifrovacky{n: n, c: c}, nil
}

func (c *cifrovacky) Check(k engine.Kombinacia) bool {
	cifrovacky := k.Cifrovacky()
	cmp := bytes.Compare(cifrovacky[:], c.c[:])
	if cmp > 0 || (len(k) == c.n && cmp != 0) {
		return false
	}
	return true
}

func (c *cifrovacky) CheckSkupina(s engine.Skupina) bool {
	return true
}

func (c *cifrovacky) String() string {
	return fmt.Sprintf("Cifrovacky: %s", c.c[:])
}
