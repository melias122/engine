package filter

import (
	"fmt"
	"strings"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type zakazane struct {
	cisla []bool
}

func NewZakazane(m int, cisla []byte) Filter {
	z := make([]bool, m)
	for _, c := range cisla {
		z[c-1] = true
	}
	return zakazane{
		cisla: z,
	}
}

func (z zakazane) String() string {
	var b []int
	for c, ok := range z.cisla {
		if ok {
			b = append(b, c+1)
		}
	}
	return fmt.Sprintf("Zakázané: %s", strings.Replace(fmt.Sprintf("%v", b), " ", ", ", -1))
}

func (z zakazane) Check(k komb.Kombinacia) bool {
	for _, cislo := range k {
		if z.cisla[cislo-1] {
			return false
		}
	}
	return true
}

func (z zakazane) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

type stlKey struct {
	STL   int
	Cislo byte
}

type zakazaneStl struct {
	n        int
	zakazane map[stlKey]bool
}

func NewZakazaneStl(n int, cisla [][]byte) Filter {
	z := zakazaneStl{
		n:        n,
		zakazane: make(map[stlKey]bool),
	}
	for i := range cisla {
		for _, c := range cisla[i] {
			z.zakazane[stlKey{STL: i, Cislo: c}] = true
		}
	}
	return z
}

func (z zakazaneStl) String() string {
	b := make([][]byte, z.n)
	for k := range z.zakazane {
		b[k.STL] = append(b[k.STL], k.Cislo)
	}
	var s string
	for i := range b {
		s += fmt.Sprintf("%d:%s ", i+1, strings.Replace(fmt.Sprintf("%v", b[i]), " ", ", ", -1))
	}
	return "Zakázané STL: " + s
}

func (z zakazaneStl) Check(k komb.Kombinacia) bool {
	for i, c := range k {
		if z.zakazane[stlKey{STL: i, Cislo: c}] {
			return false
		}
	}
	return true
}

func (z zakazaneStl) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}
