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
	sk  map[int]int      // pocetnost(skupina), pocet cisel
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
	sk := h.get(n)
	h.move(1, sk-1, sk)
}

func (h *H) move(pocet, from, to int) {
	if h.sk[from] > pocet {
		h.sk[from] -= pocet
	} else {
		delete(h.sk, from)
	}
	h.sk[to] += pocet
}

func (h *H) max() int {
	var max int
	for sk := range h.sk {
		if sk > max {
			max = sk
		}
	}
	return max
}

func (h *H) Get() float64 {
	var (
		hrx float64
		max = float64(h.max())
	)
	if max == 0 {
		return 100.0
	}
	for k, v := range h.sk {
		hrx += ((float64(v) / float64(h.m)) * math.Pow((max-float64(k))/max, 16))
	}
	hrx = math.Pow(hrx, 0.25)
	hrx *= 100
	return hrx
}

func (h *H) Simul(p Presun) float64 {
	// z aktualnej skupiny potrebujem preniest t.Max
	// do dalsej skupiny sk+1
	for _, t := range p {
		h.move(t.Max, t.Sk, t.Sk+1)
	}
	// Vypocitaj hrx pre zostavu p
	hrx := h.Get()
	// Obnov povodny stav
	for _, t := range p {
		h.move(t.Max, t.Sk+1, t.Sk)
	}
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
	Sk  int
	Max int
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
func (p Presun) Less(i, j int) bool { return p[i].Sk < p[j].Sk }
func (p Presun) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p Presun) String() string {
	if len(p) > 0 {
		s := make([]int, p[len(p)-1].Sk+1)
		s[0] = 0
		for _, v := range p {
			s[v.Sk] = v.Max
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

	max := t[0].Max
	if max > n {
		max = n
	}
	for max > 0 {
		p = append(p, Tab{t[0].Sk, max})
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
