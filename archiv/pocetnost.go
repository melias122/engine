package archiv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/melias122/psl/math"
)

func (a *Archiv) PocetnostR() error {
	var header = []string{
		"Cislo", "ZH \"r\"", "P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC",
		"Cislo", "Teor. pocet", "Teor. %",
		"Pocet R1-DO", "% R1-DO", "Pocet R1-DO (r+1)", "% R1-DO (r+1)",
		"Pocet ROD-DO", "% ROD-DO", "Pocet ROD-DO (r+1)", "% ROD-DO (r+1)",
	}
	f, err := os.Create(fmt.Sprintf("%d%d/PocetnostR_%d%d.csv", a.n, a.m, a.n, a.m))
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = ';'
	if err := w.Write(header); err != nil {
		return err
	}
	max := math.Max(1, 1, a.n, a.m).String()
	for i := 1; i <= a.m; i++ {
		c := a.Cisla[i]
		r := make([]string, 0, len(header))
		r = append(r, c.String(), itoa(a.K.Contains(c)))
		for _, e := range c.C() {
			r = append(r, itoa(int(e)))
		}
		r = append(r,
			c.String(), max, "1",
			itoa(c.PocR1()), ftoa(c.R1()), itoa(c.PocR1()+1), ftoa(math.Value(c.PocR1()+1, 1, 1, a.n, a.m)),
			itoa(c.PocR2()), ftoa(c.R2()), itoa(c.PocR2()+1), ftoa(math.Value(c.PocR2()+1, 1, 1, a.n, a.m)),
		)
		if err := w.Write(r); err != nil {
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
	f, err := os.Create(fmt.Sprintf("%d%d/PocetnostSTL_%d%d.csv", a.n, a.m, a.n, a.m))
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = ';'
	if err := w.Write(header); err != nil {
		return err
	}
	for i := 1; i <= a.m; i++ {
		for j := 1; j <= a.n; j++ {
			c := a.Cisla[i]

			r := make([]string, 0, len(header))
			r = append(r, c.String(), itoa(a.K.Contains(c)))
			for _, e := range c.C() {
				r = append(r, itoa(int(e)))
			}
			teorPoc := math.Max(i, j, a.n, a.m)
			r = append(r, fmt.Sprintf("stlchce(%d):%d", j, i), teorPoc.String())

			var s1, s2, pocS1, pocS2 string
			if teorPoc.String() == "0" {
				r = append(r, "0")
				s1, s2, pocS1, pocS2 = ftoa(0.0), ftoa(0.0), "0", "0"
			} else {
				r = append(r, "1")
				s1, pocS1 = ftoa(math.Value(c.PocS1(j)+1, i, j, a.n, a.m)), itoa(c.PocS1(j)+1)
				s2, pocS2 = ftoa(math.Value(c.PocS2(j)+1, i, j, a.n, a.m)), itoa(c.PocS2(j)+1)
			}
			r = append(r,
				itoa(c.PocS1(j)), ftoa(c.S1(j)), pocS1, s1,
				itoa(c.PocS2(j)), ftoa(c.S2(j)), pocS2, s2,
			)
			if err := w.Write(r); err != nil {
				return err
			}
		}
	}
	w.Flush()
	return w.Error()
}
