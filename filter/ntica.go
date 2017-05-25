package filter

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/melias122/engine"
)

type filterNtica struct {
	n     int
	ntica engine.Tica
}

func NewFilterNtica(n int, tica engine.Tica) Filter {
	return &filterNtica{
		n:     n,
		ntica: tica,
	}
}

func (n *filterNtica) Check(k engine.Kombinacia) bool {
	nticaK := engine.Ntica(k)
	if len(k) == n.n {
		return bytes.Equal(nticaK, n.ntica)
	}
	return true
}

func (n *filterNtica) CheckSkupina(skupina engine.Skupina) bool {
	return true
}

func (n *filterNtica) String() string {
	return "Ntica: " + n.ntica.String()
}

type filterSTLNtica struct {
	n       int
	pozicie []byte
	ntica   Filter
}

func NewFilterSTLNtica(n int, tica engine.Tica, pozicie []byte) Filter {
	return &filterSTLNtica{
		n:       n,
		ntica:   NewFilterNtica(n, tica),
		pozicie: pozicie,
	}
}

func (s *filterSTLNtica) Check(k engine.Kombinacia) bool {
	if !s.ntica.Check(k) {
		return false
	}
	if s.n == len(k) {
		return bytes.Equal(engine.NticaPozicie(k), s.pozicie)
	}
	return true
}

func (s *filterSTLNtica) CheckSkupina(h engine.Skupina) bool {
	return true
}

func (s *filterSTLNtica) String() string {
	var pozicie []string
	for i, p := range s.pozicie {
		if p == 1 {
			pozicie = append(pozicie, strconv.Itoa(i+1))
		}
	}
	return s.ntica.String() + "\n" + "STL Ntica: " + strings.Join(pozicie, ", ")
}
