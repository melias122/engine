package filter

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type cifrovacky struct {
	n int
	c komb.Cifrovacky
}

func NewCifrovacky(n, m int, c komb.Cifrovacky) (Filter, error) {
	var sum int
	for i := range c {
		sum += int(c[i])
	}
	if sum != n {
		return nil, fmt.Errorf("súčet cifrovaciek != %d", n)
	}
	tmax := komb.CifrovackyTeorMax(n, m)
	for i := range c {
		if c[i] > tmax[i] {
			return nil, fmt.Errorf("cifra(%d): %d je viac ako maximum %d", (i+1)%10, c[i], tmax[i])
		}
	}
	return cifrovacky{n: n, c: c}, nil
}

func (c cifrovacky) Check(k komb.Kombinacia) bool {
	cifrovacky := k.Cifrovacky()
	cmp := bytes.Compare(cifrovacky[:], c.c[:])
	if cmp > 0 || (len(k) == c.n && cmp != 0) {
		return false
	}
	return true
}

func (c cifrovacky) CheckSkupina(s hrx.Skupina) bool {
	return true
}

func (c cifrovacky) String() string {
	return fmt.Sprintf("Cifrovacky: %s", c.c[:])
}

func ParseCifrovacky(n, m int, s string) (komb.Cifrovacky, error) {
	var C komb.Cifrovacky
	for i, s := range strings.Split(s, " ") {
		c, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return komb.Cifrovacky{}, err
		}
		C[i] = byte(c)
	}
	return C, nil
}
