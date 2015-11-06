package psl

import (
	"bytes"
	"math"
	"sort"
	"strconv"

	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
)

// vystup filter
type V2 struct {
	n, m      int
	Hrx, HHrx *H
	r         Riadok
	p         Presun

	hrx                  float64
	zhoda, sucet         map[int]int
	hhrx, r1, s1, r2, s2 map[float64]int
	nKombi               uint64
	ntice, xtice         map[string]int
}

func NewV2(a *Archiv, sk Skupina) V2 {
	return V2{
		n:    a.n,
		m:    a.m,
		Hrx:  a.Hrx,
		HHrx: a.HHrx,
		r:    a.Riadok,
		p:    sk.Presun,

		hrx:   sk.Hrx,
		zhoda: make(map[int]int),
		sucet: make(map[int]int),
		hhrx:  make(map[float64]int),
		r1:    make(map[float64]int),
		s1:    make(map[float64]int),
		r2:    make(map[float64]int),
		s2:    make(map[float64]int),
		ntice: make(map[string]int),
		xtice: make(map[string]int),
	}
}

var HeaderV2 = []string{
	"ZH \"r\"/\"r+1\"", "HRX pre r+1", "ΔHRX", "X-cisla",
	"Počet Kombi",
	"ƩROD-DO (min)", "ƩROD-DO (max)", "ƩROD-DO (počet)",
	"N-tice", "X-tice",
	"ƩSTLOD-DO (min)", "ƩSTLOD-DO (max)", "ƩSTLOD-DO (počet)",
	"ƩKombinacie (min)", "ƩKombinacie (max)", "ƩKombinacie (počet)",
	"HHRX (min)", "HHRX (max)", "HHRX (počet)",
	"ƩR1-DO (min)", "ƩR1-DO (max)", "ƩR1-DO (počet)",
	"ƩSTL1-DO (min)", "ƩSTL1-DO (max)", "ƩSTL1-DO (počet)",
}

func (v *V2) Add(k Kombinacia) {
	v.zhoda[Zhoda(v.r.K, k)]++

	sucet := k.Sucet()
	if _, ok := v.sucet[sucet]; !ok {
		v.sucet[sucet] = 1
	}
	v.nKombi++

	R2, S2 := k.SucetRSNext(v.Hrx.Cisla)
	if _, ok := v.r2[R2]; !ok {
		v.r2[R2] = 1
	}
	if _, ok := v.s2[S2]; !ok {
		v.s2[S2] = 1
	}

	R1, S1 := k.SucetRSNext(v.HHrx.Cisla)
	if _, ok := v.r1[R1]; !ok {
		v.r1[R1] = 1
	}
	if _, ok := v.s1[S1]; !ok {
		v.s1[S1] = 1
	}

	hhrx := v.HHrx.ValueKombinacia(k)
	if _, ok := v.hhrx[hhrx]; !ok {
		v.hhrx[hhrx] = 1
	}

	v.ntice[Ntica(k).String()]++
	v.xtice[Xtica(v.m, k).String()]++
}

func (v V2) Riadok() []string {
	var r []string
	r = append(r,
		v.formatZhoda(v.zhoda),
		ftoa(v.hrx),
		ftoa(v.r.Hrx-v.hrx),
		v.p.String(),
		strconv.FormatUint(v.nKombi, 10),
	)
	r = append(r, v.formatFloatMap(v.r2)...)
	r = append(r, v.formatTica(v.ntice), v.formatTica(v.xtice))
	r = append(r, v.formatFloatMap(v.s2)...)
	r = append(r, v.formatIntMap(v.sucet)...)
	r = append(r, v.formatFloatMap(v.hhrx)...)
	r = append(r, v.formatFloatMap(v.r1)...)
	r = append(r, v.formatFloatMap(v.s1)...)
	return r
}

func (v *V2) formatZhoda(m map[int]int) string {
	if len(m) == 0 {
		return "0:(0)"
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var buf bytes.Buffer
	for i, k := range keys {
		v := m[k]
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(itoa(k) + ":(" + itoa(v) + ")")
	}
	return buf.String()
}

func (v *V2) formatFloatMap(m map[float64]int) []string {
	if len(m) == 0 {
		return []string{"0", "0", "0"}
	}
	var (
		n   int
		min = math.MaxFloat64
		max float64
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
		"P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC",
		"ZH", "ZH presun (r/r+1)", "Sm", "Kk", "Ntica", "Ntica súčet",
		"Ntica súčin pozície a stĺpca", "X-tice", "ƩR1-DO", "ΔƩR1-DO",
		"ƩSTL1-DO", "ΔƩSTL1-DO", "Δ(ƩR1-DO-ƩSTL1-DO)", "HHRX", "ΔHHRX",
		"ƩR OD-DO", "ΔƩR OD-DO", "ƩSTL OD-DO", "ΔƩSTL OD-DO", "Δ(ƩROD-DO-ƩSTLOD-DO)",
		"HRX", "ΔHRX", "ƩKombinacie", "Cifra 1", "Cifra 2", "Cifra 3", "Cifra 4", "Cifra 5",
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
		line   = make([]string, 0, 35)
		r1, s1 = k.SucetRSNext(v.hhrx.Cisla)
		r2, s2 = k.SucetRSNext(v.hrx.Cisla)
		hrx    = v.hrx.ValueKombinacia(k)
		hhrx   = v.hhrx.ValueKombinacia(k)
	)
	for _, cislo := range k {
		line = append(line, strconv.Itoa(int(cislo)))
	}
	cislovacky := k.Cislovacky()
	line = append(line, cislovacky.Strings()...)
	line = append(line,
		itoa(Zhoda(v.riadok.K, k)),
		ZhodaPresun(v.riadok.K, k).String(),
		ftoa(Smernica(v.n, v.m, k)),
		ftoa(Korelacia(v.n, v.m, v.riadok.K, k)),
		Ntica(k).String(),
		NticaSucet(k).String(),
		NticaSucin(k).String(),
		Xtica(v.m, k).String(),

		ftoa(r1),
		ftoa(v.riadok.R1-r1), //dt
		ftoa(s1),
		ftoa(v.riadok.S1-s1), //dt
		ftoa(r1-s1),
		ftoa(hhrx),
		ftoa(v.riadok.HHrx-hhrx), //dt

		ftoa(r2),
		ftoa(v.riadok.R2-r2), //dt
		ftoa(s2),
		ftoa(v.riadok.S2-s2), //dt
		ftoa(r2-s2),
		ftoa(hrx),
		ftoa(v.riadok.Hrx-hrx), //dt

		itoa(k.Sucet()),
	)
	line = append(line, k.Cifrovacky().Strings()...)
	return line
}
