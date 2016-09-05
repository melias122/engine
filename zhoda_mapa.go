package engine

import "gitlab.com/melias122/engine/csv"

func (a *Archiv) mapaZhoda2() error {
	mz := newMapaZhoda(a.riadky, a.n)

	w := csv.NewCsvMaxWriter("MapaZhoda2", a.Workdir, csv.SetHeader(mz.header()))
	defer w.Close()

	for _, s := range mz.strings() {
		if err := w.Write(s); err != nil {
			return err
		}
	}

	return nil
}

type statistikaZhoda struct {
	// maximalny pocet udalosti
	max int
	// udalost v danom stlpci (1..n)
	// oznacuje celkovy pocet udalosti
	// ktore nastali
	pocet []int

	// posledny vyskyt
	vyskyt []int

	// kroky danej udalosti
	krok [][]float64
}

func newStatistikaZhoda(max, n int) *statistikaZhoda {
	s := &statistikaZhoda{
		max:    max,
		pocet:  make([]int, n),
		vyskyt: make([]int, n),
		krok:   make([][]float64, n),
	}
	for i := range s.krok {
		s.krok[i] = []float64{1}
	}
	return s
}

func (s *statistikaZhoda) add(vyskyt int, udalost []int) {
	// oznacenie udalosti ktore nastali
	// ak udalost nastala zvysim pocet
	// a pridam krok
	// inak zvysim krok
	flag := make([]bool, len(s.pocet))
	for _, i := range udalost {
		flag[i-1] = true
	}
	for i, ok := range flag {
		if ok {
			s.pocet[i]++
			s.krok[i] = append(s.krok[i], 1)
			s.vyskyt[i] = vyskyt
		} else {
			k := s.krok[i]
			k[len(k)-1]++
		}
	}
}

func (sz *statistikaZhoda) strings() [][]string {

	var s [][]string

	addRow := func() {
		s = append(s, []string{})
	}
	addCol := func(str string) {
		i := len(s) - 1
		s[i] = append(s[i], str)
	}

	addRow()
	addCol("")
	for i := range sz.pocet {
		addCol("Udalost " + itoa(i+1))
	}

	addRow()
	addCol("Počet udalosť u")
	for _, u := range sz.pocet {
		addCol(itoa(u))
	}

	addRow()
	addCol("Súčet differencií krok udalosť")
	for _, u := range sz.pocet {
		addCol(itoa(sz.max - u))
	}

	addRow()
	addCol("Krok aritmeticky priemer udalosť")
	for _, k := range sz.krok {
		i := int(aritmetickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Krok harmonický priemer udalosť")
	for _, k := range sz.krok {
		i := int(harmonickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Krok vážený priemer udalosť")
	for _, k := range sz.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Krok geometrický priemer udalosť")
	for _, k := range sz.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Riadok posledný výskyt udalosť")
	for _, k := range sz.vyskyt {
		addCol(itoa(k))
	}

	addRow()
	addCol("Krok aritmeticky priemer udalosť + riadok posledný výskyt")
	for j, k := range sz.krok {
		i := int(aritmetickyPriemer(k))
		addCol(itoa(i + sz.vyskyt[j]))
	}

	addRow()
	addCol("Krok harmonický priemer udalosť + riadok posledný výskyt")
	for j, k := range sz.krok {
		i := int(harmonickyPriemer(k))
		addCol(itoa(i + sz.vyskyt[j]))
	}

	addRow()
	addCol("Krok vážený priemer udalosť + riadok posledný výskyt")
	for j, k := range sz.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i + sz.vyskyt[j]))
	}

	addRow()
	addCol("Krok geometrický priemer udalosť + riadok posledný výskyt")
	for j, k := range sz.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i + sz.vyskyt[j]))
	}
	return s
}

type mapaZhodaRiadok struct {
	k  Kombinacia
	zh int

	// oznacenie zhody s r+1 riadkom
	// lava strana 1..len(k)
	poz0 []int

	// oznacenie zhody v riaku
	// prava strana 1..len(k)
	poz []int

	zp ZhodaPresun

	krok int
}

func newMapaZhodaRiadok(pred, akt Kombinacia) *mapaZhodaRiadok {
	return &mapaZhodaRiadok{
		k:    akt,
		zh:   Zhoda(pred, akt),
		poz0: make([]int, akt.Len()),
		poz:  make([]int, akt.Len()),
		zp:   NewZhodaPresun(pred, akt),
	}
}

