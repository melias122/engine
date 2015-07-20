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
	n, m int

	Pc             uint16
	K              komb.Kombinacia
	C              num.C
	Zh             int
	Sm, DtSm       float64
	Kk, DtKk       float64
	Ntica          komb.Tica
	Xtica          komb.Tica
	R1, DtR1       float64
	HHrx, DtHHrx   float64
	R2, DtR2       float64
	Hrx, DtHrx     float64
	Sucet, DtSucet int
	Uc
}

func (r *Riadok) Add(k komb.Kombinacia, n1, n2 *num.N, hrx, hhrx float64) {

	Sm := komb.Smernica(r.n, r.m, k)

	if r.K != nil {
		r.Zh = komb.Zhoda(r.K, k)
		Kk := komb.Korelacia(r.n, r.m, r.K, k)
		r.DtKk = Kk - r.Kk
		r.Kk = Kk

		r.DtR1 = n1.R() - r.R1
		r.DtR2 = n2.R() - r.R2
		r.DtHrx = hrx - r.Hrx
		r.DtHHrx = hhrx - r.HHrx
		r.DtSm = Sm - r.Sm
		r.DtSucet = n1.Cislo() - r.Sucet
	}
	r.Sm = Sm
	r.R1 = n1.R()
	r.R2 = n2.R()
	r.Hrx = hrx
	r.HHrx = hhrx
	r.Sucet = n1.Cislo()

	r.K = k
	r.C = n1.C()
	r.Ntica = komb.Ntica(k)
	r.Xtica = komb.Xtica(r.m, k)
}

func (r Riadok) record() []string {
	rec := make([]string, 0, len(header))
	rec = append(rec, itoa(int(r.Pc)), r.K.String())
	for _, c := range r.C {
		rec = append(rec, itoa(int(c)))
	}
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
		ftoa(r.HHrx),
		ftoa(r.DtHHrx),
		ftoa(r.R2),
		ftoa(r.DtR2),
		ftoa(r.Hrx),
		ftoa(r.DtHrx),
		itoa(r.Sucet),
		itoa(r.DtSucet),
		itoa(int(r.Cislo)),
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
