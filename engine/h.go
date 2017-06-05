package engine

import (
	"math/big"
)

var zero = big.NewInt(0)

// H vypocita hodnotu R a STL.
// Pre vypocet R je pozicia konstantne 1.
func H(cislo, pozicia, pocet, n, m int) float64 {
	var (
		a, b big.Int
		r    big.Rat
	)
	a.Binomial(int64(cislo-1), int64(pozicia-1))
	b.Binomial(int64(m-cislo), int64(n-pozicia))
	a.Mul(&a, &b)
	b.SetInt64(int64(pocet))
	if a.Cmp(zero) == 0 {
		return 0
	}
	hodnota, _ := r.SetFrac(&b, &a).Float64()
	return hodnota
}

// func vrati maximalnu teoreticku
// pocetnost cisla v stlpci
// func Max(x, y, n, m int) *big.Int {
// 	var a, b big.Int
// 	return a.Mul(a.Binomial(int64(m-x), int64(n-y)), b.Binomial(int64(x-1), int64(y-1)))
// }
