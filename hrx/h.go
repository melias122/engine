// package hrx

// import "psl2/num"

// type Func func(*num.N) int

// type s struct {
// 	nums map[int]*num.N
// }

// type H struct {
// 	m   int
// 	max int
// 	sk  map[int]int // pocetnost, pocet cisel
// 	f   Func        // cislo -> pocet
// }

// func New(m int, f Func) *H {
// 	h := hrx{
// 		m:  m,
// 		sk: make(map[int]int, m),
// 		f:  f,
// 	}
// 	h.sk[0] = m
// 	return &h
// }

// func (h *H) add(n *num.N) {
// 	p := h.f(n)
// 	if p > 0 {
// 		if h.sk[p-1] > 1 {
// 			h.sk[p-1]--
// 		} else {
// 			delete(h.sk, p-1)
// 		}
// 	}
// 	h.sk[p]++
// 	if p > h.max {
// 		h.max = p
// 	}
// }

// func (h *H) hrx() float64 {
// 	if h.max == 0 {
// 		return 100.0
// 	}
// 	var hrx float64
// 	for k, v := range h.sk {
// 		hrx += ((float64(v) / float64(h.m)) *
// 			math.Pow((float64(h.max)-float64(k))/float64(h.max), 16))
// 	}
// 	hrx = math.Pow(hrx, 0.25)
// 	hrx *= 100
// 	return hrx
// }
