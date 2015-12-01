package psl

import (
	"bytes"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

type zhodaRiadok struct {
	// k je aktualna v riadku
	k      Kombinacia
	zh     int
	krok   int
	krokZh bool
	presun []int // r-1/r
}

func makeZhodaRiadok(k Kombinacia, z0 *zhodaRiadok) *zhodaRiadok {
	if z0 == nil {
		z0 = &zhodaRiadok{
			k:      Kombinacia{},
			presun: make([]int, len(k)),
			krok:   1,
		}
	}
	z := &zhodaRiadok{
		k:      k,
		zh:     Zhoda(k, z0.k),
		presun: make([]int, len(k)),
	}
	if z0.krokZh {
		z.krok = 1
	} else {
		z.krok = z0.krok
	}
	if z.zh > 0 {
		z.krokZh = true
		for i, c1 := range z.k {
			for j, c0 := range z0.k {
				if c0 == c1 {
					z.presun[i] = j + 1
					break
				}
			}
		}
	} else {
		z.krok++
	}
	return z
}

func zhodaRiadokHeader(n int) []string {
	header := []string{"ZH"}
	for i := 0; i < n; i++ {
		header = append(header, strconv.Itoa(i+1))
	}
	header = append(header, "Pozícia zhoda", "Krok")
	for i := 0; i < n; i++ {
		header = append(header, strconv.Itoa(i+1)+" stl r-1/r")
	}
	return header
}

func (z *zhodaRiadok) Strings() []string {
	var s []string
	s = append(s, strconv.Itoa(z.zh))
	for i, c := range z.presun {
		if c > 0 {
			s = append(s, strconv.Itoa(int(z.k[i])))
		} else {
			s = append(s, "")
		}
	}
	pZH := []string{}
	for i, c := range z.presun {
		if c > 0 {
			pZH = append(pZH, strconv.Itoa(c)+"|"+strconv.Itoa(i+1))
		}
	}
	s = append(s, strings.Join(pZH, ", "))
	if z.zh > 0 {
		s = append(s, strconv.Itoa(z.krok))
	} else {
		s = append(s, "")
	}
	for i, c := range z.presun {
		if c > 0 {
			s = append(s, strconv.Itoa(c)+"|"+strconv.Itoa(i+1))
		} else {
			s = append(s, "")
		}
	}
	return s
}

func (a *Archiv) mapaZhoda() error {
	header := make([]string, len(a.origHeader))
	copy(header, a.origHeader)
	header = append(header, zhodaRiadokHeader(a.n)...)

	w := NewCsvMaxWriter(a.WorkingDir, "MapaZhoda", [][]string{header})
	defer w.Close()

	var zh0 *zhodaRiadok
	var mapaZhoda []*zhodaRiadok
	for _, r := range a.riadky {
		mapaZhoda = append(mapaZhoda, makeZhodaRiadok(r.K, zh0))
		zh0 = mapaZhoda[len(mapaZhoda)-1]
	}
	for i, z := range mapaZhoda {
		var record []string
		record = append(record, a.riadky[i].origStrings...)
		record = append(record, z.Strings()...)
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func nticaPozicie(k Kombinacia) []string {
	var s []string
	for i, ok := range NticaPozicie(k) {
		if ok == 1 {
			s = append(s, itoa(int(k[i])))
		} else {
			s = append(s, "")
		}
	}
	return s
}

func krokNtica(ntica Tica, nticeVsetky []string, counter map[string]int) []string {
	var s []string
	for _, nticaString := range nticeVsetky {
		if nticaString == ntica.String() {
			s = append(s, itoa(counter[nticaString]))
			counter[nticaString] = 1
		} else {
			if counter[nticaString] > 0 {
				counter[nticaString]++
			}
			s = append(s, "")
		}
	}
	return s
}

func (a *Archiv) mapaNtice() error {
	var (
		s           [][]string
		nticeVsetky = nticeStr(a.n)
		counter     = make(map[string]int)
	)
	for _, tica := range nticeVsetky {
		counter[tica] = 0
	}
	for _, r := range a.riadky {
		var riadok []string
		riadok = append(riadok, r.Ntica.String())
		riadok = append(riadok, nticaPozicie(r.K)...)
		riadok = append(riadok, NticaSucet(r.K).String())
		riadok = append(riadok, NticaSucin(r.K).String())
		riadok = append(riadok, krokNtica(r.Ntica, nticeVsetky, counter)...)
		s = append(s, riadok)
	}

	header := make([]string, len(a.origHeader))
	copy(header, a.origHeader)
	header = append(header, "N-tica")
	for i := 1; i <= a.n; i++ {
		header = append(header, itoa(i))
	}
	header = append(header, "Sucet N-tic", "Sucin pozicie a stlpca")
	header = append(header, nticeVsetky...)

	w := NewCsvMaxWriter(a.WorkingDir, "MapaNtice", [][]string{header})
	defer w.Close()

	for i, r := range a.riadky {
		var riadok []string
		riadok = append(riadok, r.origStrings...)
		riadok = append(riadok, s[i]...)
		if err := w.Write(riadok); err != nil {
			return err
		}
	}
	return nil
}

func ntice(n int, emit func([]int)) {
	if n < 2 {
		panic("ntice: n < 2")
	}
	tica := make([]int, n)
	end := make([]int, n)
	tica[0] = n
	end[n-1] = 1

	i := 0
	for {
		sum := 0
		for j, e := range tica {
			sum += int(e) * (j + 1)
		}
		if sum == n {
			emit(tica)
			tica[i]--
			i++
		} else if sum < n {
			tica[i]++
		} else {
			tica[i]--
			i++
		}
		if i == n {
			i--
			for tica[i] == 0 {
				i--
			}
			tica[i]--
			i++
		}

		equal := true
		for p := 0; p < n; p++ {
			if tica[p] != end[p] {
				equal = false
			}
		}
		if equal {
			break
		}
	}
	emit(tica)
}

func nticeNtice(n int) []Tica {
	var nticeVsetky []Tica
	ntice(n, func(tica []int) {
		n := make(Tica, len(tica))
		for i := range tica {
			n[i] = byte(tica[i])
		}
		nticeVsetky = append(nticeVsetky, n)
	})
	return nticeVsetky
}

func nticeStr(n int) []string {
	var (
		tice []string
		buf  bytes.Buffer
	)
	ntice(n, func(tica []int) {
		buf.Reset()
		for i, t := range tica {
			if i > 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(itoa(t))
		}
		tice = append(tice, buf.String())
	})
	return tice
}

func (a *Archiv) mapaXtice() error {

	xlsxFile := xlsx.NewFile()
	sheet := xlsxFile.AddSheet("Mapa Xtice")

	for _, riadok := range a.riadky {
		row := sheet.AddRow()
		for _, s := range riadok.origStrings {
			cell := row.AddCell()
			cell.SetString(s)
		}
		for _, cislo := range riadok.K {
			cell := row.AddCell()
			cell.SetInt(int(cislo))

			style := xlsx.NewStyle()
			switch (int(cislo) - 1) / 10 {
			case 0: // light red
				style.Fill = *xlsx.NewFill("solid", "DB7F7F", "DB7F7F")
			case 1: // light green
				style.Fill = *xlsx.NewFill("solid", "7FBF7F", "7FBF7F")
			case 2: // blue
				style.Fill = *xlsx.NewFill("solid", "6666FF", "6666FF")
			case 3: // yelow
				style.Fill = *xlsx.NewFill("solid", "FFFF99", "FFFF99")
			case 4: // white
				style.Fill = *xlsx.NewFill("solid", "FFFFFF", "FFFFFF")
			case 5: // purple
				style.Fill = *xlsx.NewFill("solid", "FFCCFF", "FFCCFF")
			case 6: // nobody knows
				style.Fill = *xlsx.NewFill("solid", "CCFFFF", "CCFFFF")
			case 7: //  nobody knows 2
				style.Fill = *xlsx.NewFill("solid", "FFCCCC", "FFCCCC")
			case 8: // green ?
				style.Fill = *xlsx.NewFill("solid", "80FF80", "80FF80")
			default:
			}
			cell.SetStyle(style)
		}
	}
	return xlsxFile.Save(filepath.Join(a.WorkingDir, "MapaXtice.xlsx"))
}
