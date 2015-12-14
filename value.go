package psl

import "math/big"

var (
	fcache = make(map[key]float64, 2048)
)

type key struct {
	x, y, n, m byte
	pocet      uint32
}

// Value vracia hodnotu cisla x v stlpci y
// podla pocetnosti p.. n a m je rozmer databazy
// Vzorec: hodnota = pocetnostCisla / (binom(m-x nad n-y) * binom(x-1 nad y-1))
func Value(pocet, x, y, n, m int) float64 {
	// Kvoli symetrickosti binomickych cisiel
	// maju cisla na poziciach rovnaku pocetnost
	// Priklad db n=5, m=35..
	// cislo 1 == 35, 2 == 34 ...
	// stlpec 1 == 5, 2 == 4 ...
	if x > (m/2)+m%2 {
		x = (x - (m + 1)) * (-1)
		y = n - y + 1
	}
	var key = key{byte(x), byte(y), byte(n), byte(m), uint32(pocet)}
	v, ok := fcache[key]
	if !ok {
		v = value(pocet, x, y, n, m)
		fcache[key] = v
		return v
	} else {
		return v
	}
}

func value(pocet, x, y, n, m int) float64 {
	var (
		a, b  big.Int
		value float64
	)
	a.Binomial(int64(x-1), int64(y-1))
	b.Binomial(int64(m-x), int64(n-y))

	if a.Int64() != 0 && b.Int64() != 0 {
		var r big.Rat
		a.Mul(&a, &b)
		b.SetInt64(int64(pocet))
		value, _ = r.SetFrac(&b, &a).Float64()
	}
	return value
}

// func vrati maximalnu teoreticku
// pocetnost cisla v stlpci
// func Max(x, y, n, m int) *big.Int {
// 	var a, b big.Int
// 	return a.Mul(a.Binomial(int64(m-x), int64(n-y)), b.Binomial(int64(x-1), int64(y-1)))
// }
