package filter

import (
	"bytes"
	"strconv"
	"strings"

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
	nticaK := komb.Ntica(k)
	if len(k) == n.n {
		return bytes.Equal(nticaK, n.ntica)
	}
	return true
}

func (n ntica) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (n ntica) String() string {
	return "Ntica: " + n.ntica.String()
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
	if s.n == len(k) {
		return bytes.Equal(komb.NticaPozicie(k), s.pozicie)
	}
	return true
}

func (s stlNtica) CheckSkupina(h hrx.Skupina) bool {
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
