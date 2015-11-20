package psl

import (
	"fmt"
	"math"
	"strconv"
)

func aritmetickyPriemer(s []float64) float64 {
	if len(s) == 0 {
		return .0
	}
	var (
		m   = make(map[float64]float64)
		sum float64
	)
	for _, f := range s {
		m[f] = 1.0
	}
	for f := range m {
		sum += f
	}
	return sum / float64(len(m))
}

func vazenyAritmetickyPriemer(s []float64) float64 {
	var (
		m          = make(map[float64]float64)
		sum1, sum2 float64
	)
	for _, f := range s {
		m[f]++
	}
	for k, v := range m {
		sum1 += k * v
		sum2 += v
	}
	if sum2 == .0 {
		return .0
	}
	return sum1 / sum2
}

func harmonickyPriemer(s []float64) float64 {
	var (
		n   float64
		sum float64
	)
	for _, f := range s {
		if f > .0 {
			sum += 1 / f
			n++
		}
	}
	if n > 0 {
		sum = n / sum
	}
	return sum
}

func geometrickyPriemer(s []float64) float64 {
	if len(s) == 0 {
		return 0
	}

	var p float64
	for _, n := range s {
		if p == 0 {
			p = n
		} else {
			p *= n
		}
	}
	return math.Pow(p, 1/float64(len(s)))
}

type statCifrovacky struct {
	nUdalost             map[int]int // pocet udalost
	diffSum, diffCnt     map[int]int
	diff                 map[int][]float64
	poslednyVyskytRiadok map[int]int
}

func makeStatCifrovacky(n, m int, r []Riadok, f func(r Riadok) []byte, tmax func() []byte) []*statCifrovacky {
	var (
		// tmax = komb.CifrovackyTeorMax(n, m)
		tMax = tmax()
		sc   = make([]*statCifrovacky, len(tMax))
	)
	// zratanie
	for _, r := range r {
		for si, u := range f(r) {
			if sc[si] == nil {
				sc[si] = &statCifrovacky{
					nUdalost:             make(map[int]int),
					diffSum:              make(map[int]int),
					diffCnt:              make(map[int]int),
					diff:                 make(map[int][]float64),
					poslednyVyskytRiadok: make(map[int]int),
				}
			}
			stat := sc[si]
			u := int(u)
			stat.nUdalost[u]++

			for i := 0; i < int(tMax[si]); i++ {
				if i == u {
					if stat.diffCnt[i] > 0 {
						stat.diff[i] = append(stat.diff[i], float64(stat.diffCnt[i]))
						stat.diffSum[i] += stat.diffCnt[i]
					}
					stat.diffCnt[i] = 0
				} else {
					stat.diffCnt[i]++
				}
			}
			stat.poslednyVyskytRiadok[u] = int(r.Pc)
		}
	}
	return sc
}

