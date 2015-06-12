package num

import (
	"math/big"
	"sync"
)

type ph struct {
	x, y    byte
	n, m    byte
	pocet   uint16
	hodnota float64
}

func newph(x, y, n, m int) *ph {
	return &ph{
		x: byte(x),
		y: byte(y),
		n: byte(n),
		m: byte(m),
	}
}

func (p *ph) inc() {
	p.pocet++
	p.hodnota = Value(int(p.pocet), int(p.x), int(p.y), int(p.n), int(p.m))
}

func (p *ph) reset() {
	p.pocet = 0
	p.hodnota = 0
}

func (p *ph) plus(r *ph) {
	p.hodnota += r.hodnota
	// return p
}

func (p *ph) minus(r *ph) {
	p.hodnota -= r.hodnota
	// return p
}

func (p ph) Pocet() int {
	return int(p.pocet)
}

func (p ph) Hodnota() float64 {
	return p.hodnota
}

type key struct {
	p uint16
	x byte
	y byte
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
	c    = cache{data: make(map[key]float64, 1024)}
	zero = big.NewInt(int64(0))
)

// Value vracia hodnotu cisla x v stlpci y
// podla pocetnosti p.. n a m je rozmer databazy
func Value(p, x, y, n, m int) float64 {
	// Kvoli symetrickosti binomickych cisiel
	// maju cisla na poziciach rovnaku pocetnost
	// Priklad db n=5, m=35..
	// cislo 1 == 35, 2 == 34 ...
	// stlpec 1 == 5, 2 == 4 ...
	if x > (m/2)+m%2 {
		x = (x - (m + 1)) * (-1)
		y = n - y + 1
	}
	// vytvorime kluc a pozrieme sa predtym do cache
	// ci sme uz dane cislo v stlpci s pocetnostou nevyratali
	k := key{uint16(p), byte(x), byte(y)}
	v, ok := c.get(k)
	if !ok {
		// ak cislo nie je v cache ratame jeho hodnotu
		// a vlozime ho do cache pre dalsie pouzitie
		// Vzorec: hodnota = pocetnostCisla / (binom(m-x nad n-y) * binom(x-1 nad y-1))
		var a, b big.Int
		b.Mul(
			b.Binomial(int64(m-x), int64(n-y)),
			a.Binomial(int64(x-1), int64(y-1)),
		)
		if b.Cmp(zero) == 0 {
			v = 0.0
		} else {
			var rat big.Rat
			a.SetInt64(int64(k.p))
			rat.SetFrac(&a, &b) // a/b
			v, _ = rat.Float64()
		}
		c.put(k, v)
	}
	return v
}
