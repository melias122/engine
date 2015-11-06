package psl

import (
	"bytes"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
)

type xticaFilter struct {
	n, m  int
	xtica Tica
}

func NewXtica(n, m int, tica Tica) Filter {
	return xticaFilter{
		n:     n,
		m:     m,
		xtica: tica,
	}
}

func (x xticaFilter) String() string {
	return "Xtica: " + x.xtica.String()
}

func (x xticaFilter) Check(k Kombinacia) bool {
	cmp := bytes.Compare(Xtica(x.m, k), x.xtica)
	if (len(k) == x.n && cmp != 0) || cmp > 0 {
		return false
	}
	return true
}

func (x xticaFilter) CheckSkupina(skupina Skupina) bool {
	return true
}
