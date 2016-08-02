package engine

import "math/big"

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
	var k0 Kombinacia
	for _, r := range a.riadky {
		zh := Zhoda(k0, r.K)
		stat.celkom[zh]++
		stat.zh[zh][NewZhodaPresun(k0, r.K).String()]++
		k0 = r.K
	}
	//

	header := []string{"Zhoda", "Pocetnost teor.", "Teoreticka moznost v %",
		"Pocetnost", "Realne dosiahnute %"}

	w := NewCsvMaxWriter("StatistikaZhoda", a.WorkingDir, setHeader(header))
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
			[]string{"Zhoda " + itoa(i), "Pocetnost", "Realne %"},
			[]string{
				"Zhoda " + itoa(i),
				itoa(stat.celkom[i]),
				ftoa((float64(stat.celkom[i]) / dbLen) * 100),
			},
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

func (a *Archiv) statistikaNtice2() error {
	stat := newStatistikaNtice(a.riadky, a.n, a.m)

	w := NewCsvMaxWriter("StatistikaNtice2", a.WorkingDir)
	defer w.Close()

	for _, s := range stat.strings() {
		if err := w.Write(s); err != nil {
			return err
		}
	}

	return nil
}

type statistikaNtica struct {
	// je dana ntica
	ntica Tica

	// max je teoreticke maximum ktore moze nadobudnut ntica
	max         *big.Int
	maxPercento float64

	// pocet vyskytov danej ntice v databaze
	pocet int

	// dblen je maximalny pocet ntic v databaze == dlzka databazy
	dblen int

	// krok opakovania
	krok []float64

	// riadok posledny vyskyt
	vyskyt int

	// nTyp predstavuje pocet ntice podla sucinu pozicie a stlpca
	nTyp map[string]*struct {
		pocet   int
		pozicie []byte
	}
}

func newStatistikaNtica(ntica Tica, dblen, n, m int) *statistikaNtica {
	var (
		k   = m - n + 1
		max = big.NewInt(1)
		b   big.Int
	)
	for _, ni := range ntica {
		if ni == 0 {
			continue
		}
		max.Mul(max, b.Binomial(int64(k), int64(k-int(ni))))
		k -= int(ni)
	}

	// Teor max %
	var r big.Rat
	r.SetFrac(max, b.Binomial(int64(m), int64(n)))
	f, _ := r.Float64()

	return &statistikaNtica{
		ntica:       ntica,
		max:         max,
		maxPercento: f * 100,
		dblen:       dblen,
		krok:        []float64{1},
		nTyp: make(map[string]*struct {
			pocet   int
			pozicie []byte
		}),
	}
}

func (s *statistikaNtica) incKrok() {
	s.krok[len(s.krok)-1]++
}

func (s *statistikaNtica) add(k Kombinacia, riadok int) {
	s.pocet++
	s.krok = append(s.krok, 1)

	// posledny vyskyt
	s.vyskyt = riadok

	// typ ntice
	id := NticaSucin(k).String()
	if v, ok := s.nTyp[id]; !ok {
		new := struct {
			pocet   int
			pozicie []byte
		}{
			pocet:   1,
			pozicie: NticaPozicie(k),
		}
		s.nTyp[id] = &new
	} else {
		v.pocet++
	}
}

func (sn *statistikaNtica) strings() [][]string {

	s := [][]string{
		{},
		{
			"", "Ntica", "Súčin pozície a stĺpca", "Počet výskytov", "%",
		},
	}

	addRow := func() {
		s = append(s, []string{})
	}
	addCol := func(str ...string) {
		i := len(s) - 1
		s[i] = append(s[i], str...)
	}

	for i := range sn.ntica {
		addCol("STL " + itoa(i+1))
	}

	addRow()
	addCol("", sn.ntica.String(), "vsetky", itoa(sn.pocet), ftoa(float64(sn.pocet)/float64(sn.dblen)*100))

	addRow()
	addCol("Počet udalosť", itoa(sn.pocet))

	addRow()
	addCol("Súčet diferencií krok udalosť", itoa(sn.dblen-sn.pocet))

	addRow()
	addCol("Krok aritmeticky priemer udalosť", itoa(int(aritmetickyPriemer(sn.krok))))

	addRow()
	addCol("Krok harmonický priemer udalosť", itoa(int(harmonickyPriemer(sn.krok))))

	addRow()
	addCol("Krok vážený priemer udalosť", itoa(int(vazenyAritmetickyPriemer(sn.krok))))

	addRow()
	addCol("Krok geometrický priemer udalosť", itoa(int(geometrickyPriemer(sn.krok))))

	addRow()
	addCol("Riadok posledný výskyt udalosť", itoa(sn.vyskyt))

	addRow()
	addCol("Krok aritmeticky priemer udalosť + riadok posledný výskyt", itoa(sn.vyskyt+int(aritmetickyPriemer(sn.krok))))

	addRow()
	addCol("Krok harmonický priemer udalosť + riadok posledný výskyt", itoa(sn.vyskyt+int(harmonickyPriemer(sn.krok))))

	addRow()
	addCol("Krok vážený priemer udalosť + riadok posledný výskyt", itoa(sn.vyskyt+int(vazenyAritmetickyPriemer(sn.krok))))

	addRow()
	addCol("Krok geometrický priemer udalosť + riadok posledný výskyt", itoa(sn.vyskyt+int(geometrickyPriemer(sn.krok))))

	ntica := sn.ntica.String()
	for ss, v := range sn.nTyp {
		addRow()
		addCol("", ntica, ss, itoa(v.pocet), ftoa(float64(v.pocet)/float64(sn.dblen)*100))
		for _, i := range v.pozicie {
			if i == 1 {
				addCol("X")
			} else {
				addCol("")
			}
		}
	}

	return s
}

