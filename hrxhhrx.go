package psl

import (
	"math/big"
	"sort"
)

type hrxHHrxTab struct {
	n, m        int
	HrxXcisla   Xcisla
	Xcisla      Xcisla
	Hrx         *H
	HHrx        *H
	KombinaciaR Kombinacia

	hhrxMin, hhrxMax Xcisla

	skupiny Skupiny

	HrxNums  map[int]Nums
	HHrxNums map[int]Nums

	riadokHrx     float64
	riadokHHrx    float64
	min, max      int
	pocetSucet    big.Int
	rod           float64
	r1Min, r1Max  float64
	hhrxR0, hrxR0 float64
	cislovackyMin Cislovacky
	cislovackyMax Cislovacky
	zhMin, zhMax  int

	w *CsvMaxWriter

	psCache         map[Tab]*big.Int
	mmCache         map[Tab][2]int
	rodCache        map[Tab]float64
	rCache          map[Tab][2]float64
	cislovackyCache map[Tab][2]Cislovacky
	zhodaCache      map[Tab][2]int

	headerLen int
}

func makeSkupiny(archiv *Archiv) (Skupiny, error) {

	h := hrxHHrxTab{
		n:          archiv.n,
		m:          archiv.m,
		HrxXcisla:  archiv.Hrx.Xcisla(),
		Xcisla:     make(Xcisla, 0, n),
		Hrx:        archiv.Hrx,
		HHrx:       archiv.HHrx,
		riadokHrx:  archiv.Hrx.Value(),
		riadokHHrx: archiv.HHrx.Value(),
		pocetSucet: *big.NewInt(1),

		KombinaciaR: archiv.K,

		hhrxMin: archiv.HHrx.Xcisla(),
		hhrxMax: archiv.HHrx.Xcisla(),

		HrxNums:  make(map[int]Nums),
		HHrxNums: make(map[int]Nums),

		psCache:         make(map[Tab]*big.Int),
		mmCache:         make(map[Tab][2]int),
		rodCache:        make(map[Tab]float64),
		rCache:          make(map[Tab][2]float64),
		cislovackyCache: make(map[Tab][2]Cislovacky),
		zhodaCache:      make(map[Tab][2]int),
	}

	// priradenie skutocnych cisel (Num)
	// do skupin podla pocetnosti R2
	for i := 1; i <= h.m; i++ {
		hrxNum := h.Hrx.GetNum(i)
		sk := hrxNum.PocetR()
		h.HrxNums[sk] = append(h.HrxNums[sk], hrxNum)

		hhrxNum := h.HHrx.GetNum(i)
		h.HHrxNums[sk] = append(h.HHrxNums[sk], hhrxNum)
	}

	for _, hhrxNums := range h.HHrxNums {
		sort.Sort(ByPocetR{hhrxNums})
	}

	h.precompute()

	h.w = NewCsvMaxWriter("HrxHHrx", archiv.WorkingDir,
		setHeaders(h.header()),
		setSuffixFunc(IntSuffix()),
	)
	defer h.w.Close()

	if err := h.make(h.HrxXcisla.copy(), h.n); err != nil {
		return nil, err
	}
	return h.skupiny, nil
}

// TODO: Najst lepsi sposob hladania min, max v STL
// pozn: spravit kontrolu, toto je asi nekorektne
func (h *hrxHHrxTab) sMinMax() (float64, float64, float64, float64) {
	var (
		sums [4]float64

		nums          = make(Nums, 0, h.m)
		kombinaciaMin = make([]int, h.n)
		kombinaciaMax = make([]int, h.n)

		zS2min = make(map[int]bool, h.m)
		zS2max = make(map[int]bool, h.m)
		zS1min = make(map[int]bool, h.m)
		zS1max = make(map[int]bool, h.m)
	)

	for _, t := range h.Xcisla {
		nums = append(nums, h.HrxNums[t.Sk]...)
	}
	nums.Sort()
	for i := 0; i < h.n; i++ {
		kombinaciaMin[i] = nums[i].Cislo()
		kombinaciaMax[h.n-1-i] = nums[len(nums)-1-i].Cislo()
	}

	for i := 0; i < h.n; i++ {
		var (
			z     [4]int
			s1Min float64 = 1
			s1Max float64
			s2Min float64 = 1
			s2Max float64
		)
		for _, n2 := range nums {
			if !n2.HasSTL(i) {
				continue
			}
			if n2.Cislo() >= kombinaciaMin[i] && n2.Cislo() <= kombinaciaMax[i] {
				// STL OD-DO
				if n2.SNext(i+1) < s2Min && !zS2min[n2.Cislo()] {
					s2Min = n2.SNext(i + 1)
					z[2] = n2.Cislo()
				}
				if n2.SNext(i+1) > s2Max && !zS2max[n2.Cislo()] {
					s2Max = n2.SNext(i + 1)
					z[3] = n2.Cislo()
				}

				n1 := h.HHrx.GetNum(n2.Cislo())
				if n1.SNext(i+1) < s1Min && !zS1min[n2.Cislo()] {
					s1Min = n1.SNext(i + 1)
					z[0] = n2.Cislo()
				}
				if n1.SNext(i+1) > s1Max && !zS1max[n2.Cislo()] {
					s1Max = n1.SNext(i + 1)
					z[1] = n2.Cislo()
				}
			}
		}
		zS1min[z[0]] = true
		zS1max[z[1]] = true
		zS2min[z[2]] = true
		zS2max[z[3]] = true

		sums[0] += s1Min
		sums[1] += s1Max
		sums[2] += s2Min
		sums[3] += s2Max
	}
	return sums[0], sums[1], sums[2], sums[3]
}

