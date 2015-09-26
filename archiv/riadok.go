package archiv

import (
	"strconv"
	"strings"

	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

type Uc struct {
	Cislo  byte
	Riadok int
}

type Riadok struct {
	n, m        int
	origStrings []string

	Pc             uint16
	K              komb.Kombinacia
	C              num.Cislovacky
	Zh             int
	Sm, DtSm       float64
	Kk, DtKk       float64
	Ntica          komb.Tica
	Xtica          komb.Tica
	R1, DtR1       float64
	Rp1            float64
	S1, DtS1       float64
	Sp1            float64
	R1mS1          float64
	HHrx, DtHHrx   float64
	R2, DtR2       float64
	S2, DtS2       float64
	R2mS2          float64
	Hrx, DtHrx     float64
	Sucet, DtSucet int
	Uc
	Cifrovacky komb.Cifrovacky
}

func (r *Riadok) Add(k komb.Kombinacia, n1, n2 num.Nums, hrx, hhrx float64) {

	Sm := komb.Smernica(r.n, r.m, k)
	R1, S1 := k.SucetRS(n1)
	R2, S2 := k.SucetRS(n2)

	if r.K != nil {
		r.Zh = komb.Zhoda(r.K, k)
		Kk := komb.Korelacia(r.n, r.m, r.K, k)
		r.DtKk = Kk - r.Kk
		r.Kk = Kk

		// Dt
		r.DtR1 = R1 - r.R1
		r.DtS1 = S1 - r.S1
		r.DtR2 = R2 - r.R2
		r.DtS2 = S2 - r.S2
		r.DtHrx = hrx - r.Hrx
		r.DtHHrx = hhrx - r.HHrx
		r.DtSm = Sm - r.Sm
		r.DtSucet = k.Sucet() - r.Sucet
	}
	r.Sm = Sm
	r.R1 = R1
	// r.Rp1 = n1.RNext()
	r.S1 = S1
	// r.Sp1 = sumSnext(k, n1)
	r.Rp1, r.Sp1 = k.SucetRSNext(n1)
	r.R1mS1 = R1 - S1
	r.R2 = R2
	r.S2 = S2
	r.R2mS2 = R2 - S2
	r.Hrx = hrx
	r.HHrx = hhrx
	r.Sucet = k.Sucet()

	r.K = k
	r.C = k.Cislovacky()
	r.Ntica = komb.Ntica(k)
	r.Xtica = komb.Xtica(r.m, k)

	r.Cifrovacky = k.Cifrovacky()
}

func (r Riadok) record() []string {
	rec := make([]string, 0, len(header))
	rec = append(rec, itoa(int(r.Pc)), r.K.String())
	rec = append(rec, r.C.Strings()...)
	rec = append(rec,
		itoa(r.Zh),
		ftoa(r.Sm),
		ftoa(r.DtSm),
		ftoa(r.Kk),
		ftoa(r.DtKk),
		r.Ntica.String(),
		r.Xtica.String(),
		ftoa(r.R1),
		ftoa(r.DtR1),
		ftoa(r.Rp1),
		ftoa(r.S1),
		ftoa(r.DtS1),
		ftoa(r.Sp1),
		ftoa(r.R1mS1),
		ftoa(r.HHrx),
		ftoa(r.DtHHrx),
		ftoa(r.R2),
		ftoa(r.DtR2),
		ftoa(r.S2),
		ftoa(r.DtS2),
		ftoa(r.R2mS2),
		ftoa(r.Hrx),
		ftoa(r.DtHrx),
		itoa(r.Sucet),
		itoa(r.DtSucet),
		itoa(int(r.Cislo)),
		itoa(r.Riadok),
	)
	rec = append(rec, r.Cifrovacky.Strings()...)
	return rec
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func ftoa(f float64) string {
	s := strconv.FormatFloat(f, 'g', -1, 64)
	return strings.Replace(s, ".", ",", 1)
}
