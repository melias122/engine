package komb

import (
	"math"
	"strings"

	"github.com/melias122/psl/num"
)

type K struct {
	n, m int
	nptr int
	ptr  []*num.N
	num  *num.N

	ntica Ntica
	xtica Tica
}

func New(n, m int) *K {
	return &K{
		n:   n,
		m:   m,
		ptr: make([]*num.N, n+1),
		num: num.Zero(n, m),

		ntica: newNtica(n),
		xtica: make(Tica, ((m-1)/10)+1),
	}
}

// func (k *K) Nil() {
// 	for i := range k.ptr {
// 		k.ptr[i] = nil
// 	}
// 	k.ptr = nil
// 	k.num = nil
// 	k.ntica.n = nil
// 	k.ntica.t = nil
// 	k.xtica = nil
// }

func (k *K) R1() float64 {
	return k.num.R1()
}

func (k *K) R2() float64 {
	return k.num.R2()
}

func (k *K) Sucet() int {
	return k.num.Cislo()
}

func (k *K) Xtica() Tica {
	return k.xtica
}

func (k *K) Ntica() Tica {
	return k.ntica.t
}

func (k0 *K) Zh(k1 *K) int {
	zh, i, j := 0, 0, 0
	for i < k0.Len() && j < k1.Len() {
		x1 := k0.At(i).Cislo()
		x2 := k1.At(j).Cislo()
		if x1 == x2 {
			zh++
			i++
			j++
		} else if x1 < x2 {
			i++
		} else {
			j++
		}
	}
	return zh
}

func (k0 *K) Kk(k1 *K) float64 {
	kk := 0.0

	for i := 0; i < k0.n; i++ {
		a := float64(k1.At(i).Cislo())
		p := float64(k0.At(i).Cislo())
		kk += math.Pow((a-p)/float64(k0.m), 4) / float64(k0.n)
	}

	return math.Pow(float64(1)-math.Sqrt(kk), 8)
}

func (k *K) Sm() float64 {
	var (
		sm  float64
		nSm int
	)

	for i := 0; i < k.n-1; i++ {
		for j := i + 1; j < k.n; j++ {
			p1 := float64(k.At(j).Cislo()) - float64(k.At(i).Cislo())
			p2 := float64(j) - float64(i)

			p1 /= float64(k.m - 1)
			p2 /= float64(k.n - 1)
			p1 /= p2

			sm += p1
			nSm++
		}
	}

	if nSm > 0 {
		sm /= float64(nSm)
	}
	return sm
}

func (k *K) C() num.C {
	return k.num.C()
}

func (k *K) Push(n *num.N) {
	k.num.Plus(n)
	k.ptr[k.nptr] = n
	k.xtica[(n.Cislo()-1)/10]++
	k.ntica.push(n.Cislo())
	k.nptr++
}

func (k *K) Pop() *num.N {
	k.nptr--
	n := k.ptr[k.nptr]
	k.num.Minus(n)
	k.xtica[(n.Cislo()-1)/10]--
	k.ntica.pop()
	return n
}

func (k *K) Len() int {
	return k.nptr
}

func (k *K) At(i int) *num.N {
	return k.ptr[i]
}

func (k *K) Contains(n *num.N) int {
	for _, c := range k.ptr {
		if n == c {
			return 1
		}
	}
	return 0
}

func (k *K) String() string {
	s := make([]string, k.nptr)
	for i := 0; i < k.nptr; i++ {
		s[i] = k.ptr[i].String()
	}
	return strings.Join(s, " ")
}
