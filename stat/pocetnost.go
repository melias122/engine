package stat

import (
	"math/big"

	"github.com/melias122/engine/engine"
)

type ph struct {
	Pocet   int
	Hodnota float64
}

type RiadokR struct {
	Cislo int
	Zhoda int
	engine.Cislovacka
	MaxPocet *big.Int
	R1       ph
	R2       ph
}

//	var header = []string{
//		"Cislo", "ZH \"r\"", "P", "N", "PR", "Mc", "Vc", "C19", "C0", "cC", "Cc", "CC",
//		"Cislo", "Teor. pocet", "Teor. max Ʃ",
//		"Pocet R 1-DO", "R 1-DO", "Pocet R 1-DO (r+1)", "R 1-DO (r+1)",
//		"Pocet R OD-DO", "R OD-DO", "Pocet R OD-DO (r+1)", "R OD-DO (r+1)",
//	}

func PocetnostR(last engine.Kombinacia, hhrx engine.Rc, hrx engine.Rc, n, m int) []RiadokR {
	var (
		maxPocet = big.NewInt(0).Binomial(int64(m-1), int64(n-1))
		r        []RiadokR
	)
	for i := 1; i <= m; i++ {
		r = append(r, RiadokR{
			Cislo:      i,
			Zhoda:      engine.Zhoda(last, engine.Kombinacia{i}),
			Cislovacka: engine.NewCislovacka(engine.Kombinacia{i}),
			MaxPocet:   maxPocet,
			R1: ph{
				Pocet:   hhrx.Rp(i),
				Hodnota: hhrx.Rh(i),
			},
			R2: ph{
				Pocet:   hrx.Rp(i),
				Hodnota: hrx.Rh(i),
			},
		})
	}
	return r
}

type RiadokS struct {
	Cislo  int
	Stlpec int
	Zhoda  int
	engine.Cislovacka
	MaxPocet *big.Int
	MaxSucet float64
	S1       ph
	S2       ph
}

// 	var header = []string{
// 		"Cislo", "ZH \"r\"", "P", "N", "PR", "Mc", "Vc", "C19", "C0", "cC", "Cc", "CC", "Stlpec/Cislo",
// 		"Teor. pocet", "Teor. max Ʃ", "Pocet STL 1-DO", "STL 1-DO", "Pocet STL 1-DO (r+1)", "STL 1-DO (r+1)",
// 		"Pocet STL OD-DO", "STL OD-DO", "Pocet STL OD-DO (r+1)", "STL OD-DO (r+1)",
// 	}

func PocetnostS(last engine.Kombinacia, hhrx engine.STLc, hrx engine.STLc, n, m int) []RiadokS {
	var r []RiadokS
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//			var a, maxp big.Int
			r = append(r, RiadokS{
				Cislo:      i,
				Stlpec:     j,
				Zhoda:      engine.Zhoda(last, engine.Kombinacia{i}),
				Cislovacka: engine.NewCislovacka(engine.Kombinacia{i}),
				//				MaxPocet:   maxp.Mul(maxp.Binomial(int64(m-i), int64(n-j)), a.Binomial(int64(i-1), int64(j-1))),
				//				MaxSucet:   engine.CalculateRS(i,j,),
				S1: ph{
					Pocet:   hhrx.STLp(i, j),
					Hodnota: hhrx.STLh(i, j),
				},
				S2: ph{
					Pocet:   hrx.STLp(i, j),
					Hodnota: hrx.STLh(i, j),
				},
			})
		}
	}
	return r
}
