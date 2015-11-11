package psl

import (
	"bytes"
	"fmt"
	"math/big"
)

func (a *Archiv) PocetnostR() error {
	var header = []string{
		"Cislo", "ZH \"r\"", "P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC",
		"Cislo", "Teor. pocet", "Teor. %",
		"Pocet R1-DO", "% R1-DO", "Pocet R1-DO (r+1)", "% R1-DO (r+1)",
		"Pocet ROD-DO", "% ROD-DO", "Pocet ROD-DO (r+1)", "% ROD-DO (r+1)",
	}
	w := NewCsvMaxWriter(a.WorkingDir, "PocetnostR", [][]string{header})
	defer w.Close()
	var (
		max    = big.NewInt(0).Binomial(int64(a.m-1), int64(a.n-1)).String()
		riadok = make([]string, 0, len(header))
	)
	for i := 1; i <= a.m; i++ {
		riadok = riadok[:0]

		N1 := a.HHrx.GetNum(i)
		N2 := a.Hrx.GetNum(i)

		// Cislo
		riadok = append(riadok, N1.String())

		// Zhoda s r
		if bytes.Contains(a.K, []byte{byte(N1.Cislo())}) {
			riadok = append(riadok, "1")
		} else {
			riadok = append(riadok, "0")
		}

		// Cislovacky
		cislovacky := NewCislovacky(i)
		riadok = append(riadok, cislovacky.Strings()...)
		// for _, e := range num.NewC(i) {
		// 	riadok = append(riadok, itoa(int(e)))
		// }
		riadok = append(riadok,
			N1.String(), max, "1",
			itoa(N1.PocetR()),
			ftoa(N1.R()),
			itoa(N1.PocetR()+1),
			ftoa(Value(N1.PocetR()+1, 1, 1, a.n, a.m)),

			itoa(N2.PocetR()),
			ftoa(N2.R()),
			itoa(N2.PocetR()+1),
			ftoa(Value(N2.PocetR()+1, 1, 1, a.n, a.m)),
		)
		if err := w.Write(riadok); err != nil {
			return err
		}
	}
	return nil
}

func (a *Archiv) PocetnostS() error {
	var header = []string{
		"Cislo", "ZH \"r\"", "P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC", "Stlpec/Cislo",
		"Teor. pocet", "Teor. %", "Pocet STL1-DO", "% STL1-DO", "Pocet STL1-DO (r+1)", "% STL1-DO (r+1)",
		"Pocet STLOD-DO", "% STLOD-DO", "Pocet STLOD-DO (r+1)", "% STLOD-DO (r+1)",
	}

	w := NewCsvMaxWriter(a.WorkingDir, "PocetnostSTL", [][]string{header})
	defer w.Close()

	var (
		teorPocet, bi big.Int
	)
	for i := 1; i <= a.m; i++ {
		N1 := a.HHrx.GetNum(i)
		N2 := a.Hrx.GetNum(i)
		for j := 1; j <= a.n; j++ {

			r := make([]string, 0, len(header))

			// Cislo
			r = append(r, N1.String())

			// Zhoda s r
			// if bytes.Contains(a.K, []byte{byte(N1.Cislo())}) {
			if int(a.K[j-1]) == i {
				r = append(r, "1")
			} else {
				r = append(r, "0")
			}

			// Cislovacky
			cislovacky := NewCislovacky(i)
			r = append(r, cislovacky.Strings()...)
			// teorPoc := num.Max(i, j, a.n, a.m)

			teorPocet.Mul(teorPocet.Binomial(int64(a.m-i), int64(a.n-j)), bi.Binomial(int64(i-1), int64(j-1)))
			r = append(r,
				fmt.Sprintf("stlchce(%d):%d", j, i),
				teorPocet.String(),
			)

			var s1, s2, pocS1, pocS2 string
			if teorPocet.String() == "0" {
				r = append(r, "0")
				s1, s2, pocS1, pocS2 = ftoa(0.0), ftoa(0.0), "0", "0"
			} else {
				r = append(r, "1")
				s1, pocS1 = ftoa(Value(N1.PocetS(j)+1, i, j, a.n, a.m)), itoa(N1.PocetS(j)+1)
				s2, pocS2 = ftoa(Value(N2.PocetS(j)+1, i, j, a.n, a.m)), itoa(N2.PocetS(j)+1)
			}
			r = append(r,
				itoa(N1.PocetS(j)),
				ftoa(N1.S(j)),
				pocS1,
				s1,

				itoa(N2.PocetS(j)),
				ftoa(N2.S(j)),
				pocS2,
				s2,
			)
			if err := w.Write(r); err != nil {
				return err
			}
		}
	}
	// w.Flush()
	return nil
}
