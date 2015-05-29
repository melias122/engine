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

	ntica Tica
	xtica Tica
}

func New(n, m int) *K {
	return &K{
		n:   n,
		m:   m,
		ptr: make([]*num.N, n+1),
		num: num.Zero(n, m),

		ntica: make(Tica, n),
		xtica: make(Tica, ((m-1)/10)+1),
	}
}

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

// TODO: priadat do push
func (k *K) Ntica() Tica {
	// var tica int
	// k.ntica = make(Tica, k.n)
	// for i := 1; i < k.Len(); i++ {
	// 	if (k.At(i-1).Cislo() - k.At(i).Cislo()) == 1 {
	// 		tica++
	// 	} else if tica > 0 {
	// 		k.ntica[tica]++
	// 		tica = 0
	// 	} else {
	// 		k.ntica[0]++
	// 	}
	// }
	return k.ntica
}

func (k0 *K) Zh(k1 *K) int {
	zh, i, j := 0, 0, 0
	for i < k0.Len() && j < k1.Len() {
		if k0.ptr[i] == k1.ptr[i] {
			zh++
			i++
			j++
		} else if k0.At(i).Cislo() < k1.At(j).Cislo() {
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
	k.nptr++
}

func (k *K) Pop() *num.N {
	k.nptr--
	n := k.ptr[k.nptr]
	k.num.Minus(n)
	k.xtica[(n.Cislo()-1)/10]--
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
