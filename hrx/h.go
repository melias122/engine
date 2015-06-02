package hrx

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/melias122/psl/num"
)

type H struct {
	m   int
	max int
	sk  map[int]int      // pocetnost, pocet cisel
	get func(*num.N) int // cislo -> pocet
}

func New(m int, f func(*num.N) int) *H {
	h := &H{
		m:   m,
		sk:  make(map[int]int, m),
		get: f,
	}
	h.sk[0] = m
	return h
}

func (h *H) Add(n *num.N) {
	p := h.get(n)
	if p > 0 {
		if h.sk[p-1] > 1 {
			h.sk[p-1]--
		} else {
			delete(h.sk, p-1)
		}
	}
	h.sk[p]++
	if p > h.max {
		h.max = p
	}
}

func (h *H) Get() float64 {
	if h.max == 0 {
		return 100.0
	}
	var hrx float64
	for k, v := range h.sk {
		hrx += ((float64(v) / float64(h.m)) *
			math.Pow((float64(h.max)-float64(k))/float64(h.max), 16))
	}
	hrx = math.Pow(hrx, 0.25)
	hrx *= 100
	return hrx
}

func (h *H) Presun() Presun {
	p := make(Presun, 0, len(h.sk))
	for k, v := range h.sk {
		p = append(p, Tab{k, v})
	}
	sort.Sort(p)
	return p
}

type Tab struct {
	sk  int
	max int
}

type Presun []Tab

func (p Presun) copy() Presun {
	cp := make(Presun, len(p))
	for i := range p {
		cp[i] = p[i]
	}
	return cp
}

func (p Presun) Len() int           { return len(p) }
func (p Presun) Less(i, j int) bool { return p[i].sk < p[j].sk }
func (p Presun) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p Presun) String() string {
	if len(p) > 0 {
		s := make([]int, p[len(p)-1].sk+1)
		s[0] = 0
		for _, v := range p {
			s[v.sk] = v.max
		}
		s2 := make([]string, len(s))
		for i := range s {
			s2[i] = strconv.Itoa(s[i])
		}
		return strings.Join(s2, " ")
	} else {
		return ""
	}
}

func GenerujPresun(t []Tab, n int) chan Presun {
	ch := make(chan Presun)
	go func() {
		defer close(ch)
		generujPresun(t,
			n,
			make(Presun, 0, len(t)),
			ch)
	}()
	return ch
}

func generujPresun(t []Tab, n int, p Presun, ch chan Presun) {
	if len(t) == 0 {
		return
	}

	max := t[0].max
	if max > n {
		max = n
	}
	for max > 0 {
		p = append(p, Tab{t[0].sk, max})
		if n-max > 0 {
			generujPresun(t[1:], n-max, p, ch)
		} else {
			ch <- p.copy()
		}
		p = p[:len(p)-1]
		max--
	}
	generujPresun(t[1:], n, p, ch)
}
