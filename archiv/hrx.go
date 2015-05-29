package archiv

import (
	"fmt"
	"math"
	"psl2/num"
)

func (h *hrx) dbg() string {
	return fmt.Sprint(h.max, h.sk)
}

type hrx struct {
	m   int
	max int
	sk  map[int]int      // pocetnost, pocet cisel
	get func(*num.N) int // cislo -> pocet
}

func newhrx(m int, f func(*num.N) int) *hrx {
	h := hrx{
		m:   m,
		sk:  make(map[int]int, m),
		get: f,
	}
	h.sk[0] = m
	return &h
}

func (h *hrx) add(n *num.N) {
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

func (h *hrx) hrx() float64 {
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

// func (a *Archiv) HrxHHrx() error {
// 	header := []string{"p.c.", "HRX pre r+1", "dHRX diferencia s \"r\"", "presun z r do (r+1)cisla", "∑%ROD-DO", "∑%STLOD-DO od do", "∑ kombi od do",
// 		"Pocet ∑ kombi", "HHRX pre r+1", "dHHRX diferencia s \"r\"", "∑%R1-DO od do", "Teor. max. pocet", "∑%R1-DO", "∑%STL1-DO od do",
// 	}
// 	f, err := os.Open(fmt.Sprintf("%d%d/HrxHHrx_%d%d.csv", a.n, a.m, a.n, a.m))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	w := csv.NewWriter(f)
// 	if err = w.Write(header); err != nil {
// 		return err
// 	}
// 	makeHrx()
// 	w.Flush()
// 	return w.Error()
// }

// type x struct {
// 	sk int
// 	c  []int
// }

// var (
// 	sk = []x{
// 		{3, []int{1, 9, 12, 15, 17, 21, 22, 24, 25, 26}},
// 		{2, []int{6, 13, 16, 18, 31, 35}},
// 		{5, []int{5, 10, 19, 27, 30}},
// 		{4, []int{2, 14, 33, 34}},
// 		{6, []int{4, 11, 29}},
// 		{8, []int{3, 20, 32}},
// 		{7, []int{23, 28}},
// 		{1, []int{8}},
// 		{9, []int{8}},
// 		{10, []int{7}},
// 	}
// )

// type w struct {
// 	s, w int
// }

// var y = make([]w, 0, 3)
// var spolu int

// func ss(sk []x, n int) {
// 	if len(sk) == 0 {
// 		return
// 	}
// 	var xx int
// 	if n > len(sk[0].c) {
// 		xx = len(sk[0].c)
// 	} else {
// 		xx = n
// 	}
// 	for i := xx; i > 0; i-- {
// 		y = append(y, w{sk[0].sk, i})
// 		if n-i > 0 {
// 			ss(sk[1:], n-i)
// 		} else {
// 			fmt.Println(y)
// 			spolu++
// 		}
// 		y = y[:len(y)-1]
// 	}
// 	ss(sk[1:], n)
// }

// func test() {
// 	ss(sk, 30)
// 	fmt.Println(spolu)
// }
