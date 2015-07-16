package num

import "math/big"

var (
	bcache = make(map[bkey]big.Int)
	fcache = make(map[key]float64)
)

type bkey struct {
	a, b byte
}

type key struct {
	x, y, n, m byte
	pocet      uint32
}

// Value vracia hodnotu cisla x v stlpci y
// podla pocetnosti p.. n a m je rozmer databazy
// Vzorec: hodnota = pocetnostCisla / (binom(m-x nad n-y) * binom(x-1 nad y-1))
func Value(pocet, x, y, n, m int) float64 {
	return value(key{byte(x), byte(y), byte(n), byte(m), uint32(pocet)})
}

func value(k key) float64 {

	if v, ok := fcache[k]; ok {
		return v
	}
	var (
		c, d big.Int
		r    big.Rat
	)
	b, ok := bcache[bkey{k.m - k.x, k.n - k.y}]
	if !ok {
		b.Binomial(int64(k.m-k.x), int64(k.n-k.y))
		bcache[bkey{k.m - k.x, k.n - k.y}] = b
	}
	a, ok := bcache[bkey{k.x - 1, k.y - 1}]
	if !ok {
		a.Binomial(int64(k.x-1), int64(k.y-1))
		bcache[bkey{k.x - 1, k.y - 1}] = a
	}
	c.Mul(&b, &a)
	d.SetInt64(int64(0))
	if c.Cmp(&d) > 0 {
		d.SetInt64(int64(k.pocet))
		r.SetFrac(&d, &c)
		v, _ := r.Float64()

		fcache[k] = v
		return v
	} else {
		return 0
	}
}

// func vrati maximalnu teoreticku
// pocetnost cisla v stlpci
// func Max(x, y, n, m int) *big.Int {
// 	var a, b big.Int
// 	return a.Mul(a.Binomial(int64(m-x), int64(n-y)), b.Binomial(int64(x-1), int64(y-1)))
// }
