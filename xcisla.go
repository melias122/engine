package psl

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Tab struct {
	Sk  int
	Max int
}

func newTab(sk, max int) Tab {
	return Tab{Sk: sk, Max: max}
}

type Xcisla []Tab

func NewXcisla(m int) Xcisla {
	presun := make([]Tab, 1, m)
	presun[0] = Tab{Sk: 0, Max: m}
	return presun
}

func (p *Xcisla) move(n, from, to int) {
	if from > to {
		from, to = to, from
		n = -n
	}
	p.add(-n, from)
	p.add(n, to)
}

func (p Xcisla) index(sk int) (int, bool) {
	i := sort.Search(len(p), func(j int) bool { return p[j].Sk >= sk })
	if i < len(p) && p[i].Sk == sk {
		return i, true
	} else {
		return i, false
	}
}

func (p *Xcisla) add(n, sk int) {
	i, ok := p.index(sk)
	// ak neexistuje skupina vytvorime ju na pozicii i
	if !ok {
		*p = append(*p, Tab{})
		copy((*p)[i+1:], (*p)[i:])
		(*p)[i] = newTab(sk, 0)
	}
	// pridanie odcitanie z danej skupiny
	(*p)[i].Max += n
	// ak pocet(max) danej skupiny je 0 vymazeme ju
	if (*p)[i].Max <= 0 {
		*p = append((*p)[:i], (*p)[i+1:]...)
	}
}

func (p Xcisla) Max() int {
	return p[len(p)-1].Sk
}

func (p Xcisla) copy() Xcisla {
	presun := make(Xcisla, len(p))
	copy(presun, p)
	return presun
}

func (p Xcisla) Len() int           { return len(p) }
func (p Xcisla) Less(i, j int) bool { return p[i].Sk < p[j].Sk }
func (p Xcisla) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p Xcisla) String() string {
	var (
		i int
		s = make([]string, 0, 2*len(p))
	)
	for _, p := range p {
		for i < p.Sk {
			s = append(s, "0")
			i++
		}
		s = append(s, strconv.Itoa(p.Max))
		i++
	}
	return strings.Join(s, " ")
}

func (p Xcisla) Contains(t Tab) bool {
	i := sort.Search(len(p), func(j int) bool { return p[j].Sk >= t.Sk })
	if i < len(p) && p[i].Sk == t.Sk && p[i].Max == t.Max {
		return true
	} else {
		return false
	}
}

// idea: {skupina1:[pocet1, pocet2, ...], skupina1:[pocet1, pocet2, ...], ...}
// pr: 	{1:[1, 2, ...], 2:[0, 2], 3:[], 6:[], ...}
type filterXcisla struct {
	x [][]Tab
}

// XcislaFromString implementuju filter nad Xcislami.
// Format: "1:1,2,3; 2:1-3,5; 3:1; 3:2"
func NewFilterXcislaFromString(s string, n, m int) (Filter, error) {
	p := NewParser(strings.NewReader(s), n, m)
	mapInts, err := p.ParseMapInts()
	if err != nil {
		return nil, err
	}

	var x Xcisla
	for j, ints := range mapInts {
		for _, i := range ints {
			x = append(x, Tab{Sk: j, Max: i})
		}
	}
	return NewFilterXcisla(x), nil
}

// Xcisla impletuju filter pre hrx.Presun (xcisla). Vstup je vektor tabuliek
func NewFilterXcisla(xcisla Xcisla) Filter {
	var (
		x           filterXcisla
		skupinaLast = -1
	)
	sort.Sort(xcisla)
	for _, t := range xcisla {
		if t.Sk != skupinaLast {
			x.x = append(x.x, []Tab{})
			skupinaLast = t.Sk
		}
		i := len(x.x) - 1
		x.x[i] = append(x.x[i], t)
	}
	return x
}

func (x filterXcisla) Check(Kombinacia) bool {
	return true
}

func (x filterXcisla) CheckSkupina(h Skupina) bool {
	for _, tabs := range x.x {
		var count int
		for _, t := range tabs {
			if h.Xcisla.Contains(t) {
				count++
			}
		}
		if count == 0 {
			return false
		}
	}
	return true
}

func (x filterXcisla) String() string {
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
