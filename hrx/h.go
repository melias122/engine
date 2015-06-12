package hrx

import (
	"math"

	"github.com/melias122/psl/num"
)

type H struct {
	m   int
	sk  map[int]int      // pocetnost(skupina), pocet cisel
	get func(*num.N) int // cislo -> pocet
}

func new(m int, f func(*num.N) int) *H {
	h := &H{
		m:   m,
		sk:  make(map[int]int, m),
		get: f,
	}
	h.sk[0] = m
	return h
}

func NewHHrx(m int) *H {
	return new(m, func(N *num.N) int { return N.PocR1() })
}

func NewHrx(m int) *H {
	return new(m, func(N *num.N) int { return N.PocR2() })
}

func (h *H) Add(n *num.N) {
	sk := h.get(n)
	h.Move(1, sk-1, sk)
}

func (h *H) Move(pocet, from, to int) {
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
		// max = float64(19)
	)
	if max == 0 {
		return 100.0
	}
	for k, v := range h.sk {
		hrx += ((float64(v) / float64(h.m)) * math.Pow((max-float64(k))/max, 16))
	}
	hrx = math.Sqrt(math.Sqrt(hrx)) * 100
	return hrx
}

// func (h *H) Simul(p Presun) float64 {
// 	// z aktualnej skupiny potrebujem preniest t.Max
// 	// do dalsej skupiny sk+1
// 	for _, t := range p {
// 		h.Move(t.Max, t.Sk, t.Sk+1)
// 	}
// 	// Vypocitaj hrx pre zostavu p
// 	hrx := h.Get()
// 	// Obnov povodny stav
// 	for _, t := range p {
// 		h.Move(t.Max, t.Sk+1, t.Sk)
// 	}
// 	return hrx
// }

// func (h *H) Presun() Presun {
// 	p := make(Presun, 0, len(h.sk))
// 	for k, v := range h.sk {
// 		p = append(p, Tab{k, v})
// 	}
// 	sort.Sort(p)
// 	return p
// }

// type Tab struct {
// 	Sk  int
// 	Max int
// }

// type Presun []Tab

// func (p Presun) copy() Presun {
// 	cp := make(Presun, len(p))
// 	for i := range p {
// 		cp[i] = p[i]
// 	}
// 	return cp
// }

// func (p Presun) Len() int           { return len(p) }
// func (p Presun) Less(i, j int) bool { return p[i].Sk < p[j].Sk }
// func (p Presun) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// func (p Presun) String() string {
// 	if len(p) > 0 {
// 		s := make([]int, p[len(p)-1].Sk+1)
// 		s[0] = 0
// 		for _, v := range p {
// 			s[v.Sk] = v.Max
// 		}
// 		s2 := make([]string, len(s))
// 		for i := range s {
// 			s2[i] = strconv.Itoa(s[i])
// 		}
// 		return strings.Join(s2, " ")
// 	} else {
// 		return ""
// 	}
// }

// func GenerujPresun(t []Tab, n int) chan Presun {
// 	ch := make(chan Presun)
// 	go func() {
// 		defer close(ch)
// 		generujPresun(t,
// 			n,
// 			make(Presun, 0, len(t)),
// 			ch)
// 	}()
// 	return ch
// }

// func generujPresun(t []Tab, n int, p Presun, ch chan Presun) {
// 	if len(t) == 0 {
// 		return
// 	}

// 	max := t[0].Max
// 	if max > n {
// 		max = n
// 	}
// 	for max > 0 {
// 		p = append(p, Tab{t[0].Sk, max})
// 		if n-max > 0 {
// 			generujPresun(t[1:], n-max, p, ch)
// 		} else {
// 			ch <- p.copy()
// 		}
// 		p = p[:len(p)-1]
// 		max--
// 	}
// 	generujPresun(t[1:], n, p, ch)
// }

type Skupina []*num.N

func (s Skupina) Id() int {
	return s[0].PocR2()
}

func (s Skupina) Max() int {
	return len(s)
}

type Presun []Skupina

func (p Presun) Len() int           { return len(p) }
func (p Presun) Less(i, j int) bool { return p[i].Id() < p[j].Id() }
func (p Presun) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// func (p Presun) String() string {
// 	if len(p) > 0 {
// 		s := make([]int, p[len(p)-1].Sk+1)
// 		s[0] = 0
// 		for _, v := range p {
// 			s[v.Sk] = v.Max
// 		}
// 		s2 := make([]string, len(s))
// 		for i := range s {
// 			s2[i] = strconv.Itoa(s[i])
// 		}
// 		return strings.Join(s2, " ")
// 	} else {
// 		return ""
// 	}
// }

// type HrxTab struct {
// 	hrx       *H
// 	hrxR      float64
// 	presun    []byte
// 	sucetMin  int
// 	sucetMax  int
// 	pocetKomb big.Int
// 	hhrxMin   float64
// 	hhrxMax   float64
// }

// func (h *HrxTab) Add(s []*num.N, max int) {
// 	skupina := s[0].PocR2()
// 	// presun add skupina

// }

// func (h *HrxTab) Delete(s []*num.N, max int) {

// }

// func (h *HrxTab) Make(skupiny [][]*num.N, n int) {
// 	if len(skupiny) == 0 {
// 		return
// 	}
// 	max := len(skupiny[0])
// 	if max > n {
// 		max = n
// 	}
// 	for max > 0 {
// 		s := skupiny[0]
// 		h.Add(s, max)
// 		if n-max > 0 {
// 			generujPresun(skupiny[1:], n-max, p, ch)
// 		} else {
// 			// nasli sme hrx
// 		}
// 		h.Delete(s, max)
// 		max--
// 	}
// 	generujPresun(t[1:], n, p, ch)
// }

// func sucetMinMax(m map[int][]*num.N, p hrx.Presun) string {
// 	min, max := 0, 0
// 	for _, t := range p {
// 		arr := m[t.Sk]
// 		for i := 0; i < t.Max; i++ {
// 			min += arr[i].Cislo()
// 			max += arr[len(arr)-1-i].Cislo()
// 		}
// 	}
// 	return strings.Join([]string{itoa(min), itoa(max)}, "-")
// }

// func pocetSucet(m map[int][]*num.N, p hrx.Presun) string {
// 	pocet := big.NewInt(1)
// 	var b big.Int
// 	for _, t := range p {
// 		arr := m[t.Sk]
// 		pocet.Mul(pocet, b.Binomial(int64(len(arr)), int64(t.Max)))
// 	}
// 	return pocet.String()
// }