func (h *hrxHHrxTab) formatCislovackyMinMax() []string {
	s := make([]string, 20)
	for i := 0; i < 10; i++ {
		j := i * 2
		s[j] = itoa(int(h.cislovackyMin[i]))
		s[j+1] = itoa(int(h.cislovackyMax[i]))
	}
	return s
}

func (h *hrxHHrxTab) record() []string {
	hrx := h.Hrx.valuePresun(h.HrxXcisla)

	// var s1min, s1max, s2min, s2max float64
	s1min, s1max, s2min, s2max := h.sMinMax()

	hhrxMin := h.HHrx.valuePresun(h.hhrxMin)
	hhrxMax := h.HHrx.valuePresun(h.hhrxMax)

	h.skupiny = append(h.skupiny, Skupina{
		R1:   [2]float64{h.r1Min, h.r1Max},
		S1:   [2]float64{s1min, s1max},
		HHrx: [2]float64{hhrxMin, hhrxMax},

		R2:  h.rod,
		S2:  [2]float64{s2min, s2max},
		Hrx: hrx,

		Sucet: [2]uint16{uint16(h.min), uint16(h.max)},

		Cislovacky: [2]Cislovacky{h.cislovackyMin, h.cislovackyMax},
		Zh:         [2]byte{byte(h.zhMin), byte(h.zhMax)},

		// Cifrovacky:,

		Xcisla: h.Xcisla.copy(),
	})

	r := make([]string, 0, h.headerLen)
	r = append(r,
		ftoa(hrx),
		ftoa(hrx-h.riadokHrx),
		h.Xcisla.String(),
		ftoa(h.rod),
		// ntice
		// xtice
		ftoa(s2min),
		ftoa(s2max),
		// pocet s2
		itoa(h.min),
		itoa(h.max),
		h.pocetSucet.String(),
		// "#Kombinacie"
		ftoa(hhrxMin),
		ftoa(hhrxMax),
		// "HHrx (počet)"
		ftoa(h.riadokHHrx-hhrxMin),
		ftoa(h.riadokHHrx-hhrxMax),
		ftoa(h.r1Min),
		ftoa(h.r1Max),
		// "ƩR 1-DO (počet)"
		// "Teor. max. pocet"
		ftoa(s1min),
		ftoa(s1max),
		// "ƩSTL 1-DO (počet)"
	)
	r = append(r, h.formatCislovackyMinMax()...)
	r = append(r, itoa(h.zhMin), itoa(h.zhMax))

	return r
}

func (h *hrxHHrxTab) append(t Tab) {
	// presun retazec
	h.Xcisla = append(h.Xcisla, t)

	// presun v Hrx
	h.HrxXcisla.move(t.Max, t.Sk, t.Sk+1)

	// presun v HHrx min, max
	N2 := h.HHrxNums[t.Sk]
	N2LastIndex := len(N2) - 1
	for i := 0; i < t.Max; i++ {
		h.hhrxMin.move(1, N2[i].PocetR(), N2[i].PocetR()+1)
		h.hhrxMax.move(1, N2[N2LastIndex-i].PocetR(), N2[N2LastIndex-i].PocetR()+1)
	}

	// pocet suctov
	h.pocetSucet.Mul(&h.pocetSucet, h.psCache[t])

	// min, max sucet
	mm := h.mmCache[t]
	h.min += mm[0]
	h.max += mm[1]

	// max ROD-DO
	h.rod += h.rodCache[t]

	// min, max R1-DO
	r1mm := h.rCache[t]
	h.r1Min += r1mm[0]
	h.r1Max += r1mm[1]

	// Cislovacky min, max
	cMinMax := h.cislovackyCache[t]
	h.cislovackyMin.Plus(cMinMax[0])
	h.cislovackyMax.Plus(cMinMax[1])

	// Zhoda min, max
	zhMinMax := h.zhodaCache[t]
	h.zhMin += zhMinMax[0]
	h.zhMax += zhMinMax[1]

}

