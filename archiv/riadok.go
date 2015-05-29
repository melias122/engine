package archiv

import (
	"psl2/komb"
	"strconv"
	"strings"
)

type uc struct {
	Cislo  int
	Riadok int
}

type Riadok struct {
	Pc      int
	K       *komb.K
	Zh      int
	Sm      float64
	Sm_dt   float64
	Kk      float64
	Kk_dt   float64
	R1      float64
	R1_dt   float64
	HHrx    float64
	HHrx_dt float64
	R2      float64
	R2_dt   float64
	Hrx     float64
	Hrx_dt  float64
	// Sucet    int
	Sucet_dt int
	uc
}

func (r *Riadok) add(k *komb.K, Hrx, HHrx float64) {
	r.K = k
	r.Sm = k.Sm()
	r.R1 = k.R1()
	r.HHrx = HHrx
	r.R2 = k.R2()
	r.Hrx = Hrx

}

// funkcia spravi rozdiel riadov r1 a r0
// a vysledok ulozi do r1
func (r1 *Riadok) diff(r0 Riadok) {
	r1.Zh = r1.K.Zh(r0.K)
	r1.Sm_dt = r1.Sm - r0.Sm
	r1.Kk = r1.K.Kk(r0.K)
	r1.Kk_dt = r1.Kk - r0.Kk
	r1.R1_dt = r1.R1 - r0.R1
	r1.HHrx_dt = r1.HHrx - r0.HHrx
	r1.R2_dt = r1.R2 - r0.R2
	r1.Hrx_dt = r1.Hrx - r0.Hrx
	r1.Sucet_dt = r1.K.Sucet() - r0.K.Sucet()
}

func (r Riadok) record() []string {
	rec := make([]string, 0, len(header))
	rec = append(rec,
		itoa(r.Pc),
		r.K.String(),
	)
	for _, c := range r.K.C() {
		rec = append(rec, itoa(int(c)))
	}
	// rec = append(rec, strings.Split(" ", c.String())...)
	rec = append(rec,
		itoa(r.Zh),
		ftoa(r.Sm),
		ftoa(r.Sm_dt),
		ftoa(r.Kk),
		ftoa(r.Kk_dt),
		r.K.Ntica().String(),
		r.K.Xtica().String(),
		ftoa(r.R1),
		ftoa(r.R1_dt),
		ftoa(r.HHrx),
		ftoa(r.HHrx_dt),
		ftoa(r.R2),
		ftoa(r.R2_dt),
		ftoa(r.Hrx),
		ftoa(r.Hrx_dt),
		itoa(r.K.Sucet()),
		itoa(r.Sucet_dt),
		itoa(r.Cislo),
		itoa(r.Riadok),
	)
	return rec
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func ftoa(f float64) string {
	s := strconv.FormatFloat(f, 'g', -1, 64)
	return strings.Replace(s, ".", ",", 1)
}
