package engine

import (
	"strconv"
	"strings"
)

type filterPovinne struct {
	povinne  []bool
	n, count int
	filterPriority
}

func NewFilterPovinneFromString(s string, k0 Kombinacia, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	p.Zhoda = k0
	ints, err := p.ParseInts()
	if err != nil {
		return nil, err
	}
	return NewFilterPovinne(ints, n, m), nil
}

func NewFilterPovinne(ints []int, n, m int) Filter {
	p := filterPovinne{
		povinne:        make([]bool, m),
		n:              n,
		filterPriority: P1,
	}
	for _, i := range ints {
		if i > 0 && i <= m && !p.povinne[i-1] {
			p.povinne[i-1] = true
			p.count++
		}
	}
	if p.count > n {
		p.count = n
	}
	return p
}

func (p filterPovinne) Check(k Kombinacia) bool {
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

func (p filterPovinne) CheckSkupina(s Skupina) bool {
	return true
}

func (p filterPovinne) String() string {
	var s []string
	for i, ok := range p.povinne {
		if ok {
			s = append(s, strconv.Itoa(i+1))
		}
	}
	return "Povinne: " + strings.Join(s, ", ")
}

type filterPovinneSTL struct {
	povinne [][]bool
	count   int
	filterPriority
}

func NewFilterPovinneSTLFromString(s string, k0 Kombinacia, n, m int) (Filter, error) {
	parser := NewParser(strings.NewReader(s), n, m)
	parser.Zhoda = k0
	mapInts, err := parser.ParseMapInts()
	if err != nil {
		return nil, err
	}
	return NewFilterPovinneSTL(mapInts, n, m), nil
}

func NewFilterPovinneSTL(mapInts MapInts, n, m int) Filter {
	p := filterPovinneSTL{
		povinne:        make([][]bool, n),
		filterPriority: P1,
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

func (p filterPovinneSTL) Check(k Kombinacia) bool {
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

func (p filterPovinneSTL) CheckSkupina(s Skupina) bool {
	return true
}

func (p filterPovinneSTL) String() string {
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
