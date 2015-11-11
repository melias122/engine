package psl

import (
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type HrxTab struct {
	n, m int
	r0   Xcisla
	r1   Xcisla
	Hrx  *H
	HHrx *H

	hhrxMin, hhrxMax Xcisla

	skupiny Skupiny

	skupinyN1 map[int]Nums
	skupinyN2 map[int]Nums

	header []string

	pc            uint
	riadokHrx     float64
	riadokHHrx    float64
	min, max      int
	pocetSucet    big.Int
	rod           float64
	r1Min, r1Max  float64
	hhrxR0, hrxR0 float64

	w *CsvMaxWriter

	psCache  map[Tab]*big.Int
	mmCache  map[Tab][2]int
	rodCache map[Tab]float64
	rCache   map[Tab][2]float64
}

func NewHrxTab(Hrx, HHrx *H, n, m int) *HrxTab {
	// hlavicka suboru HrxHHrx
	header := []string{
		"p.c.", "HRX pre r+1", "dHRX diferencia s \"r\"", "presun z r do (r+1)cisla",
		"∑%ROD-DO", "∑%STLOD-DO (min)", "∑%STLOD-DO (max)", "∑ kombi (min)", "∑ kombi (max)",
		"Pocet ∑ kombi", "HHRX pre r+1 (min)", "HHRX pre r+1 (max)", "dHHRX s \"r\" (min)",
		"dHHRX s \"r\" (max)", "∑%R1-DO (min)", "∑%R1-DO (max)", "Teor. max. pocet",
		"∑%STL1-DO (min)", "∑%STL1-DO (max)",
	}

	return &HrxTab{
		n:          n,
		m:          m,
		r0:         Hrx.Xcisla(),
		r1:         make(Xcisla, 0, n),
		Hrx:        Hrx,
		HHrx:       HHrx,
		riadokHrx:  Hrx.Value(),
		riadokHHrx: HHrx.Value(),
		pocetSucet: *big.NewInt(1),

		hhrxMin: HHrx.Xcisla(),
		hhrxMax: HHrx.Xcisla(),

		// sMM:   make([][4]float64, n),
		// sNums: make(num.Nums, 0, m),

		psCache:  make(map[Tab]*big.Int),
		mmCache:  make(map[Tab][2]int),
		rodCache: make(map[Tab]float64),
		rCache:   make(map[Tab][2]float64),

		header: header,
	}
}

// TODO: Najst lepsi sposob hladania min, max v STL
// pozn: spravit kontrolu, toto je asi nekorektne
func (h *HrxTab) sMinMax() (float64, float64, float64, float64) {
	var (
		sums [4]float64

		nums          = make(Nums, 0, h.m)
		kombinaciaMin = make([]int, h.n)
		kombinaciaMax = make([]int, h.n)

		zS2min = make(map[int]bool)
		zS2max = make(map[int]bool)
		zS1min = make(map[int]bool)
		zS1max = make(map[int]bool)
	)

	for _, t := range h.r1 {
		nums = append(nums, h.skupinyN1[t.Sk]...)
	}
	sort.Sort(nums)
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

func (h *HrxTab) record() []string {
	h.pc++
	hrx := h.Hrx.valuePresun(h.r0)
	s1min, s1max, s2min, s2max := h.sMinMax()

	hhrxMin := h.HHrx.valuePresun(h.hhrxMin)
	hhrxMax := h.HHrx.valuePresun(h.hhrxMax)

	h.skupiny = append(h.skupiny, Skupina{
		Hrx:    hrx,
		HHrx:   [2]float64{hhrxMin, hhrxMax},
		R1:     [2]float64{h.r1Min, h.r1Max},
		R2:     h.rod,
		Sucet:  [2]uint16{uint16(h.min), uint16(h.max)},
		Xcisla: h.r1.copy(),
	})

	r := make([]string, 0, len(h.header))
	r = append(r,
		itoa(int(h.pc)), // pc
		ftoa(hrx),
		ftoa(h.riadokHrx-hrx),
		h.r1.String(),
		ftoa(h.rod),
		ftoa(s2min),
		ftoa(s2max),
		itoa(h.min),
		itoa(h.max),
		h.pocetSucet.String(),
		ftoa(hhrxMin),
		ftoa(hhrxMax),
		ftoa(h.riadokHHrx-hhrxMin),
		ftoa(h.riadokHHrx-hhrxMax),
		ftoa(h.r1Min),
		ftoa(h.r1Max),
		"",
		ftoa(s1min),
		ftoa(s1max),
	)
	return r
}

func (h *HrxTab) append(t Tab) {
	// presun retazec
	h.r1 = append(h.r1, t)

	// presun v Hrx
	h.r0.move(t.Max, t.Sk, t.Sk+1)

	// presun v HHrx min, max
	N2 := h.skupinyN2[t.Sk]
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
}

func (h *HrxTab) delete() {
	// presun retazec
	t := h.r1[len(h.r1)-1]
	h.r1 = h.r1[:len(h.r1)-1]

	// presun v Hrx
	h.r0.move(t.Max, t.Sk+1, t.Sk)

	// presun v HHrx min, max
	N2 := h.skupinyN2[t.Sk]
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
}

func (h *HrxTab) make(r0 Xcisla, n int) error {
	if len(r0) == 0 {
		return nil
	}
	max := r0[0].Max
	if max > n {
		max = n
	}
	for max > 0 {
		h.append(newTab(r0[0].Sk, max))
		if n-max > 0 {
			if err := h.make(r0[1:], n-max); err != nil {
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
	return h.make(r0[1:], n)
}

func (h *HrxTab) Make(workingDir string) (Skupiny, error) {

	// priradenie skutocnych cisel(num.N)
	// do skupin podla pocetnosti R2
	skupinyN1 := map[int]Nums{}
	skupinyN2 := map[int]Nums{}
	for i := 1; i <= h.m; i++ {
		N := h.Hrx.GetNum(i)
		sk := N.PocetR()
		skupinyN1[sk] = append(skupinyN1[sk], N)

		N2 := h.HHrx.GetNum(i)
		skupinyN2[sk] = append(skupinyN2[sk], N2)
	}
	h.skupinyN1 = skupinyN1

	for i := range skupinyN2 {
		sort.Sort(ByPocetR{Nums(skupinyN2[i])})
	}
	h.skupinyN2 = skupinyN2
	// skupiny cisel a ich zoradenie kvoli
	// pridaniu do suboru pred hlavicku
	// Pozn.: zbytocna blbost, pocetnosti,
	// resp. skupiny je uz vidiet PocetnostR subore
	mKeys := make([]int, 0, len(skupinyN1))
	for k := range skupinyN1 {
		mKeys = append(mKeys, k)
	}
	sort.Ints(mKeys)
	var PreHeader [][]string
	for _, k := range mKeys {
		var r1, r2, r3 []string
		r1 = append(r1, "Cislo")
		r2 = append(r2, "Pocet R1-DO")
		r3 = append(r3, "Pocet ROD-DO")
		for _, c := range skupinyN1[k] {
			r1 = append(r1, c.String())
			r2 = append(r2, itoa(h.HHrx.GetNum(c.Cislo()).PocetR()))
			r3 = append(r3, itoa(c.PocetR()))
		}
		PreHeader = append(PreHeader, r1, r2, r3, []string{""})
	}
	PreHeader = append(PreHeader, h.header)

	// Prepocitanie znamych hodnot
	for _, t := range h.r0 {
		max := t.Max
		if max > h.n {
			max = h.n
		}

		var (
			i, smin, smax int
			rmin, rmax    float64
		)
		skN := skupinyN1[t.Sk]
		skN2 := skupinyN2[t.Sk]
		for ; max > 0; max-- {

			// pocet suctov v skupine
			var b big.Int
			b.Binomial(int64(t.Max), int64(max))
			h.psCache[newTab(t.Sk, max)] = &b

			// max, min sucet v skupine
			smin += skN[i].Cislo()
			smax += skN[len(skN)-1-i].Cislo()
			h.mmCache[newTab(t.Sk, i+1)] = [2]int{smin, smax}

			// max ROD-DO hodnota v skupine
			h.rodCache[newTab(t.Sk, max)] = skupinyN1[t.Sk][0].RNext() * float64(max)

			// min,max R1-DO hodnota v skupine
			rmin += h.HHrx.GetNum(skN2[i].Cislo()).RNext()
			rmax += h.HHrx.GetNum(skN2[len(skN2)-1-i].Cislo()).RNext()
			h.rCache[newTab(t.Sk, i+1)] = [2]float64{rmin, rmax}

			i++
		}
	}
	h.w = NewCsvMaxWriter(workingDir, "HrxHHrx", PreHeader)
	h.w.Suffix = IntSuffix()
	defer h.w.Close()
	if err := h.make(h.r0.copy(), h.n); err != nil {
		return nil, err
	}
	return h.skupiny, nil
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func ftoa(f float64) string {
	s := strconv.FormatFloat(f, 'g', -1, 64)
	return strings.Replace(s, ".", ",", 1)
}
