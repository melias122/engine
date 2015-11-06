package psl

import (
	"bytes"
)

func Xtica(m int, k Kombinacia) Tica {
	xtica := make(Tica, (m+9)/10)
	for _, n := range k {
		xtica[(n-1)/10]++
	}
	return xtica
}

type filterXtica struct {
	n, m  int
	xtica Tica
}

func NewFilterXtica(n, m int, tica Tica) Filter {
	return filterXtica{
		n:     n,
		m:     m,
		xtica: tica,
	}
}

func (f filterXtica) String() string {
	return "Xtica: " + f.xtica.String()
}

func (f filterXtica) Check(k Kombinacia) bool {
	cmp := bytes.Compare(Xtica(f.m, k), f.xtica)
	if (len(k) == f.n && cmp != 0) || cmp > 0 {
		return false
	}
	return true
}

func (f filterXtica) CheckSkupina(skupina Skupina) bool {
	return true
}