func statCifrovackyStrings(n, m, lenOrigHeader int, sc []*statCifrovacky, tmax func() []byte, doPad bool) [][]string {
	var (
		s [][]string
		// tmax       = komb.CifrovackyTeorMax(n, m)
		tMax       = tmax()
		maxUdalost int
	)
	for _, u := range tMax {
		u := int(u)
		if u > maxUdalost {
			maxUdalost = u
		}
	}

	for i := 0; i < maxUdalost; i++ {
		s = append(s,
			[]string{fmt.Sprintf("Pocet udalost %d", i)},
			[]string{fmt.Sprintf("Sucet diferencii krok udalost %d", i)},
			[]string{fmt.Sprintf("Krok aritmeticky priemer udalost %d", i)},
			[]string{fmt.Sprintf("Krok harmonicky priemer udalost %d", i)},
			[]string{fmt.Sprintf("Krok vazeny priemer udalost %d", i)},
			[]string{fmt.Sprintf("Krok geometricky priemer udalost %d", i)},
			[]string{fmt.Sprintf("Riadok posledny vyskyt udalost %d", i)},
			[]string{fmt.Sprintf("Krok aritmeticky priemer + riadok posledny vyskyt %d", i)},
			[]string{fmt.Sprintf("Krok harmonicky priemer + riadok posledny vyskyt %d", i)},
			[]string{fmt.Sprintf("Krok vazeny priemer + riadok posledny vyskyt %d", i)},
			[]string{fmt.Sprintf("Krok geometricky priemer + riadok posledny vyskyt %d", i)},
			[]string{},
			[]string{},
		)
	}
	if doPad {
		for i := range s {
			for len(s[i]) < lenOrigHeader {
				s[i] = append(s[i], "")
			}
		}
	}
	// var pad int
	// if doPad {
	pad := len(s) / maxUdalost
	// } else {
	// 	pad = 1
	// }
	for j, stat := range sc {
		for u := 0; u <= maxUdalost && u < int(tMax[j]); u++ {
			ap := int(aritmetickyPriemer(stat.diff[u]) + .5)
			hp := int(harmonickyPriemer(stat.diff[u]) + .5)
			vp := int(vazenyAritmetickyPriemer(stat.diff[u]) + .5)
			gp := int(geometrickyPriemer(stat.diff[u]) + .5)
			last := stat.poslednyVyskytRiadok[u]

			i := u * pad
			s[i] = append(s[i], itoa(stat.nUdalost[u]))
			s[i+1] = append(s[i+1], itoa(stat.diffSum[u]))
			s[i+2] = append(s[i+2], itoa(ap))
			s[i+3] = append(s[i+3], itoa(hp))
			s[i+4] = append(s[i+4], itoa(vp))
			s[i+5] = append(s[i+5], itoa(gp))
			s[i+6] = append(s[i+6], itoa(last))
			s[i+7] = append(s[i+7], itoa(ap+last))
			s[i+8] = append(s[i+8], itoa(hp+last))
			s[i+9] = append(s[i+9], itoa(vp+last))
		}
	}
	s = append([][]string{[]string{}}, s...)
	s = append([][]string{[]string{}}, s...)

	s = append([][]string{[]string{"Teoreticky možné"}}, s...)
	s = append([][]string{[]string{"Počet celkom"}}, s...)
	if doPad {
		for len(s[0]) < lenOrigHeader {
			s[0] = append(s[0], "")
			s[1] = append(s[1], "")
		}
	}
	for i := 0; i < len(tMax); i++ {
		var sum int
		for _, v := range sc[i].nUdalost {
			sum += v
		}
		s[0] = append(s[0], itoa(sum))
		s[1] = append(s[1], itoa(int(tMax[i])))
	}
	s = append([][]string{[]string{}}, s...)
	s = append([][]string{[]string{}}, s...)

	return s
}

func (a *Archiv) statistikaCifrovacky() error {
	header := a.origHeader
	for i := 1; i <= 10; i++ {
		header = append(header, "Cifra ("+strconv.Itoa(i%10)+")")
	}
	w := NewCsvMaxWriter(a.WorkingDir, "KrokCifrovacky", [][]string{header})
	defer w.Close()

	// dokumentacia
	for _, r := range a.riadky {
		line := []string{itoa(int(r.Pc))}
		line = append(line, r.origStrings[1:]...)
		line = append(line, r.Cifrovacky.Strings()...)
		if err := w.Write(line); err != nil {
			return err
		}
	}

	// statistika 1-do
	w.Write([]string{})
	w.Write([]string{})
	w.Write([]string{"R 1-DO"})

	f := func(r Riadok) []byte { return r.Cifrovacky[:] }
	tmax := func() []byte {
		max := CifrovackyTeorMax(a.n, a.m)
		return max[:]
	}
	stat1do := makeStatCifrovacky(a.n, a.m, a.riadky, f, tmax)
	stat1doStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), stat1do, tmax, true)
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
		statoddoStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), statoddo, tmax, true)
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
