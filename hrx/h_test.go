package hrx

import (
	"math"
	"testing"
)

func BenchmarkPow16(b *testing.B) {
	x := 1.1234567890
	y := 16.0
	for i := 0; i < b.N; i++ {
		math.Pow(x, y)
	}
}

func BenchmarkPow16Multi(b *testing.B) {
	x := 1.1234567890
	// y := 16.0
	for i := 0; i < b.N; i++ {
		x *= x
		x *= x
		x *= x
		x *= x
	}
}

func TestPow(t *testing.T) {
	x := 1.1234567890
	t.Log(math.Pow(x, 16))
	x *= x
	x *= x
	x *= x
	x *= x
	t.Log(x)
}

// func TestG(t *testing.T) {
// 	p := Presun{
// 		{0, 0},
// 		{1, 1},
// 		{2, 6},
// 		{3, 10},
// 		{4, 4},
// 		{5, 5},
// 		{6, 3},
// 		{7, 2},
// 		{8, 3},
// 		{9, 0},
// 		{10, 1},
// 	}
// 	h := NewHrxTab(p, 5)
// 	h.Make()
// }

// func TestGenerujPresun(t *testing.T) {
// 	// sk := []Tab{
// 	// 	{1, []int{8}},
// 	// 	{2, []int{6, 13, 16, 18, 31, 35}},
// 	// 	{4, []int{2, 14, 33, 34}},
// 	// 	{3, []int{1, 9, 12, 15, 17, 21, 22, 24, 25, 26}},
// 	// 	{5, []int{5, 10, 19, 27, 30}},
// 	// 	{6, []int{4, 11, 29}},
// 	// 	{7, []int{23, 28}},
// 	// 	{8, []int{3, 20, 32}},
// 	// 	{10, []int{7}},
// 	// }
// 	n := 5
// 	sk := Presun{
// 		{0, 0},
// 		{1, 1},
// 		{2, 6},
// 		// {3, 10},
// 		// {4, 4},
// 		// {5, 5},
// 		// {6, 3},
// 		// {7, 2},
// 		// {8, 3},
// 		// {9, 0},
// 		// {10, 1},
// 	}
// 	sums := 0
// 	p := Presun{}
// 	i := 0
// 	for {

// 		max := sk[i].Max
// 		if max > n {
// 			max = n
// 		}

// 		n -= max
// 		if n > 0 {
// 			p = append(p, Tab{i, max})
// 			i++
// 			continue
// 		} else {
// 			p = append(p, Tab{i, max})
// 		}

// 		fmt.Println(p)

// 		sums++
// 		if sums == 10 {
// 			break
// 		}

// 		if i == len(sk)-1 {
// 			n += sk[i].Max
// 			sk[i].Max = 0
// 			p = p[:i]
// 			i--
// 		}

// 		for i > 0 {
// 			if p[i].Max > 0 {
// 				p[i].Max--
// 				n++
// 			}
// 			if p[i].Max == 0 {
// 				p = p[:i]
// 				i--
// 				continue
// 			} else {
// 				break
// 			}
// 		}

// 		// if i < len(sk)-1 {
// 		// 	i++
// 		// }
// 	}
// }

// var (
// 	//Hrx
// 	h *H = &H{
// 		m:   35,
// 		max: 10,
// 		sk: map[int]int{
// 			1:  1,
// 			2:  6,
// 			3:  10,
// 			4:  4,
// 			5:  5,
// 			6:  3,
// 			7:  2,
// 			8:  3,
// 			10: 1,
// 		}}
// 	//HHrx
// 	h2 *H = &H{
// 		m:   35,
// 		max: 182,
// 		sk: map[int]int{
// 			122: 1,
// 			131: 2,
// 			132: 2,
// 			135: 1,
// 			137: 1,
// 			138: 3,
// 			139: 1,
// 			140: 1,
// 			141: 1,
// 			142: 2,
// 			144: 2,
// 			147: 1,
// 			148: 2,
// 			149: 1,
// 			150: 1,
// 			152: 1,
// 			154: 3,
// 			155: 1,
// 			157: 1,
// 			158: 1,
// 			159: 1,
// 			166: 1,
// 			167: 1,
// 			169: 1,
// 			170: 1,
// 			182: 1,
// 		}}
// )

// func TestHrx(t *testing.T) {

// p := Presun{
// 	Tab{10, 1},
// 	Tab{8, 3},
// 	Tab{7, 1},
// }
// t.Log(h.Simul(p))
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