type statistikaNtice struct {
	n []*statistikaNtica
}

func newStatistikaNtice(r []Riadok, n, m int) statistikaNtice {
	s := statistikaNtice{
	// n: make([]*statistikaNtica),
	}
	for _, tica := range nticeNtice(n) {
		s.n = append(s.n, newStatistikaNtica(tica, len(r), n, m))
	}
	for _, r := range r {
		s.add(r)
	}
	return s
}

func (s *statistikaNtice) add(r Riadok) {
	for _, sn := range s.n {
		if sn.ntica.Equal(r.Ntica) {
			sn.add(r.K, int(r.Pc))
		} else {
			sn.incKrok()
		}
	}
}

func (sn *statistikaNtice) strings() [][]string {
	s := [][]string{
		{
			"", "Ntica", "Teoretické maximum", "Teoretické %", "Počet výskytov", "%",
		},
	}
	for _, n := range sn.n {
		s = append(s, []string{""})
		i := len(s) - 1
		s[i] = append(s[i],
			n.ntica.String(),
			n.max.String(),
			ftoa(n.maxPercento),
			itoa(n.pocet),
			ftoa(float64(n.pocet)/float64(n.dblen)*100),
		)
	}

	for _, n := range sn.n {
		s = append(s, n.strings()...)
	}

	return s
}

// func (a *Archiv) statistikaNtice2() error {
// 	stat := struct {
// 		teorMax map[string]*big.Int
// 		celkom  map[string]int
// 		sucin   map[string]map[string]int
// 	}{
// 		teorMax: make(map[string]*big.Int),
// 		celkom:  make(map[string]int),
// 		sucin:   make(map[string]map[string]int),
// 	}
// 	var (
// 		nticeVsetky = nticeStr(a.n)
// 		counter     = make(map[string]int)
// 	)
// 	for _, ntica := range nticeNtice(a.n) {
// 		var (
// 			k        = a.m - a.n + 1
// 			pocetMax = big.NewInt(1)
// 			b        big.Int
// 		)
// 		for _, n := range ntica {
// 			if n == 0 {
// 				continue
// 			}
// 			pocetMax.Mul(pocetMax, b.Binomial(int64(k), int64(k-int(n))))
// 			k -= int(n)
// 		}
// 		stat.teorMax[ntica.String()] = pocetMax
// 	}
// 	for _, tica := range nticeVsetky {
// 		counter[tica] = 0
// 	}
// 	for _, r := range a.riadky {
// 		ntica := r.Ntica.String()
// 		stat.celkom[ntica]++
// 		sucin := NticaSucin(r.K).String()
// 		if _, ok := stat.sucin[ntica]; !ok {
// 			stat.sucin[ntica] = make(map[string]int)
// 		}
// 		stat.sucin[ntica][sucin]++
// 	}

// 	var (
// 		dbLen = float64(len(a.riadky))
// 		s     [][]string
// 	)
// 	for _, ntica := range nticeVsetky {
// 		var r big.Rat
// 		r.SetFrac(stat.teorMax[ntica], big.NewInt(0).Binomial(int64(a.m), int64(a.n)))
// 		teorPercento, _ := r.Float64()
// 		s = append(s, []string{
// 			ntica,
// 			stat.teorMax[ntica].String(),                    // teor max pocet
// 			ftoa(teorPercento * 100),                        // teor percento
// 			itoa(stat.celkom[ntica]),                        // skutocny pocet za DB
// 			ftoa(float64(stat.celkom[ntica]) / dbLen * 100), // skutocne percento za DB
// 		})
// 	}
// 	s = append(s,
// 		[]string{""},
// 		)
// 	for _, ntica := range nticeVsetky {
// 		s = append(s, []string{
// 			ntica,
// 			"vsetky",
// 			itoa(stat.celkom[ntica]),
// 			ftoa(float64(stat.celkom[ntica]) / dbLen * 100),
// 		},
// 		)
// 		for k, v := range stat.sucin[ntica] {
// 			s = append(s, []string{
// 				ntica,
// 				k,
// 				itoa(v),
// 				ftoa(float64(v) / dbLen * 100),
// 			})
// 		}
// 		s = append(s, []string{""})
// 	}
// 	header := []string{
// 		"N-tica", "Pocetnost teor.", "Teoreticka moznost v %",
// 		"Realne dosiahnuta pocetnost", "Realne dosiahnute %",
// 	}
// 	w := NewCsvMaxWriter("StatistikaNtice", a.WorkingDir, setHeader(header))
// 	defer w.Close()

// 	for _, r := range s {
// 		if err := w.Write(r); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (a *Archiv) statistikaCislovacky() error {

	f := func(r Riadok) []byte {
		b := r.C[:]
		b = append(b, byte(r.Zh))
		return b
	}

	header := []string{
		"", "P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC", "ZH",
	}
	w := NewCsvMaxWriter("StatistikaCislovacky", a.WorkingDir, setHeader(header))
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
	stat1doStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), stat1do, tmax, false)
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
		statoddoStrings := statCifrovackyStrings(a.n, a.m, len(a.origHeader), statoddo, tmax, false)
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
