package engine

import (
	"sort"
	"strconv"
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
	cmp := func(j int) bool {
		return p[j].Sk >= sk
	}
	i := sort.Search(len(p), cmp)
	ok := i < len(p) && p[i].Sk == sk
	return i, ok
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

func (x Xcisla) copy() Xcisla {
	xcisla := make(Xcisla, len(x))
	copy(xcisla, x)
	return xcisla
}

func (p Xcisla) Len() int           { return len(p) }
func (p Xcisla) Less(i, j int) bool { return p[i].Sk < p[j].Sk }
func (p Xcisla) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (x Xcisla) String() string {
	var (
		zero  = []byte("0 ")
		space = []byte(" ")
		buf   = make([]byte, 0, 128)
		i     int
	)
	for _, t := range x {
		for i < t.Sk {
			buf = append(buf, zero...)
			i++
		}
		buf = strconv.AppendUint(buf, uint64(t.Max), 10)
		buf = append(buf, space...)
		i++
	}
	return string(buf[:len(buf)])
}

func (p Xcisla) Contains(t Tab) bool {
	i := sort.Search(len(p), func(j int) bool { return p[j].Sk >= t.Sk })
	if i < len(p) && p[i].Sk == t.Sk && p[i].Max == t.Max {
		return true
	} else {
		return false
	}
}
