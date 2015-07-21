package hrx

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/melias122/psl/num"
)

type HrxTab struct {
	n, m int
	r0   Presun
	r1   Presun
	Hrx  *H
	HHrx *H

	skupiny Skupiny

	skupinyN1 map[int][]*num.N
	skupinyN2 map[int][]*num.N

	header []string

	pc            uint
	riadokHrx     float64
	riadokHHrx    float64
	min, max      int
	pocetSucet    big.Int
	rod           float64
	r1Min, r1Max  float64
	hhrxR0, hrxR0 float64

	//sMinMax
	sMM   [][4]float64
	sNums num.Nums

	w *csv.Writer

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

	r0 := Hrx.Presun()
	return &HrxTab{
		n:          n,
		m:          m,
		r0:         r0,
		r1:         make(Presun, 0, len(r0)+1),
		Hrx:        Hrx,
		HHrx:       HHrx,
		riadokHrx:  Hrx.Value(),
		riadokHHrx: HHrx.Value(),
		pocetSucet: *big.NewInt(1),

		sMM:   make([][4]float64, n),
		sNums: make(num.Nums, 0, m),

		psCache:  make(map[Tab]*big.Int),
		mmCache:  make(map[Tab][2]int),
		rodCache: make(map[Tab]float64),
		rCache:   make(map[Tab][2]float64),

		header: header,
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// TODO: Najst lepsi sposob hladania min, max v STL
// pozn: spravit kontrolu, toto je asi nekorektne
func (h *HrxTab) sMinMax() (float64, float64, float64, float64) {
	var (
		s        = h.sMM
		nums     = h.sNums[:0]
		sums     [4]float64
		from, to int
	)
	for i := 0; i < h.n; i++ {
		s[i][0] = math.MaxFloat64
		s[i][1] = 0
		s[i][2] = math.MaxFloat64
		s[i][3] = 0
	}

	for _, t := range h.r1 {
		nums = append(nums, h.skupinyN1[t.Sk]...)
	}
	sort.Sort(nums)

	for i, n2 := range nums {
		if i < h.n-1 { // zaciatok pyramidy
			from = 0
			to = i + 1
		} else if i > len(nums)-h.n-1 { // vrch pyramidy
			from = h.n - (len(nums) - i)
			to = h.n
		} else { // stred
			from = 0
			to = h.n
		}
		n1 := h.HHrx.GetN(n2.Cislo())
		for j := from; j < to; j++ {
			//STL1-DO
			s[j][0] = min(n1.S(j+1), s[j][0])
			s[j][1] = max(n1.S(j+1), s[j][1])

			//STLOD-DO
			s[j][2] = min(n2.S(j+1), s[j][2])
			s[j][3] = max(n2.S(j+1), s[j][3])
		}
	}

	for i := 0; i < h.n; i++ {
		sums[0] += s[i][0]
		sums[1] += s[i][1]
		sums[2] += s[i][2]
		sums[3] += s[i][3]
	}
	return sums[0], sums[1], sums[2], sums[3]
}

func (h *HrxTab) hhrxMinMax() (float64, float64) {
	// presun v HHrx (HHrx min)
	for _, t := range h.r1 {
		c := h.skupinyN2[t.Sk]
		for i := 0; i < t.Max; i++ {
			h.HHrx.move(1, c[i].PocetR(), c[i].PocetR()+1)
		}
	}
	hhrxMin := h.HHrx.Value()
	for _, t := range h.r1 {
		c := h.skupinyN2[t.Sk]
		for i := 0; i < t.Max; i++ {
			h.HHrx.move(1, c[i].PocetR()+1, c[i].PocetR())
		}
	}

	// presun v HHrx (HHrx max)
	for _, t := range h.r1 {
		c := h.skupinyN2[t.Sk]
		lastIndex := len(c) - 1
		for i := 0; i < t.Max; i++ {
			h.HHrx.move(1, c[lastIndex-i].PocetR(), c[lastIndex-i].PocetR()+1)
		}
	}
	hhrxMax := h.HHrx.Value()
	for _, t := range h.r1 {
		c := h.skupinyN2[t.Sk]
		lastIndex := len(c) - 1
		for i := 0; i < t.Max; i++ {
			h.HHrx.move(1, c[lastIndex-i].PocetR()+1, c[lastIndex-i].PocetR())
		}
	}
	return hhrxMin, hhrxMax
}

func (h *HrxTab) record() []string {
	h.pc++
	hrx := h.Hrx.Value()

	s1min, s1max, s2min, s2max := h.sMinMax()

	// Presunut do append a delete
	// Treba mat 2 kopie HHrx
	// Jedna na min, druha na max
	hhrxMin, hhrxMax := h.hhrxMinMax()

	h.skupiny = append(h.skupiny, Skupina{
		Hrx:    hrx,
		HHrx:   [2]float64{hhrxMin, hhrxMax},
		Presun: h.r1.copyNonZero(),
	})

	r := make([]string, 0, len(h.header))
	r = append(r,
		itoa(int(h.pc)), // pc
		ftoa(hrx),
		ftoa(hrx-h.riadokHrx),
		h.r1.String(),
		ftoa(h.rod),
		ftoa(s2min),
		ftoa(s2max),
		itoa(h.min),
		itoa(h.max),
		h.pocetSucet.String(),
		ftoa(hhrxMin),
		ftoa(hhrxMax),
		ftoa(hhrxMin-h.riadokHHrx),
		ftoa(hhrxMax-h.riadokHHrx),
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
	h.Hrx.move(t.Max, t.Sk, t.Sk+1)

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
	h.Hrx.move(t.Max, t.Sk+1, t.Sk)

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

func (h *HrxTab) make(r0 Presun, n int) error {
	if len(r0) == 0 {
		return nil
	}
	max := r0[0].Max
	if max > n {
		max = n
	}
	for max > 0 {
		h.append(Tab{r0[0].Sk, max})
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

func (h *HrxTab) Make() (Skupiny, error) {

	// vytvorenie suboru
	f, err := os.Create(fmt.Sprintf("%d%d/HrxHHrx_%d%d.csv", h.n, h.m, h.n, h.m))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = ';'
	h.w = w

	// priradenie skutocnych cisel(num.N)
	// do skupin podla pocetnosti R2
	skupinyN1 := map[int][]*num.N{}
	skupinyN2 := map[int][]*num.N{}
	for i := 1; i <= h.m; i++ {
		N := h.Hrx.GetN(i)
		sk := N.PocetR()
		skupinyN1[sk] = append(skupinyN1[sk], N)

		N2 := h.HHrx.GetN(i)
		skupinyN2[sk] = append(skupinyN2[sk], N2)
	}
	h.skupinyN1 = skupinyN1

	for i := range skupinyN2 {
		sort.Sort(num.ByPocetR{num.Nums(skupinyN2[i])})
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
			r2 = append(r2, itoa(c.PocetR()))
			r3 = append(r3, itoa(h.HHrx.GetN(c.Cislo()).PocetR()))
		}
		PreHeader = append(PreHeader, r1, r2, r3, []string{""})
	}
	PreHeader = append(PreHeader, h.header)
	if err = w.WriteAll(PreHeader); err != nil {
		return nil, err
	}

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
			h.psCache[Tab{t.Sk, max}] = &b

			// max, min sucet v skupine
			smin += skN[i].Cislo()
			smax += skN[len(skN)-1-i].Cislo()
			h.mmCache[Tab{t.Sk, i + 1}] = [2]int{smin, smax}

			// max ROD-DO hodnota v skupine
			h.rodCache[Tab{t.Sk, max}] = skupinyN1[t.Sk][0].R() * float64(max)

			// min,max R1-DO hodnota v skupine
			rmin += h.HHrx.GetN(skN2[i].Cislo()).R()
			rmax += h.HHrx.GetN(skN2[len(skN2)-1-i].Cislo()).R()
			h.rCache[Tab{t.Sk, i + 1}] = [2]float64{rmin, rmax}

			i++
		}
	}

	if err := h.make(h.r0, h.n); err != nil {
		return nil, err
	}

	w.Flush()
	return h.skupiny, w.Error()
}

type Tab struct {
	Sk  int
	Max int
}

type Presun []Tab

func (p Presun) copy() Presun {
	presun := make(Presun, len(p))
	for i := range p {
		presun[i] = p[i]
	}
	return presun
}

func (p Presun) copyNonZero() Presun {
	var presun Presun
	for i := range p {
		if p[i].Max > 0 {
			presun = append(presun, p[i])
		}
	}
	return presun
}

func (p Presun) Len() int           { return len(p) }
func (p Presun) Less(i, j int) bool { return p[i].Sk < p[j].Sk }
func (p Presun) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p Presun) String() string {
	if len(p) > 0 {
		s := make([]int, p[len(p)-1].Sk+1)
		for _, v := range p {
			s[v.Sk] = v.Max
		}
		s2 := make([]string, len(s))
		for i := range s {
			s2[i] = strconv.Itoa(s[i])
		}
		return strings.Join(s2, " ")
	} else {
		return ""
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func ftoa(f float64) string {
	s := strconv.FormatFloat(f, 'g', -1, 64)
	return strings.Replace(s, ".", ",", 1)
}

type Skupiny []Skupina

type Skupina struct {
	Hrx    float64
	HHrx   [2]float64
	Presun Presun
}
