package psl

import (
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

type Presun []Tab

func NewPresun(m int) Presun {
	presun := make([]Tab, 1, m)
	presun[0] = Tab{Sk: 0, Max: m}
	return presun
}

func (p *Presun) move(n, from, to int) {
	if from > to {
		from, to = to, from
		n = -n
	}
	p.add(-n, from)
	p.add(n, to)
}

func (p Presun) index(sk int) (int, bool) {
	i := sort.Search(len(p), func(j int) bool { return p[j].Sk >= sk })
	if i < len(p) && p[i].Sk == sk {
		return i, true
	} else {
		return i, false
	}
}

func (p *Presun) add(n, sk int) {
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

func (p Presun) Max() int {
	return p[len(p)-1].Sk
}

func (p Presun) copy() Presun {
	presun := make(Presun, len(p))
	copy(presun, p)
	return presun
}

func (p Presun) Len() int           { return len(p) }
func (p Presun) Less(i, j int) bool { return p[i].Sk < p[j].Sk }
func (p Presun) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p Presun) String() string {
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

func (p Presun) Contains(t Tab) bool {
	i := sort.Search(len(p), func(j int) bool { return p[j].Sk >= t.Sk })
	if i < len(p) && p[i].Sk == t.Sk && p[i].Max == t.Max {
		return true
	} else {
		return false
	}
}
