package engine

import (
	"bytes"
	"math"
	"sort"
	"strconv"
)

var HeaderV2 = []string{
	"HRX (r+1)",
	"ΔHRX",
	"Xcisla",
	"ƩR OD-DO",
	"Ntice",
	"Xtice",
	"ƩSTL OD-DO (min)",
	"ƩSTL OD-DO (max)",
	"ƩSTL OD-DO (počet)",
	"ƩKombinacie (min)",
	"ƩKombinacie (max)",
	"Kombinacie (počet)",
	"HHrx (min)",
	"HHrx (max)",
	"HHrx (počet)",
	// "ΔHHrx (min)",
	// "ΔHHrx (max)",
	"ƩR 1-DO (min)",
	"ƩR 1-DO (max)",
	"ƩR 1-DO (počet)",
	"ƩSTL 1-DO (min)",
	"ƩSTL 1-DO (max)",
	"ƩSTL 1-DO (počet)",
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

// vystup filter
type V2 struct {
	n, m      int
	Hrx, HHrx *H
	// r         Riadok
	p Xcisla

	hrx              float64
	sucet            [2]int
	hhrx, r1, s1, s2 map[float64]int
	nKombi           int
	ntice, xtice     map[string]int
	cislovacky       []Cislovacky
	zhoda            [2]int

	K0   Kombinacia
	HRX0 float64
	R2   float64
}

func NewV2(a *Archiv, sk Skupina) V2 {
	return V2{
		n:    a.n,
		m:    a.m,
		Hrx:  a.Hrx,
		HHrx: a.HHrx,
		// r:    a.Riadok,
		p: sk.Xcisla,

		hrx:   sk.Hrx,
		sucet: [2]int{math.MaxInt32, 0},
		hhrx:  make(map[float64]int),
		r1:    make(map[float64]int),
		s1:    make(map[float64]int),
		s2:    make(map[float64]int),
		ntice: make(map[string]int),
		xtice: make(map[string]int),
		zhoda: [2]int{math.MaxInt32, 0},

		K0:   a.Riadok.K,
		HRX0: a.Riadok.Hrx,
	}
}

func (v *V2) Add(k Kombinacia) {
	v.nKombi++

	// zhoda min, max
	zhoda := Zhoda(v.K0, k)
	v.zhoda[0] = min(v.zhoda[0], zhoda)
	v.zhoda[1] = max(v.zhoda[1], zhoda)

	// sucet min, max
	sucet := k.Sucet()
	v.sucet[0] = min(v.sucet[0], sucet)
	v.sucet[1] = max(v.sucet[1], sucet)

	// STL2 min, max, pocet
	R2, S2 := k.SucetRSNext(v.Hrx.Cisla)
	if _, ok := v.s2[S2]; !ok {
		v.s2[S2] = 1
	}
	v.R2 = R2

	// R1, STL1 min, max, pocet
	R1, S1 := k.SucetRSNext(v.HHrx.Cisla)
	if _, ok := v.r1[R1]; !ok {
		v.r1[R1] = 1
	}
	if _, ok := v.s1[S1]; !ok {
		v.s1[S1] = 1
	}

	// HHrx min, max, pocet
	hhrx := v.HHrx.ValueKombinacia(k)
	if _, ok := v.hhrx[hhrx]; !ok {
		v.hhrx[hhrx] = 1
	}

	//  pocet Ntic
	v.ntice[Ntica(k).String()]++

	// pocet Xtic
	v.xtice[Xtica(v.m, k).String()]++

	// cislovacky
	v.cislovacky = append(v.cislovacky, k.Cislovacky())
}

func (v V2) Riadok() []string {
	r := make([]string, 0, len(HeaderV2))
	r = append(r, ftoa(v.hrx))
	r = append(r, ftoa(v.hrx-v.HRX0))
	r = append(r, v.p.String())
	r = append(r, ftoa(v.R2))
	r = append(r, v.formatTica(v.ntice))
	r = append(r, v.formatTica(v.xtice))
	r = append(r, v.formatFloatMap(v.s2)...)
	r = append(r, itoa(v.sucet[0]), itoa(v.sucet[1]))
	r = append(r, itoa(v.nKombi))
	r = append(r, v.formatFloatMap(v.hhrx)...)
	r = append(r, v.formatFloatMap(v.r1)...)
	r = append(r, v.formatFloatMap(v.s1)...)
	r = append(r, v.formatCislovacky()...)
	r = append(r, itoa(v.zhoda[0]), itoa(v.zhoda[1]))
	return r
}

func (v *V2) formatCislovacky() []string {
	var cmin, cmax Cislovacky
	for i := range cmax {
		cmin[i] = byte(99)
	}
	for _, c := range v.cislovacky {
		for i := range c {
			cmin[i] = byte(min(int(c[i]), int(cmin[i])))
			cmax[i] = byte(max(int(c[i]), int(cmax[i])))
		}
	}
	s := make([]string, 20)
	for i := 0; i < 10; i++ {
		j := i * 2
		s[j] = itoa(int(cmin[i]))
		s[j+1] = itoa(int(cmax[i]))
	}
	return s
}

func (v *V2) formatFloatMap(m map[float64]int) []string {
	if len(m) == 0 {
		return []string{"0", "0", "0"}
	}
	var (
		n   int
		min = math.MaxFloat64
		max = math.SmallestNonzeroFloat64
	)
	for k, v := range m {
		n += v
		if k > max {
			max = k
		}
		if k < min {
			min = k
		}
	}
	return []string{ftoa(min), ftoa(max), itoa(n)}
}

func (v *V2) formatIntMap(m map[int]int) []string {
	if len(m) == 0 {
		return []string{"0", "0", "0"}
	}
	var (
		n   int
		min = math.MaxInt32
		max int
	)
	for k, v := range m {
		n += v
		if k > max {
			max = k
		}
		if k < min {
			min = k
		}
	}
	return []string{itoa(min), itoa(max), itoa(n)}
}

func (v *V2) formatTica(m map[string]int) string {
	if len(m) == 0 {
		return ""
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, k := range keys {
		v := m[k]
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(k + ":(" + itoa(v) + ")")
	}
	return buf.String()
}

func (v V2) Empty() bool {
	return v.nKombi == 0
}

// vystup generator
type V1 struct {
	n, m      int
	hrx, hhrx *H
	riadok    Riadok
	Header    []string
}

func NewV1(a *Archiv) V1 {

	var header []string
	for i := 1; i <= a.n; i++ {
		header = append(header, strconv.Itoa(i))
	}
	header = append(header,
		"P",
		"N",
		"Sled PN",
		"PR",
		"Sled PNPr",
		"Mc",
		"Vc",
		"Sled McVc",
		"C19",
		"C0",
		"cC",
		"Cc",
		"CC",
		"Sled prirodzené kritéria",
		"ZH",
		"ZH presun (r/r+1)",
		"Sm",
		"Kk",
		"Ntica",
		"Ntica súčet",
		"Ntica súčin pozície a stĺpca",
		"Xtica",
		"ƩR 1-DO",
		"ΔƩR 1-DO",
		"ƩSTL 1-DO",
		"ΔƩSTL 1-DO",
		"Δ(ƩR1-DO-ƩSTL1-DO)",
		"HHrx",
		"ΔHHrx",
		"ƩR OD-DO",
		"ΔƩR OD-DO",
		"ƩSTL OD-DO",
		"ΔƩSTL OD-DO",
		"Δ(ƩROD-DO-ƩSTLOD-DO)",
		"Hrx",
		"ΔHrx",
		"ƩKombinacie",
		"Cifra 1", "Cifra 2", "Cifra 3", "Cifra 4", "Cifra 5",
		"Cifra 6", "Cifra 7", "Cifra 8", "Cifra 9", "Cifra 0",
	)
	return V1{
		n:      a.n,
		m:      a.m,
		hrx:    a.Hrx,
		hhrx:   a.HHrx,
		riadok: a.Riadok,
		Header: header,
	}
}

func (v V1) Riadok(k Kombinacia) []string {
	var (
		line   = make([]string, 0, len(v.Header)+v.n)
		r1, s1 = k.SucetRSNext(v.hhrx.Cisla)
		r2, s2 = k.SucetRSNext(v.hrx.Cisla)
		hrx    = v.hrx.ValueKombinacia(k)
		hhrx   = v.hhrx.ValueKombinacia(k)
	)
	for _, cislo := range k {
		line = append(line, strconv.Itoa(int(cislo)))
	}
	c := k.Cislovacky()
	cislovacky := c.Strings()
	line = append(line, cislovacky[0:2]...)
	line = append(line, k.SledPN())
	line = append(line, cislovacky[2])
	line = append(line, k.SledPNPr())
	line = append(line, cislovacky[3:5]...)
	line = append(line, k.SledMcVc())
	line = append(line, cislovacky[5:]...)
	line = append(line, k.SledPrirodzene())
	line = append(line,
		itoa(Zhoda(v.riadok.K, k)),
		NewZhodaPresun(v.riadok.K, k).String(),
		ftoa(Smernica(k, v.n, v.m)),
		ftoa(Korelacia(v.riadok.K, k, v.n, v.m)),
		Ntica(k).String(),
		NticaSucet(k).String(),
		NticaSucin(k).String(),
		Xtica(v.m, k).String(),

		ftoa(r1),
		ftoa(r1-v.riadok.R1), //dt
		ftoa(s1),
		ftoa(s1-v.riadok.S1), //dt
		ftoa(r1-s1),
		ftoa(hhrx),
		ftoa(hhrx-v.riadok.HHrx), //dt

		ftoa(r2),
		ftoa(r2-v.riadok.R2), //dt
		ftoa(s2),
		ftoa(s2-v.riadok.S2), //dt
		ftoa(r2-s2),
		ftoa(hrx),
		ftoa(hrx-v.riadok.Hrx), //dt

		itoa(k.Sucet()),
	)
	line = append(line, k.Cifrovacky().Strings()...)
	return line
}
