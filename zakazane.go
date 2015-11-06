package psl

import (
	"strconv"
	"strings"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
	// "github.com/melias122/psl/parser"
)

type zakazane struct {
	cisla []bool
}

func ZakazaneFromString(s string, zhoda []byte, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	p.Zhoda = zhoda
	ints, err := p.ParseInts()
	if err != nil {
		return nil, err
	}
	return Zakazane(ints, n, m), nil
}

func Zakazane(ints []int, n, m int) Filter {
	z := make([]bool, m)
	for _, i := range ints {
		z[i-1] = true
	}
	return zakazane{
		cisla: z,
	}
}

func (z zakazane) Check(k Kombinacia) bool {
	for _, c := range k {
		if z.cisla[c-1] {
			return false
		}
	}
	return true
}

func (z zakazane) CheckSkupina(skupina Skupina) bool {
	return true
}

func (z zakazane) String() string {
	var s []string
	for c, ok := range z.cisla {
		if ok {
			s = append(s, strconv.Itoa(c+1))
		}
	}
	return "Zakázané:" + strings.Join(s, ", ")
}

type zakazaneStl struct {
	zakazane [][]bool
}

func ZakazaneSTLFromString(s string, zhoda []byte, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	p.Zhoda = zhoda
	mi, err := p.ParseMapInts()
	if err != nil {
		return nil, err
	}
	var mapInts map[int][]int
	for k, v := range mi {
		mapInts[k] = v
	}
	return ZakazaneSTL(mapInts, n, m), nil
}

func ZakazaneSTL(mapInts map[int][]int, n, m int) Filter {
	z := zakazaneStl{
		zakazane: make([][]bool, n),
	}
	for i := range mapInts {
		if z.zakazane[i-1] == nil {
			z.zakazane[i-1] = make([]bool, m)
		}
		for _, j := range mapInts[i] {
			z.zakazane[i-1][j-1] = true
		}
	}
	return z
}

func (z zakazaneStl) Check(k Kombinacia) bool {
	for i, j := range k {
		if z.zakazane[i] == nil {
			continue
		}
		if z.zakazane[i][j-1] {
			return false
		}
	}
	return true
}

func (z zakazaneStl) CheckSkupina(skupina Skupina) bool {
	return true
}

func (z zakazaneStl) String() string {
	var s []string
	for i := range z.zakazane {
		if z.zakazane[i] == nil {
			continue
		}
		var ints []string
		for j, ok := range z.zakazane[i] {
			if ok {
				ints = append(ints, strconv.Itoa(j+1))
			}
		}
		s = append(s, strconv.Itoa(i+1)+":"+strings.Join(ints, ", "))
	}
	return "Zakázané STL: " + strings.Join(s, "; ")
}
