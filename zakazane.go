package engine

import (
	"strconv"
	"strings"
)

type filterZakazane struct {
	cisla []bool
}

func NewFilterZakazaneFromString(s string, zhoda []byte, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	p.Zhoda = zhoda
	ints, err := p.ParseInts()
	if err != nil {
		return nil, err
	}
	return NewFilterZakazane(ints, n, m), nil
}

func NewFilterZakazane(ints []int, n, m int) Filter {
	z := make([]bool, m)
	for _, i := range ints {
		if i > 0 && i <= m {
			z[i-1] = true
		}
	}
	return filterZakazane{
		cisla: z,
	}
}

func (f filterZakazane) Check(k Kombinacia) bool {
	for _, c := range k {
		if f.cisla[c-1] {
			return false
		}
	}
	return true
}

func (f filterZakazane) CheckSkupina(skupina Skupina) bool {
	return true
}

func (f filterZakazane) String() string {
	var s []string
	for c, ok := range f.cisla {
		if ok {
			s = append(s, strconv.Itoa(c+1))
		}
	}
	return "Zakázané:" + strings.Join(s, ", ")
}

type filterZakazaneSTL struct {
	zakazane [][]bool
}

func NewFilterZakazaneSTLFromString(s string, zhoda []byte, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	p.Zhoda = zhoda
	mapInts, err := p.ParseMapInts()
	if err != nil {
		return nil, err
	}
	return NewFilterZakazaneSTL(mapInts, n, m), nil
}

func NewFilterZakazaneSTL(mapInts MapInts, n, m int) Filter {
	z := filterZakazaneSTL{
		zakazane: make([][]bool, n),
	}
	for i := range mapInts {
		if i <= 0 {
			continue
		}
		if z.zakazane[i-1] == nil {
			z.zakazane[i-1] = make([]bool, m)
		}
		for _, j := range mapInts[i] {
			if j > 0 && j <= m {
				z.zakazane[i-1][j-1] = true
			}
		}
	}
	return z
}

func (z filterZakazaneSTL) Check(k Kombinacia) bool {
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

func (z filterZakazaneSTL) CheckSkupina(skupina Skupina) bool {
	return true
}

func (z filterZakazaneSTL) String() string {
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
