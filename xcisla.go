package psl

import (
	"fmt"
	"sort"
	"strings"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
	// "github.com/melias122/psl/parser"
)

// idea: {skupina1:[pocet1, pocet2, ...], skupina1:[pocet1, pocet2, ...], ...}
// pr: 	{1:[1, 2, ...], 2:[0, 2], 3:[], 6:[], ...}
type xcisla struct {
	x [][]Tab
}

// XcislaFromString implementuju filter nad Xcislami.
// Format: "1:1,2,3; 2:1-3,5; 3:1; 3:2"
func XcislaFromString(s string, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	mapInts, err := p.ParseMapInts()
	if err != nil {
		return nil, err
	}

	var x Presun
	for j, ints := range mapInts {
		for _, i := range ints {
			x = append(x, Tab{Sk: j, Max: i})
		}
	}
	return Xcisla(x), nil
}

// Xcisla impletuju filter pre hrx.Presun (xcisla). Vstup je vektor tabuliek
func Xcisla(tabs Presun) Filter {
	var (
		x           xcisla
		skupinaLast = -1
	)
	sort.Sort(tabs)
	for _, t := range tabs {
		if t.Sk != skupinaLast {
			x.x = append(x.x, []Tab{})
			skupinaLast = t.Sk
		}
		i := len(x.x) - 1
		x.x[i] = append(x.x[i], t)
	}
	return x
}

func (x xcisla) Check(Kombinacia) bool {
	return true
}

func (x xcisla) CheckSkupina(h Skupina) bool {
	for _, tabs := range x.x {
		var count int
		for _, t := range tabs {
			if h.Presun.Contains(t) {
				count++
			}
		}
		if count == 0 {
			return false
		}
	}
	return true
}

func (x xcisla) String() string {
	var s []string
	for _, tabs := range x.x {
		var str string
		if len(tabs) > 0 {
			str = fmt.Sprintf("%d:", tabs[0].Sk)
		}
		var str2 []string
		for _, t := range tabs {
			str2 = append(str2, fmt.Sprint(t.Max))
		}
		str += strings.Join(str2, ",")
		s = append(s, str)
	}
	return fmt.Sprintf("%s", strings.Join(s, ";"))
}
