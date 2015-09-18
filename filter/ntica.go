package filter

import (
	"bytes"
	"fmt"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type ntica struct {
	n     int
	ntica komb.Tica
}

func NewNtica(n int, tica komb.Tica) Filter {
	return ntica{
		n:     n,
		ntica: tica,
	}
}

func (n ntica) Check(k komb.Kombinacia) bool {
	cmp := bytes.Compare(komb.Ntica(k), n.ntica)
	if (len(k) == n.n && cmp != 0) || cmp > 0 {
		return false
	}
	return true
}

func (n ntica) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (n ntica) String() string {
	return "Ntica:" + n.ntica.String()
}

type stlNtica struct {
	n       int
	pozicie []byte
	ntica   Filter
}

func NewStlNtica(n int, tica komb.Tica, pozicie []byte) Filter {
	return stlNtica{
		n:       n,
		ntica:   NewNtica(n, tica),
		pozicie: pozicie,
	}
}

func (s stlNtica) Check(k komb.Kombinacia) bool {
	if !s.ntica.Check(k) {
		return false
	}
	if s.n == len(k) && bytes.Compare(komb.NticaPozicie(k), s.pozicie) != 0 {
		return false
	}
	return true
}

func (s stlNtica) CheckSkupina(h hrx.Skupina) bool {
	return true
}

func (s stlNtica) String() string {
	return fmt.Sprintf("STL Ntica: %v", s.pozicie)
}
