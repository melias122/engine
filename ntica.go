package psl

import (
	"bytes"
	"strconv"
	"strings"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
)

type nticaFilter struct {
	n     int
	ntica Tica
}

func NewNtica(n int, tica Tica) Filter {
	return nticaFilter{
		n:     n,
		ntica: tica,
	}
}

func (n nticaFilter) Check(k Kombinacia) bool {
	nticaK := Ntica(k)
	if len(k) == n.n {
		return bytes.Equal(nticaK, n.ntica)
	}
	return true
}

func (n nticaFilter) CheckSkupina(skupina Skupina) bool {
	return true
}

func (n nticaFilter) String() string {
	return "Ntica: " + n.ntica.String()
}

type stlNtica struct {
	n       int
	pozicie []byte
	ntica   Filter
}

func NewStlNtica(n int, tica Tica, pozicie []byte) Filter {
	return stlNtica{
		n:       n,
		ntica:   NewNtica(n, tica),
		pozicie: pozicie,
	}
}

func (s stlNtica) Check(k Kombinacia) bool {
	if !s.ntica.Check(k) {
		return false
	}
	if s.n == len(k) {
		return bytes.Equal(NticaPozicie(k), s.pozicie)
	}
	return true
}

func (s stlNtica) CheckSkupina(h Skupina) bool {
	return true
}

func (s stlNtica) String() string {
	var pozicie []string
	for i, p := range s.pozicie {
		if p == 1 {
			pozicie = append(pozicie, strconv.Itoa(i+1))
		}
	}
	return s.ntica.String() + "\n" + "STL Ntica: " + strings.Join(pozicie, ", ")
}
