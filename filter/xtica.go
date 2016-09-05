package filter

import (
	"bytes"

	"gitlab.com/melias122/engine"
)

type filterXtica struct {
	n, m  int
	xtica engine.Tica
}

func NewFilterXtica(n, m int, tica engine.Tica) Filter {
	return &filterXtica{
		n:     n,
		m:     m,
		xtica: tica,
	}
}

func (f *filterXtica) String() string {
	return "Xtica: " + f.xtica.String()
}

func (f *filterXtica) Check(k engine.Kombinacia) bool {
	cmp := bytes.Compare(engine.Xtica(f.m, k), f.xtica)
	if (len(k) == f.n && cmp != 0) || cmp > 0 {
		return false
	}
	return true
}

func (f *filterXtica) CheckSkupina(skupina engine.Skupina) bool {
	return true
}
