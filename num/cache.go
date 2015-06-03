package num

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

func (c *c) get(k key) (float64, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.c[k]
	return v, ok
}

func (c *c) put(k key, v float64) {
	c.Lock()
	c.c[k] = v
	c.Unlock()
}

var (
	cache = c{c: make(map[key]float64, 1024)}
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
	var a, b big.Int
	b.Mul(b.Binomial(int64(m-k.x), int64(n-k.y)), a.Binomial(int64(k.x-1), int64(k.y-1)))
	if b.Cmp(zero) == 0 {
		return 0.0
	}
	a.SetInt64(int64(k.p))
	var rat big.Rat
	rat.SetFrac(&a, &b)
	f, _ := rat.Float64()
	return f
}

// func vrati maximalnu teoreticku
// pocetnost cisla v stlpci
func Max(x, y, n, m int) *big.Int {
	var a, b big.Int
	return a.Mul(a.Binomial(int64(m-x), int64(n-y)), b.Binomial(int64(x-1), int64(y-1)))
}
