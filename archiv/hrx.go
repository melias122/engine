package archiv

// func sucetMinMax(m map[int][]*num.N, p hrx.Presun) string {
// 	min, max := 0, 0
// 	for _, t := range p {
// 		arr := m[t.Sk]
// 		for i := 0; i < t.Max; i++ {
// 			min += arr[i].Cislo()
// 			max += arr[len(arr)-1-i].Cislo()
// 		}
// 	}
// 	return strings.Join([]string{itoa(min), itoa(max)}, "-")
// }

// func pocetSucet(m map[int][]*num.N, p hrx.Presun) string {
// 	pocet := big.NewInt(1)
// 	for _, t := range p {
// 		arr := m[t.Sk]
// 		pocet.Mul(pocet, big.NewInt(0).Binomial(int64(len(arr)), int64(t.Max)))
// 	}
// 	return pocet.String()
// }

// func (a *Archiv) HrxHHrx() error {
// 	// hlavicka suboru HrxHHrx
// 	header := []string{
// 		"p.c.", "HRX pre r+1", "dHRX diferencia s \"r\"", "presun z r do (r+1)cisla",
// 		"∑%ROD-DO", "∑%STLOD-DO od do", "∑ kombi od do", "Pocet ∑ kombi", "HHRX pre r+1",
// 		"dHHRX diferencia s \"r\"", "∑%R1-DO od do", "Teor. max. pocet", "∑%R1-DO", "∑%STL1-DO od do",
// 	}
// 	// vytvorenie suboru
// 	f, err := os.Create(fmt.Sprintf("%d%d/HrxHHrx_%d%d.csv", a.n, a.m, a.n, a.m))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	w := csv.NewWriter(f)
// 	w.Comma = ';'

// 	// priradenie skutocnych cisel(num.N)
// 	// do skupin podla pocetnosti R2
// 	m := map[int][]*num.N{}
// 	for i := 1; i <= a.m; i++ {
// 		c := a.Cisla[i-1]
// 		m[c.PocR2()] = append(m[c.PocR2()], c)
// 	}
// 	// skupiny cisel a ich zoradenie kvoli
// 	// pridaniu do suboru pred hlavicku
// 	// Pozn.: zbytocna blbost, pocetnosti,
// 	// resp. skupiny je uz vidiet PocetnostR subore
// 	mKeys := make([]int, 0, len(m))
// 	for k := range m {
// 		mKeys = append(mKeys, k)
// 	}
// 	sort.Ints(mKeys)
// 	var PreHeader [][]string
// 	for _, k := range mKeys {
// 		var r1, r2, r3 []string
// 		r1 = append(r1, "Cislo")
// 		r2 = append(r2, "Pocet R1-DO")
// 		r3 = append(r3, "Pocet ROD-DO")
// 		for _, c := range m[k] {
// 			r1 = append(r1, c.String())
// 			r2 = append(r2, itoa(c.PocR1()))
// 			r3 = append(r3, itoa(c.PocR2()))
// 		}
// 		PreHeader = append(PreHeader, r1, r2, r3, []string{""})
// 	}
// 	PreHeader = append(PreHeader, header)
// 	if err = w.WriteAll(PreHeader); err != nil {
// 		return err
// 	}

// 	ch := hrx.GenerujPresun(a.Hrx.Presun(), a.n)

// 	i := 1
// 	for p := range ch {
// 		hrx := a.Hrx.Simul(p)

// 		r := make([]string, 0, len(header))
// 		r = append(r,
// 			itoa(i),                // pc
// 			ftoa(hrx),              // hrx
// 			ftoa(hrx-a.Riadok.Hrx), // dif hrx s r
// 			p.String(),             // presun
// 			ftoa(0),                // r2
// 			ftoa(0),                // stl2
// 			sucetMinMax(m, p),      // sucet
// 			pocetSucet(m, p),       // pocet suctov
// 		)
// 		if err := w.Write(r); err != nil {
// 			return err
// 		}
// 		i++
// 	}

// 	w.Flush()
// 	return w.Error()
// }