func (h *hrxHHrxTab) delete() {
	// presun retazec
	t := h.Xcisla[len(h.Xcisla)-1]
	h.Xcisla = h.Xcisla[:len(h.Xcisla)-1]

	// presun v Hrx
	h.HrxXcisla.move(t.Max, t.Sk+1, t.Sk)

	// presun v HHrx min, max
	N2 := h.HHrxNums[t.Sk]
	N2LastIndex := len(N2) - 1
	for i := 0; i < t.Max; i++ {
		h.hhrxMin.move(1, N2[i].PocetR()+1, N2[i].PocetR())
		h.hhrxMax.move(1, N2[N2LastIndex-i].PocetR()+1, N2[N2LastIndex-i].PocetR())
	}

	// pocet suctov
	h.pocetSucet.Div(&h.pocetSucet, h.psCache[t])

	// min, max sucet
	mm := h.mmCache[t]
	h.min -= mm[0]
	h.max -= mm[1]

	// max ROD-DO
	h.rod -= h.rodCache[t]

	// min, max R1-DO
	r1mm := h.rCache[t]
	h.r1Min -= r1mm[0]
	h.r1Max -= r1mm[1]

	// Cislovacky min, max
	cMinMax := h.cislovackyCache[t]
	h.cislovackyMin.Minus(cMinMax[0])
	h.cislovackyMax.Minus(cMinMax[1])

	// Zhoda min, max
	zhMinMax := h.zhodaCache[t]
	h.zhMin -= zhMinMax[0]
	h.zhMax -= zhMinMax[1]
}

func (h *hrxHHrxTab) make(HrxXcisla Xcisla, n int) error {
	if len(HrxXcisla) == 0 {
		return nil
	}
	max := HrxXcisla[0].Max
	if max > n {
		max = n
	}
	for max > 0 {
		h.append(newTab(HrxXcisla[0].Sk, max))
		if n-max > 0 {
			if err := h.make(HrxXcisla[1:], n-max); err != nil {
				return err
			}
		} else {

			if err := h.w.Write(h.record()); err != nil {
				return err
			}
		}
		h.delete()
		max--
	}
	return h.make(HrxXcisla[1:], n)
}

// maximum vrati maximalny mozny pocet typu cisla
// v skupine. napriklad max pocet: P, N,...
// n je celkovy pocet danych cisiel v skupine
// smax je maximalny pocet cisiel
// ktore mozu by vybrate zo skupiny
func (h *hrxHHrxTab) maximum(n, smax int) int {
	return min(n, smax)
}

// minimum vrati minimalny mozny pocet typu cisla
// v skupine. napriklad min pocet: P, N,...
// n je celkovy pocet danych cisiel v skupine
// smax je maximalny pocet cisiel ktore mozu by vybrate zo skupiny
// ssize je velkost skupiny
func (h *hrxHHrxTab) minimum(n, smax, ssize int) int {
	min := ssize - n - smax
	if min >= 0 {
		min = 0
	} else {
		min = -min
	}
	return min
}

func (h *hrxHHrxTab) maxCislovacky(smax int, cislovacky Cislovacky) Cislovacky {
	var c Cislovacky
	for i, j := range cislovacky {
		c[i] = byte(h.maximum(int(j), smax))
	}
	return c
}

func (h *hrxHHrxTab) minCislovacky(smax int, ssize int, cislovacky Cislovacky) Cislovacky {
	var c Cislovacky
	for i, j := range cislovacky {
		c[i] = byte(h.minimum(int(j), smax, ssize))
	}
	return c
}

