package filter

import (
	"bytes"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type xtica struct {
	n, m  int
	xtica komb.Tica
}

func NewXtica(n, m int, tica komb.Tica) Filter {
	return xtica{
		n:     n,
		m:     m,
		xtica: tica,
	}
}

func (x xtica) String() string {
	return "Xtica: " + x.xtica.String()
}

func (x xtica) Check(k komb.Kombinacia) bool {
	cmp := bytes.Compare(komb.Xtica(x.m, k), x.xtica)
	if (len(k) == x.n && cmp != 0) || cmp > 0 {
		return false
	}
	return true
}

func (x xtica) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}
