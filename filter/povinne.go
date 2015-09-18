package filter

import (
	"fmt"
	"strings"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type povinne struct {
	n int
	p map[byte]bool
}

func NewPovinne(n int, cisla []byte) Filter {
	p := make(map[byte]bool, len(cisla))
	for _, cislo := range cisla {
		p[cislo] = true
	}
	return povinne{
		n: n,
		p: p,
	}
}

func (p povinne) Check(k komb.Kombinacia) bool {
	var count int
	for _, cislo := range k {
		if p.p[cislo] {
			count++
		}
	}
	var expected int
	if len(p.p) > p.n {
		expected = p.n
	} else {
		expected = len(p.p)
	}
	if len(k) == p.n && expected != count {
		return false
	}
	return true
}

func (p povinne) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (p povinne) String() string {
	var b []byte
	for k := range p.p {
		b = append(b, k)
	}
	return fmt.Sprintf("Povinne: %s", strings.Replace(fmt.Sprintf("%v", b), " ", ", ", -1))
}

type povinneStl struct {
	n       int
	povinne map[stlKey]bool
	stl     []bool
}

func NewPovinneStl(n int, cisla [][]byte) Filter {
	p := povinneStl{
		n:       n,
		povinne: make(map[stlKey]bool),
		stl:     make([]bool, n),
	}
	for i := range cisla {
		for _, c := range cisla[i] {
			p.povinne[stlKey{STL: i, Cislo: c}] = true
			p.stl[i] = true
		}
	}
	return p
}

func (p povinneStl) Check(k komb.Kombinacia) bool {
	for i, c := range k {
		if ok := p.povinne[stlKey{STL: i, Cislo: c}]; !ok && p.stl[i] {
			return false
		}
	}
	return true
}

func (p povinneStl) CheckSkupina(skupina hrx.Skupina) bool {
	return true
}

func (p povinneStl) String() string {
	b := make([][]byte, p.n)
	for k := range p.povinne {
		b[k.STL] = append(b[k.STL], k.Cislo)
	}
	var s string
	for i := range b {
		s += fmt.Sprintf("%d:%s ", i+1, strings.Replace(fmt.Sprintf("%v", b[i]), " ", ", ", -1))
	}
	return "Povinne STL: " + s
}