// Prepocitanie znamych hodnot
func (h *hrxHHrxTab) precompute() {
	for _, t := range h.HrxXcisla {
		max := t.Max
		if max > h.n {
			max = h.n
		}
		var (
			i, smin, smax int
			rmin, rmax    float64
			cislovacky    Cislovacky
		)
		skN := h.HrxNums[t.Sk]
		skN2 := h.HHrxNums[t.Sk]

		// navyssi mozny pocet danej cislovacky v skupine
		for _, num := range skN {
			c2 := NewCislovacky(num.Cislo())
			cislovacky.Plus(c2)
		}

		// navyssi mozny pocet zhod z poslednym riadkom v skupine
		var zhodaK Kombinacia
		for _, num := range skN {
			zhodaK.Append(byte(num.Cislo()))
		}
		zhodaMax := Zhoda(h.KombinaciaR, zhodaK)

		for ; max > 0; max-- {

			// min, max zhoda v skupine
			h.zhodaCache[newTab(t.Sk, max)] = [2]int{
				h.minimum(zhodaMax, max, len(skN)),
				h.maximum(zhodaMax, max),
			}

			// min, max cislovacky v skupine
			h.cislovackyCache[newTab(t.Sk, max)] = [2]Cislovacky{
				h.minCislovacky(max, len(skN), cislovacky),
				h.maxCislovacky(max, cislovacky),
			}

			// pocet suctov v skupine
			b := new(big.Int)
			b.Binomial(int64(t.Max), int64(max))
			h.psCache[newTab(t.Sk, max)] = b

			// max, min sucet v skupine
			smin += skN[i].Cislo()
			smax += skN[len(skN)-1-i].Cislo()
			h.mmCache[newTab(t.Sk, i+1)] = [2]int{smin, smax}

			// max ROD-DO hodnota v skupine
			h.rodCache[newTab(t.Sk, max)] = h.HrxNums[t.Sk][0].RNext() * float64(max)

			// min,max R1-DO hodnota v skupine
			rmin += h.HHrx.GetNum(skN2[i].Cislo()).RNext()
			rmax += h.HHrx.GetNum(skN2[len(skN2)-1-i].Cislo()).RNext()
			h.rCache[newTab(t.Sk, i+1)] = [2]float64{rmin, rmax}

			i++
		}
	}
}

func (h *hrxHHrxTab) header() [][]string {

	var header [][]string

	// zoradenie skupin podla pocetnosti
	// pridaniu do suboru pred hlavicku
	pocetnostiHrx := make([]int, 0, len(h.HrxNums))
	for sk := range h.HrxNums {
		pocetnostiHrx = append(pocetnostiHrx, sk)
	}
	sort.Ints(pocetnostiHrx)

	for _, sk := range pocetnostiHrx {
		var r1, r2, r3 []string
		r1 = append(r1, "Cislo")
		r2 = append(r2, "Pocet R 1-DO")
		r3 = append(r3, "Pocet R OD-DO")
		for _, num := range h.HrxNums[sk] {
			r1 = append(r1, num.String())
			r2 = append(r2, itoa(h.HHrx.GetNum(num.Cislo()).PocetR()))
			r3 = append(r3, itoa(num.PocetR()))
		}
		header = append(header, r1, r2, r3, []string{""})
	}

	// hlavicka suboru HrxHHrx
	realHeader := []string{
		"Hrx",
		"ΔHrx",
		"Xcisla",
		"ƩR OD-DO",
		// "Ntice",
		// "Xtice",
		"ƩSTL OD-DO (min)",
		"ƩSTL OD-DO (max)",
		// "ƩSTL OD-DO (počet)",
		"ƩKombinacie (min)",
		"ƩKombinacie (max)",
		"Kombinacie (počet)",
		"HHrx (min)",
		"HHrx (max)",
		// "HHrx (počet)",
		"ΔHHrx (min)",
		"ΔHHrx (max)",
		"ƩR 1-DO (min)",
		"ƩR 1-DO (max)",
		// "ƩR 1-DO (počet)",
		// "Teor. max. pocet",
		"ƩSTL 1-DO (min)",
		"ƩSTL 1-DO (max)",
		// "ƩSTL 1-DO (počet)",
		"P (min)", "P (max)",
		"N (min)", "N (max)",
		"Pr (min)", "Pr (max)",
		"Mc (min)", "Mc (max)",
		"Vc (min)", "Vc (max)",
		"C19 (min)", "C19 (max)",
		"C0 (min)", "C0 (max)",
		"cC (min)", "cC (max)",
		"Cc (min)", "Cc (max)",
		"CC (min)", "CC (max)",
		"Zh (min)", "Zh (max)",
	}
	header = append(header, realHeader)

	h.headerLen = len(realHeader)

	return header
}
