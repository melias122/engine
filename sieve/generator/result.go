package generator

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"unicode/utf8"

	"github.com/melias122/engine/csv"
	"github.com/melias122/engine/engine"
	"github.com/melias122/engine/filter"
)

type syncWriter struct {
	sync.Mutex
	w *csv.CsvMaxWriter
}

func newSyncWriter(w *csv.CsvMaxWriter) *syncWriter {
	return &syncWriter{
		w: w,
	}
}

func (s *syncWriter) Write(record []string) (err error) {
	s.Lock()
	err = s.w.Write(record)
	s.Unlock()
	return
}

func (s *syncWriter) Close() error {
	s.Lock()
	defer s.Unlock()
	return s.w.Close()
}

// result implemets engine.Filter
type result struct {
	filter.Filter

	n, m      int
	hrx, hhrx *engine.H
	riadok    engine.Riadok
	header    []string
	w         *syncWriter
}

func newResultFilter(w *csv.CsvMaxWriter, a *engine.Archiv, n, m int) *result {
	var header []string
	for i := 1; i <= n; i++ {
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
	return &result{
		n:      n,
		m:      m,
		hrx:    a.Hrx,
		hhrx:   a.HHrx,
		riadok: a.Riadok,
		header: header,
		w:      newSyncWriter(w),
	}
}

func (r *result) Check(k engine.Kombinacia) bool {
	// write only full length k, but accept it
	if len(k) != r.n {
		return true
	}
	var (
		line   = make([]string, 0, len(r.header)+r.n)
		r1, s1 = k.SucetRSNext(r.hhrx.Cisla)
		r2, s2 = k.SucetRSNext(r.hrx.Cisla)
		hrx    = r.hrx.Value(k)
		hhrx   = r.hhrx.Value(k)
	)
	for _, cislo := range k {
		line = append(line, strconv.Itoa(int(cislo)))
	}
	c := engine.NewKCislovacky(k)
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
		itoa(engine.Zhoda(r.riadok.K, k)),
		engine.NewZhodaPresun(r.riadok.K, k).String(),
		ftoa(engine.Smernica(k, r.n, r.m)),
		ftoa(engine.Korelacia(r.riadok.K, k, r.n, r.m)),
		engine.NewNtica(k).String(),
		engine.NticaSucet(k).String(),
		engine.NticaSucin(k).String(),
		engine.NewXtica(k, r.m).String(),

		ftoa(r1),
		ftoa(r1-r.riadok.R1), //dt
		ftoa(s1),
		ftoa(s1-r.riadok.S1), //dt
		ftoa(r1-s1),
		ftoa(hhrx),
		ftoa(hhrx-r.riadok.HHrx), //dt

		ftoa(r2),
		ftoa(r2-r.riadok.R2), //dt
		ftoa(s2),
		ftoa(s2-r.riadok.S2), //dt
		ftoa(r2-s2),
		ftoa(hrx),
		ftoa(hrx-r.riadok.Hrx), //dt

		itoa(k.Sucet()),
	)
	line = append(line, engine.NewCifrovacky(k).Strings()...)

	// TODO: err
	if err := r.w.Write(line); err != nil {
		log.Println("result: ", err)
		return false
	}
	return true
}

func (*result) CheckSkupina(engine.Skupina) bool {
	return true
}

func (r *result) String() string {
	return fmt.Sprint(r.w.w.TotalRowsWriten())
}

func (r *result) Close() error {
	return r.w.Close()
}

func itoa(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func ftoa(f float64) string {
	buf := make([]byte, 0, 64)
	buf = strconv.AppendFloat(buf, f, 'g', -1, 64)
	for i, w := 0, 0; i < len(buf); i += w {
		runeValue, width := utf8.DecodeRune(buf[i:])
		if runeValue == '.' {
			buf[i] = ','
			break
		}
		w = width
	}
	return string(buf[:len(buf)])
}
