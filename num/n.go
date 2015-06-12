package num

import (
	"math/big"
	"strconv"
)

type N struct {
	c C
	x int32
	r [2]*ph
	s [][]*ph
}

func Zero(n, m int) *N {
	var num N
	num.s = make([][]*ph, 2)
	for i := 0; i < 2; i++ {
		num.r[i] = &ph{}
		num.s[i] = make([]*ph, n)
		for j := 0; j < n; j++ {
			num.s[i][j] = &ph{}
		}

	}
	return &num
}

func New(x, n, m int) *N {
	num := N{
		c: newC(x),
		x: int32(x),
	}
	num.s = make([][]*ph, 2)
	for i := 0; i < 2; i++ {
		num.r[i] = newph(1, 1, n, m)
		num.s[i] = make([]*ph, n)
		for j := 1; j <= n; j++ {
			num.s[i][j-1] = newph(x, j, n, m)
		}

	}
	return &num
}

func (old *N) Copy(n, m int) *N {
	new := Zero(n, m)
	new.Plus(old)
	return new
}

func (n *N) C() C {
	return n.c
}

func (n *N) Cislo() int {
	return int(n.x)
}

func (n *N) Plus(m *N) {
	n.c.Plus(m.c)
	n.x += m.x
	for i := 0; i < 2; i++ {
		n.r[i].plus(m.r[i])
		for j := 0; j < len(n.s[i]); j++ {
			n.s[i][j].plus(m.s[i][j])
		}
	}
	// return n
}

func (n *N) Minus(m *N) {
	n.c.Minus(m.c)
	n.x -= m.x
	for i := 0; i < 2; i++ {
		n.r[i].minus(m.r[i])
		for j := 0; j < len(n.s[i]); j++ {
			n.s[i][j].minus(m.s[i][j])
		}
	}
	// return n
}

func (n *N) Inc1(y int) {
	n.r[0].inc()
	n.s[0][y].inc()
}

func (n *N) Inc2(y int) {
	n.r[1].inc()
	n.s[1][y].inc()
}

func (n *N) Reset2() {
	n.r[1].reset()
	for _, ph := range n.s[1] {
		ph.reset()
	}
}

func (n *N) R1() float64 {
	return n.r[0].Hodnota()
}

func (n *N) R2() float64 {
	return n.r[1].Hodnota()
}

func (n *N) S1(y int) float64 {
	return n.s[0][y-1].Hodnota()
}

func (n *N) S2(y int) float64 {
	return n.s[1][y-1].Hodnota()
}

func (n *N) PocR1() int {
	return n.r[0].Pocet()
}

func (n *N) PocR2() int {
	return n.r[1].Pocet()
}

func (n *N) PocS1(y int) int {
	return n.s[0][y-1].Pocet()
}

func (n *N) PocS2(y int) int {
	return n.s[1][y-1].Pocet()
}

func (n *N) String() string {
	return strconv.Itoa(n.Cislo())
}

// func vrati maximalnu teoreticku
// pocetnost cisla v stlpci
func Max(x, y, n, m int) *big.Int {
	var a, b big.Int
	return a.Mul(a.Binomial(int64(m-x), int64(n-y)), b.Binomial(int64(x-1), int64(y-1)))
}
