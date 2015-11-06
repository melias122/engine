package psl

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func (a *Archiv) statistikaZhoda() error {

	// statistika
	stat := struct {
		celkom map[int]int
		zh     map[int]map[string]int
	}{
		celkom: make(map[int]int),
		zh:     make(map[int]map[string]int),
	}
	for i := 0; i <= a.n; i++ {
		stat.zh[i] = make(map[string]int)
	}
	var zh0 *zhodaRiadok
	for _, r := range a.riadky {
		zh1 := makeZhodaRiadok(r.K, zh0)

		stat.celkom[zh1.zh]++
		pZH := []string{}
		for i, c := range zh1.presun {
			if c > 0 {
				s := strconv.Itoa(c) + "/" + strconv.Itoa(i+1)
				if len(s) > 0 {
					pZH = append(pZH, s)
				}
			}
		}
		stat.zh[zh1.zh][strings.Join(pZH, ", ")]++
		zh0 = zh1
	}
	//

	header := []string{"Zhoda", "Pocetnost teor.", "Teoreticka moznost v %", "Pocetnost", "Realne dosiahnute %"}

	w := NewCsvMaxWriter(a.WorkingDir, "StatistikaZhoda", [][]string{header})
	defer w.Close()

	dbLen := float64(len(a.riadky))
	for i := a.n; i >= 0; i-- {
		var (
			c, m big.Int
			r    big.Rat
		)
		c.Mul(c.Binomial(int64(a.n), int64(i)), m.Binomial(int64(a.m-a.n), int64(a.m+i-(2*a.n))))
		r.SetFrac(&c, m.Binomial(int64(a.m), int64(a.n)))
		f, _ := r.Float64()
		if err := w.Write([]string{
			itoa(i),
			c.String(),
			ftoa(f * 100),
			itoa(stat.celkom[i]),
			ftoa((float64(stat.celkom[i]) / dbLen) * 100),
		}); err != nil {
			return err
		}
	}

	var s [][]string
	for i := 1; i <= a.n; i++ {
		s = append(s,
			[]string{""},
			[]string{fmt.Sprintf("Zhoda %d", i), "Pocetnost", "Realne %"},
			[]string{fmt.Sprintf("Zhoda %d", i), itoa(stat.celkom[i]), ftoa((float64(stat.celkom[i]) / dbLen) * 100)},
		)
		for k, v := range stat.zh[i] {
			s = append(s, []string{
				k,
				itoa(v),
				ftoa((float64(v) / dbLen) * 100),
			})
		}
	}
	for _, r := range s {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	return nil
}

func (a *Archiv) statistikaNtice() error {
	stat := struct {
		teorMax map[string]*big.Int
		celkom  map[string]int
		sucin   map[string]map[string]int
	}{
		teorMax: make(map[string]*big.Int),
		celkom:  make(map[string]int),
		sucin:   make(map[string]map[string]int),
	}
	var (
		nticeVsetky = nticeStr(a.n)
		counter     = make(map[string]int)
	)
	for _, ntica := range nticeNtice(a.n) {
		var (
			k        = a.m - a.n + 1
			pocetMax = big.NewInt(1)
			b        big.Int
		)
		for _, n := range ntica {
			if n == 0 {
				continue
			}
			pocetMax.Mul(pocetMax, b.Binomial(int64(k), int64(k-int(n))))
			k -= int(n)
		}
		stat.teorMax[ntica.String()] = pocetMax
	}
	for _, tica := range nticeVsetky {
		counter[tica] = 0
	}
	for _, r := range a.riadky {
		ntica := r.Ntica.String()
		stat.celkom[ntica]++
		sucin := NticaSucin(r.K).String()
		if _, ok := stat.sucin[ntica]; !ok {
			stat.sucin[ntica] = make(map[string]int)
		}
		stat.sucin[ntica][sucin]++
	}

	var (
		dbLen = float64(len(a.riadky))
		s     [][]string
	)
	for _, ntica := range nticeVsetky {
		var r big.Rat
		r.SetFrac(stat.teorMax[ntica], big.NewInt(0).Binomial(int64(a.m), int64(a.n)))
		teorPercento, _ := r.Float64()
		s = append(s, []string{
			ntica,
			stat.teorMax[ntica].String(),                    // teor max pocet
			ftoa(teorPercento * 100),                        // teor percento
			itoa(stat.celkom[ntica]),                        // skutocny pocet za DB
			ftoa(float64(stat.celkom[ntica]) / dbLen * 100), // skutocne percento za DB
		})
	}
	s = append(s,
		[]string{""},
		[]string{
			"N-tica", "Sucin pozicie a stlpca", "Pocet vyskytov", "%",
		})
	for _, ntica := range nticeVsetky {
		s = append(s, []string{
			ntica,
			"vsetky",
			itoa(stat.celkom[ntica]),
			ftoa(float64(stat.celkom[ntica]) / dbLen * 100),
		},
		)
		for k, v := range stat.sucin[ntica] {
			s = append(s, []string{
				ntica,
				k,
				itoa(v),
				ftoa(float64(v) / dbLen * 100),
			})
		}
		s = append(s, []string{""})
	}
	header := []string{
		"N-tica", "Pocetnost teor.", "Teoreticka moznost v %",
		"Realne dosiahnuta pocetnost", "Realne dosiahnute %",
	}
	w := NewCsvMaxWriter(a.WorkingDir, "StatistikaNtice", [][]string{header})
	defer w.Close()

	for _, r := range s {
		if err := w.Write(r); err != nil {
			return err
		}
	}
	return nil
}

func (a *Archiv) statistikaCislovacky() error {

	f := func(r Riadok) []byte {
		b := r.C[:]
		b = append(b, byte(r.Zh))
		return b
	}

	header := []string{
		"", "", "", "", "", "", "", "",
		"P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC", "ZH",
	}
	w := NewCsvMaxWriter(a.WorkingDir, "StatistikaCislovacky", [][]string{header})
	defer w.Close()

	tmax := func() []byte {
		var c Cislovacky
		for i := 1; i <= a.m; i++ {
			cislovacky := NewCislovacky(i)
			c.Plus(cislovacky)
		}
		sc := c[:]
		sc = append(sc, byte(a.n-1))
		for i, n := range sc {
			if int(n) > a.n {
				sc[i] = byte(a.n)
			}
		}
		return sc
	}
	stat1do := makeStatCifrovacky(a.n, a.m, a.riadky, f, tmax)
	stat1doStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), stat1do, tmax)
	for _, s := range stat1doStrings {
		if err := w.Write(s); err != nil {
			return err
		}
	}

	// statistika od-do
	w.Write([]string{})
	w.Write([]string{})
	w.Write([]string{"R OD-DO"})
	if a.Uc.Cislo != 0 && a.Uc.Riadok > 1 {
		statoddo := makeStatCifrovacky(a.n, a.m, a.riadky[a.Uc.Riadok-1:], f, tmax)
		statoddoStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), statoddo, tmax)
		for _, s := range statoddoStrings {
			if err := w.Write(s); err != nil {
				return err
			}
		}
	} else {
		w.Write([]string{"Nenastala udalost 101..."})
	}

	return nil
}
