package psl

import (
	"strconv"
	"strings"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
	// "github.com/melias122/psl/parser"
)

type povinne struct {
	povinne  []bool
	n, count int
}

func PovinneFromString(s string, last Kombinacia, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	p.Zhoda = last
	ints, err := p.ParseInts()
	if err != nil {
		return nil, err
	}
	return Povinne(ints, n, m), nil
}

func Povinne(ints []int, n, m int) Filter {
	p := povinne{
		povinne: make([]bool, m),
		n:       n,
	}
	for _, i := range ints {
		if i >= 1 && i <= m && !p.povinne[i-1] {
			p.povinne[i-1] = true
			p.count++
		}
	}
	if p.count > n {
		p.count = n
	}
	return p
}

func (p povinne) Check(k Kombinacia) bool {
	// if p.n-len(k) > p.count {
	// if len(k) < p.count {
	// 	return true
	// }

	var count int
	for _, i := range k {
		if p.povinne[i-1] {
			count++
		}
	}
	if len(k) == p.n {
		return p.count == count
	}
	return true
}

func (p povinne) CheckSkupina(s Skupina) bool {
	return true
}

func (p povinne) String() string {
	var s []string
	for i, ok := range p.povinne {
		if ok {
			s = append(s, strconv.Itoa(i+1))
		}
	}
	return "Povinne: " + strings.Join(s, ", ")
}

type povinneStl struct {
	povinne [][]bool
	count   int
}

func PovinneSTLFromString(s string, zhoda []byte, n, m int) (Filter, error) {
	parser := NewParser(strings.NewReader(s), n, m)
	parser.Zhoda = zhoda
	ma, err := parser.ParseMapInts()
	if err != nil {
		return nil, err
	}
	// TODO: remove this conversion
	var mapInts map[int][]int
	for k, v := range ma {
		mapInts[k] = v
	}
	return PovinneSTL(mapInts, n, m), nil
}

func PovinneSTL(mapInts map[int][]int, n, m int) Filter {
	p := povinneStl{
		povinne: make([][]bool, n),
	}

	for i, ints := range mapInts {
		for _, j := range ints {
			if !isOkForSTL(j-1, i-1, n, m) {
				continue
			}
			// make slice
			if p.povinne[i-1] == nil {
				p.povinne[i-1] = make([]bool, m)
				p.count++
			}
			// set if not seted
			if !p.povinne[i-1][j-1] {
				p.povinne[i-1][j-1] = true
			}
		}
	}
	return p
}

func (p povinneStl) Check(k Kombinacia) bool {
	for i, j := range k {
		if p.povinne[i] == nil {
			continue
		}
		if !p.povinne[i][j-1] {
			return false
		}
	}
	return true
}

func (p povinneStl) CheckSkupina(s Skupina) bool {
	return true
}

func (p povinneStl) String() string {
	var s []string
	for i := range p.povinne {
		if p.povinne[i] == nil {
			continue
		}
		var ints []string
		for j, ok := range p.povinne[i] {
			if ok {
				ints = append(ints, strconv.Itoa(j+1))
			}
		}
		s = append(s, strconv.Itoa(i+1)+":"+strings.Join(ints, ", "))
	}
	return "Povinne STL: " + strings.Join(s, "; ")
}

func isOkForSTL(number, STL, n, m int) bool {
	// 2:1; 3:1,2; ... || 1:32,33,...; 2:33,...
	if number < STL || number > (m-n+STL) {
		return false
	}
	return true
}
