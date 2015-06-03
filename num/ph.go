package num

import (
	"math/big"
	"sync"
)

type IncFn func(int) float64

type ph struct {
	p  int
	h  float64
	fn IncFn
}

func newph(x, y, n, m int) *ph {
	return &ph{
		fn: func(p int) float64 { return Value(p, x, y, n, m) },
	}
}

func (p *ph) inc() {
	p.p++
	p.h = p.fn(p.p)
}

func (p *ph) reset() {
	p.p = 0
	p.h = 0
}

func (p *ph) plus(r *ph) *ph {
	p.h += r.h
	return p
}

func (p *ph) minus(r *ph) *ph {
	p.h -= r.h
	return p
}

type key struct {
	p, x, y int
}

type cache struct {
	sync.RWMutex
	data map[key]float64
}

func (c *cache) get(k key) (float64, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.data[k]
	return v, ok
}

func (c *cache) put(k key, v float64) {
	c.Lock()
	c.data[k] = v
	c.Unlock()
}

var (
	c    = &cache{data: make(map[key]float64, 1024)}
	zero = big.NewInt(int64(0))
)

func Value(p, x, y, n, m int) float64 {
	if x > (m/2)+m%2 {
		x = (x - (m + 1)) * (-1)
		y = n - y + 1
	}
	k := key{p, x, y}
	v, ok := c.get(k)
	if !ok {
		v = value(k, n, m)
		c.put(k, v)
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
