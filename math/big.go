package math

import (
	"math/big"
	"sync"
)

type key struct {
	p, x, y int
}

type c struct {
	sync.RWMutex
	c map[key]float64
}

func (c *c) get(k key) (v float64, ok bool) {
	c.RLock()
	v, ok = c.c[k]
	c.RUnlock()
	return
}

func (c *c) put(k key, v float64) {
	c.Lock()
	c.c[k] = v
	c.Unlock()
}

var (
	cache = c{c: make(map[key]float64, 2048)}
	zero  = big.NewInt(int64(0))
)

func Value(p, x, y, n, m int) float64 {
	if x > (m/2)+m%2 {
		x = (x - (m + 1)) * (-1)
		y = n - y + 1
	}
	k := key{p, x, y}
	v, ok := cache.get(k)
	if !ok {
		v = value(k, n, m)
		cache.put(k, v)
	}
	return v
}

func value(k key, n, m int) float64 {
	b := big.NewInt(int64(k.p))
	c := Max(k.x, k.y, n, m)
	if c.Cmp(zero) == 0 {
		return 0.0
	}
	var a big.Rat
	a.SetFrac(b, c)
	f, _ := a.Float64()
	return f
}

// func vrati maximalnu teoreticku
// pocetnost cisla v stlpci
func Max(x, y, n, m int) *big.Int {
	a := new(big.Int)
	return a.Mul(Binomial(n-y, m-x), Binomial(y-1, x-1))
}

// func vrati nCm
// nCm(5,35) = 324632
func Binomial(n, m int) *big.Int {
	x := new(big.Int)
	return x.Binomial(int64(m), int64(n))
}