// ak nastala zhoda v r (aktualny riadok) pozicie nastavia sa
// pozicie presunov v danych stlpcoch v m (r-1) a r
func (m *mapaZhodaRiadok) merge(r *mapaZhodaRiadok) {
	if r.zh == 0 {
		return
	}

	for _, p := range r.zp.p {
		// citatel i
		i := p[0] - 1
		// menovatel j
		j := p[1] - 1

		// nastavenie cisiel
		m.poz0[i] = int(m.k[i])
		r.poz[j] = int(r.k[j])
	}
}

func (m *mapaZhodaRiadok) strings() []string {
	s := []string{
		m.k.String(),
		itoa(m.zh),
	}
	for _, i := range m.poz0 {
		if i == 0 {
			s = append(s, "")
		} else {
			s = append(s, itoa(i))
		}
	}
	s = append(s, "")
	for _, i := range m.poz {
		if i == 0 {
			s = append(s, "")
		} else {
			s = append(s, itoa(i))
		}
	}
	s = append(s, m.zp.String())
	if m.zh > 0 {
		s = append(s, itoa(m.krok))
	} else {
		s = append(s, "")
	}

	presun := make([]string, m.k.Len())
	for _, p := range m.zp.p {
		i := p[1] - 1
		presun[i] = itoa(p[0]) + "|" + itoa(p[1])
	}
	s = append(s, presun...)
	return s
}

type mapaZhoda struct {
	n           int
	zhodaCelkom []int
	krok        []int

	r         []*mapaZhodaRiadok
	citatel   *statistikaZhoda
	menovatel *statistikaZhoda
}

func newMapaZhoda(r []Riadok, n int) *mapaZhoda {
	m := &mapaZhoda{
		n:           n,
		zhodaCelkom: make([]int, n+1),
		krok:        make([]int, 1),
		r:           make([]*mapaZhodaRiadok, 0, len(r)),
		citatel:     newStatistikaZhoda(len(r), n),
		menovatel:   newStatistikaZhoda(len(r), n),
	}

	var pred Kombinacia
	for _, r := range r {
		m.add(r.Pc, pred, r.K)
		pred = r.K
	}

	return m
}

func (m *mapaZhoda) add(riadok int, pred, akt Kombinacia) {

	// vytvorim aktualny riadok
	r := newMapaZhodaRiadok(pred, akt)

	// z je zhoda
	z := r.zh

	// incrementujem zhodu 0..n
	m.zhodaCelkom[z]++

	// inkrementujem posledny krok
	m.krok[len(m.krok)-1]++

	// nastala zhoda?
	if z > 0 {
		// v riadku zapisem krok
		r.krok = m.krok[len(m.krok)-1]
		// pridam dalsi krok
		m.krok = append(m.krok, 0)
	}

	// mame uz viac ako jeden riadkov?
	if len(m.r) > 0 {
		r0 := m.r[len(m.r)-1]
		// ak nastala zhoda v riadku
		// nastavim pozicie v riadku-1
		r0.merge(r)
	}
	m.r = append(m.r, r)

	// statistika
	if z > 0 {
		var udalostC []int
		var udalostM []int
		for _, p := range r.zp.p {
			udalostC = append(udalostC, p[0])
			udalostM = append(udalostM, p[1])
		}
		m.citatel.add(riadok-1, udalostC)
		m.menovatel.add(riadok, udalostM)
	}
}

func (m *mapaZhoda) header() []string {
	h := []string{"Kombinacie", "Zh"}
	for i := 1; i <= m.n; i++ {
		h = append(h, itoa(i))
	}
	h = append(h, "")
	for i := 1; i <= m.n; i++ {
		h = append(h, itoa(i))
	}
	h = append(h, "Pozicia Zh")
	h = append(h, "Krok")
	for i := 1; i <= m.n; i++ {
		h = append(h, itoa(i)+" stl r-1/r")
	}
	return h
}

func (m *mapaZhoda) strings() [][]string {
	strings := make([][]string, len(m.r))
	for i, r := range m.r {
		strings[i] = r.strings()
	}
	cstrings := m.citatel.strings()
	mstrings := m.menovatel.strings()
	for i := range cstrings {
		s := []string{""}
		s = append(s, cstrings[i]...)
		// s = append(s, "")
		s = append(s, mstrings[i]...)
		strings = append(strings, s)
	}
	return strings
}
