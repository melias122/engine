package archiv

import (
	"bytes"
	"math"
	"strconv"
	"strings"

	"github.com/melias122/psl/num"
)

type Uc struct {
	Cislo  byte
	Riadok int
}

type Riadok struct {
	n, m int

	Pc             uint16
	K              []byte
	C              num.C
	Zh             int
	Sm, DtSm       float64
	Kk, DtKk       float64
	Ntica          []byte
	Xtica          []byte
	R1, DtR1       float64
	HHrx, DtHHrx   float64
	R2, DtR2       float64
	Hrx, DtHrx     float64
	Sucet, DtSucet int
	Uc
}

func (r *Riadok) Add(k []byte, n1, n2 *num.N, hrx, hhrx float64) {

	Sm := smernica(r.n, r.m, k)

	if r.K != nil {
		r.Zh = zhoda(r.K, k)
		Kk := korelacia(r.n, r.m, r.K, k)
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
}

func (r Riadok) record() []string {
	rec := make([]string, 0, len(header))
	rec = append(rec, itoa(int(r.Pc)), kombToString(r.K))
	for _, c := range r.C {
		rec = append(rec, itoa(int(c)))
	}
	rec = append(rec,
		itoa(r.Zh),
		ftoa(r.Sm),
		ftoa(r.DtSm),
		ftoa(r.Kk),
		ftoa(r.DtKk),
		"", // ntica
		"", // xtica
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

func smernica(n, m int, k []byte) float64 {
	if len(k) != n {
		return 0
	}
	var (
		sm  float64
		nSm int
	)
	for i := 0; i < n-1; i++ {
		// for i, ki := range k[:len(k)-1] {
		for j := i + 1; j < n; j++ {
			// for j, kj := range k[i+1:] {
			p1 := float64(k[j]) - float64(k[i])
			p2 := float64(j) - float64(i)

			p1 /= float64(m - 1)
			p2 /= float64(n - 1)
			p1 /= p2

			sm += p1
			nSm++
		}
	}
	if nSm > 0 {
		sm /= float64(nSm)
	}
	return sm
}

func zhoda(k0, k1 []byte) int {
	var zh, i, j int
	for i < len(k0) && j < len(k1) {
		if k0[i] == k1[j] {
			zh++
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return zh
}

func korelacia(n, m int, k0, k1 []byte) float64 {
	if len(k0) == 0 || len(k1) == 0 {
		return 0
	}
	var kk float64
	for i := 0; i < n; i++ {
		a := float64(k1[i])
		p := float64(k0[i])
		kk += math.Pow((a-p)/float64(m), 4) / float64(n)
	}
	return math.Pow(float64(1)-math.Sqrt(kk), 8)
}

func kombToString(k []byte) string {
	var buf bytes.Buffer
	for i, el := range k {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(int(el)))
	}
	return buf.String()
}
