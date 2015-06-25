package archiv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/melias122/psl/num"
)

func (a *Archiv) PocetnostR() error {
	var header = []string{
		"Cislo", "ZH \"r\"", "P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC",
		"Cislo", "Teor. pocet", "Teor. %",
		"Pocet R1-DO", "% R1-DO", "Pocet R1-DO (r+1)", "% R1-DO (r+1)",
		"Pocet ROD-DO", "% ROD-DO", "Pocet ROD-DO (r+1)", "% ROD-DO (r+1)",
	}
	f, err := os.Create(filepath.Join(a.Dir, "PocetnostR_"+a.Dir+".csv"))
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = ';'
	if err := w.Write(header); err != nil {
		return err
	}
	var (
		max    = big.NewInt(0).Binomial(int64(a.m-1), int64(a.n-1)).String()
		riadok = make([]string, 0, len(header))
	)
	for i := 1; i <= a.m; i++ {
		riadok = riadok[:0]

		N1 := a.HHrx.GetN(i)
		N2 := a.Hrx.GetN(i)

		// Cislo
		riadok = append(riadok, N1.String())

		// Zhoda s r
		if bytes.Contains(a.K, []byte{byte(N1.Cislo())}) {
			riadok = append(riadok, "1")
		} else {
			riadok = append(riadok, "0")
		}

		// Cislovacky
		for _, e := range N1.C() {
			riadok = append(riadok, itoa(int(e)))
		}
		riadok = append(riadok,
			N1.String(), max, "1",
			itoa(N1.PocetR()),
			ftoa(N1.R()),
			itoa(N1.PocetR()+1),
			ftoa(num.Value(N1.PocetR()+1, 1, 1, a.n, a.m)),

			itoa(N2.PocetR()),
			ftoa(N2.R()),
			itoa(N2.PocetR()+1),
			ftoa(num.Value(N2.PocetR()+1, 1, 1, a.n, a.m)),
		)
		if err := w.Write(riadok); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func (a *Archiv) PocetnostS() error {
	var header = []string{
		"Cislo", "ZH \"r\"", "P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC", "Stlpec/Cislo",
		"Teor. pocet", "Teor. %", "Pocet STL1-DO", "% STL1-DO", "Pocet STL1-DO (r+1)", "% STL1-DO (r+1)",
		"Pocet STLOD-DO", "% STLOD-DO", "Pocet STLOD-DO (r+1)", "% STLOD-DO (r+1)",
	}
	// fmt.Sprintf("%d%d/PocetnostSTL_%d%d.csv", a.n, a.m, a.n, a.m)
	f, err := os.Create(filepath.Join(a.Dir, "PocetnostS_"+a.Dir+".csv"))
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = ';'
	if err := w.Write(header); err != nil {
		return err
	}

	var (
		teorPocet, bi big.Int
	)
	for i := 1; i <= a.m; i++ {
		N1 := a.HHrx.GetN(i)
		N2 := a.Hrx.GetN(i)
		for j := 1; j <= a.n; j++ {

			r := make([]string, 0, len(header))

			// Cislo
			r = append(r, N1.String())

			// Zhoda s r
			if bytes.Contains(a.K, []byte{byte(N1.Cislo())}) {
				r = append(r, "1")
			} else {
				r = append(r, "0")
			}

			// Cislovacky
			for _, e := range N1.C() {
				r = append(r, itoa(int(e)))
			}
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
				s1, pocS1 = ftoa(num.Value(N1.PocetS(j)+1, i, j, a.n, a.m)), itoa(N1.PocetS(j)+1)
				s2, pocS2 = ftoa(num.Value(N2.PocetS(j)+1, i, j, a.n, a.m)), itoa(N2.PocetS(j)+1)
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
	w.Flush()
	return w.Error()
}
