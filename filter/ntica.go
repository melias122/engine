package filter

import (
	"bytes"

	"github.com/melias122/psl/komb"
)

// TODO: stlNtica
type stlNtica struct {
}

func NewStlNtica() Filter {
	return nil
}

func (s stlNtica) Check(k komb.Kombinacia) bool {
	return true
}

type ntica struct {
	n     int
	ntica komb.Tica
}

func NewNtica(n int, tica komb.Tica) Filter {
	return ntica{
		ntica: tica,
	}
}

func (n ntica) Check(k komb.Kombinacia) bool {
	if len(k) == n.n {
		if bytes.Compare(komb.Ntica(k), n.ntica) != 0 {
			return false
		}
	}
	return true
}
