package psl

import (
	"fmt"
	"path/filepath"

	"github.com/tealeg/xlsx"
)

type statistikaXticeXn struct {
	max int

	pocet []int

	krok [][]float64

	vyskyt []int
}

func newStatistikaXticeXn(max, i int) *statistikaXticeXn {
	sx := statistikaXticeXn{
		max:    max,
		pocet:  make([]int, i),
		krok:   make([][]float64, i),
		vyskyt: make([]int, i),
	}
	for n := range sx.krok {
		sx.krok[n] = []float64{1}
	}
	return &sx
}

func (s *statistikaXticeXn) add(i int, vyskyt int) {
	s.pocet[i]++
	s.krok[i] = append(s.krok[i], 1)
	s.vyskyt[i] = vyskyt
}

func (s *statistikaXticeXn) inc(i int) {
	k := s.krok[i]
	k[len(k)-1]++
}

func (sx *statistikaXticeXn) strings(udalost string) [][]string {
	s := [][]string{
		[]string{},
	}

	addRow := func() {
		s = append(s, []string{})
	}
	addCol := func(str ...string) {
		i := len(s) - 1
		s[i] = append(s[i], str...)
	}

	addRow()
	addCol("Počet udalosť " + udalost)
	for _, u := range sx.pocet {
		addCol(itoa(u))
	}

	addRow()
	addCol("Súčet differencií krok udalosť")
	for _, u := range sx.pocet {
		addCol(itoa(sx.max - u))
	}

	addRow()
	addCol("Krok aritmeticky priemer udalosť")
	for _, k := range sx.krok {
		i := int(aritmetickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Krok harmonický priemer udalosť")
	for _, k := range sx.krok {
		i := int(harmonickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Krok vážený priemer udalosť")
	for _, k := range sx.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Krok geometrický priemer udalosť")
	for _, k := range sx.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i))
	}

	addRow()
	addCol("Riadok posledný výskyt udalosť")
	for _, k := range sx.vyskyt {
		addCol(itoa(k))
	}

	addRow()
	addCol("Krok aritmeticky priemer udalosť + riadok posledný výskyt")
	for j, k := range sx.krok {
		i := int(aritmetickyPriemer(k))
		addCol(itoa(i + sx.vyskyt[j]))
	}

	addRow()
	addCol("Krok harmonický priemer udalosť + riadok posledný výskyt")
	for j, k := range sx.krok {
		i := int(harmonickyPriemer(k))
		addCol(itoa(i + sx.vyskyt[j]))
	}

	addRow()
	addCol("Krok vážený priemer udalosť + riadok posledný výskyt")
	for j, k := range sx.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i + sx.vyskyt[j]))
	}

	addRow()
	addCol("Krok geometrický priemer udalosť + riadok posledný výskyt")
	for j, k := range sx.krok {
		i := int(vazenyAritmetickyPriemer(k))
		addCol(itoa(i + sx.vyskyt[j]))
	}

	return s
}

type mapaXtice2 struct {
	n, m int
	k    []Kombinacia

	s  []*statistikaXticeXn
	s2 []*statistikaXticeXn
}

func newMapaXtice2(r []Riadok, n, m int) mapaXtice2 {

	// maximalny pocet cisiel v jednom stlpci xtice je 10 + 0
	s2max := n + 1
	if s2max > 11 {
		s2max = 11
	}

	x := mapaXtice2{
		n:  n,
		m:  m,
		k:  make([]Kombinacia, 0, len(r)),
		s:  make([]*statistikaXticeXn, (m+9)/10),
		s2: make([]*statistikaXticeXn, s2max),
	}

	for i := range x.s {
		x.s[i] = newStatistikaXticeXn(len(r), n)
	}

	for i := range x.s2 {
		x.s2[i] = newStatistikaXticeXn(len(r), (m+9)/10)
	}

	for _, r := range r {
		x.k = append(x.k, r.K)

		// statistika stl n
		for stl, c := range r.K {
			i := (int(c) - 1) / 10
			x.s[i].add(stl, r.Pc)
			for j, xs := range x.s {
				if j != i {
					xs.inc(stl)
				}

			}
		}

		// statistika xn
		for stl, c := range r.Xtica {
			i := int(c)
			x.s2[i].add(stl, r.Pc)
			for j, xs2 := range x.s2 {
				if j != i {
					xs2.inc(stl)
				}

			}
		}
	}

	return x
}

func (m *mapaXtice2) strings() [][]string {
	s := [][]string{
		[]string{
			"Kombinacie",
		},
	}

	addRow := func() {
		s = append(s, []string{})
	}
	addCol := func(str ...string) {
		i := len(s) - 1
		s[i] = append(s[i], str...)
	}

	for i := 1; i <= m.n; i++ {
		addCol("STL " + itoa(i))
	}
	addCol("")
	addCol("")
	for i := 0; i < (m.m+9)/10; i++ {
		// 1, 11, 21, 31, ...
		j := 1 + i*10
		addCol(fmt.Sprintf("X%d (%d-%d)", i+1, j, j+9))
	}

	for _, k := range m.k {
		addRow()
		addCol(k.String())
		xtica := Xtica(m.m, k)
		for i, j := range xtica {
			for k := 0; k < int(j); k++ {
				addCol("X" + itoa(i+1))
			}
		}
		addCol("")
		addCol("")
		for _, i := range xtica {
			addCol(itoa(int(i)))
		}
	}
	var xnstrings [][]string
	for i, s := range m.s {
		xnstrings = append(xnstrings, s.strings("X"+itoa(i+1))...)
	}

	var cstrings [][]string
	for i, s := range m.s2 {
		cstrings = append(cstrings, s.strings(itoa(i))...)
	}

	lens := len(s)

	s = append(s, xnstrings...)

	for i, str := range cstrings {
		if lens+i == len(s) {
			addRow()
			// add pad
			addCol(make([]string, m.n+1)...)
		}
		// addCol(str...)
		s[lens+i] = append(s[lens+i], "")
		s[lens+i] = append(s[lens+i], str...)
	}

	return s
}

func (a *Archiv) mapaXtice2() error {

	w := NewCsvMaxWriter("MapaXtice2", a.WorkingDir)
	defer w.Close()

	mx := newMapaXtice2(a.riadky, a.n, a.m)
	for _, s := range mx.strings() {
		if err := w.Write(s); err != nil {
			return err
		}
	}

	return nil
}

func (a *Archiv) mapaXtice() error {

	xlsxFile := xlsx.NewFile()
	sheet := xlsxFile.AddSheet("Mapa Xtice (farebna)")

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
	return xlsxFile.Save(filepath.Join(a.WorkingDir, "MapaXtice(farebna).xlsx"))
}
